package etcd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	hash "github.com/mitchellh/hashstructure"
	"github.com/pangdogs/galaxy/plugin/logger"
	"github.com/pangdogs/galaxy/plugin/registry"
	"github.com/pangdogs/galaxy/service"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/client/v3"
	"path"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	prefix = "/galaxy/registry/"
)

func newRegistry(options ...Option) registry.Registry {
	opts := Options{}
	Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return &_EtcdRegistry{
		options:  opts,
		register: make(map[string]uint64),
		leases:   make(map[string]clientv3.LeaseID),
	}
}

type _EtcdRegistry struct {
	options    Options
	serviceCtx service.Context
	client     *clientv3.Client
	register   map[string]uint64
	leases     map[string]clientv3.LeaseID
	sync.RWMutex
}

func (e *_EtcdRegistry) Init(ctx service.Context) {
	e.serviceCtx = ctx

	client, err := clientv3.New(e.configure())
	if err != nil {
		panic(err)
	}
	e.client = client
}

func (e *_EtcdRegistry) Shut() {
	if e.client != nil {
		e.client.Close()
	}
}

func (e *_EtcdRegistry) Register(ctx context.Context, service registry.Service, ttl time.Duration) error {
	if len(service.Nodes) <= 0 {
		return errors.New("require at least one node")
	}

	var anyErr error

	for _, node := range service.Nodes {
		if err := e.registerNode(ctx, service, node, ttl); err != nil {
			anyErr = err
		}
	}

	return anyErr
}

func (e *_EtcdRegistry) Deregister(ctx context.Context, service registry.Service) error {
	if len(service.Nodes) <= 0 {
		return errors.New("require at least one node")
	}

	for _, node := range service.Nodes {
		np := nodePath(service.Name, node.Id)

		e.Lock()
		// delete our hash of the service
		delete(e.register, np)
		// delete our lease of the service
		delete(e.leases, np)
		e.Unlock()

		ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
		defer cancel()

		logger.Tracef(e.serviceCtx, "deregistering %s", np)

		_, err := e.client.Delete(ctx, np)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *_EtcdRegistry) GetService(ctx context.Context, serviceName string) ([]registry.Service, error) {
	ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
	defer cancel()

	rsp, err := e.client.Get(ctx, servicePath(serviceName)+"/", clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	if len(rsp.Kvs) <= 0 {
		return nil, registry.ErrNotFound
	}

	serviceMap := map[string]registry.Service{}

	for _, n := range rsp.Kvs {
		if sn := decode(n.Value); sn != nil {
			s, ok := serviceMap[sn.Version]
			if !ok {
				s = registry.Service{
					Name:      sn.Name,
					Version:   sn.Version,
					Metadata:  sn.Metadata,
					Endpoints: sn.Endpoints,
				}
				serviceMap[s.Version] = s
			}

			s.Nodes = append(s.Nodes, sn.Nodes...)
		}
	}

	services := make([]registry.Service, 0, len(serviceMap))
	for _, service := range serviceMap {
		services = append(services, service)
	}

	return services, nil
}

func (e *_EtcdRegistry) ListServices(ctx context.Context) ([]registry.Service, error) {
	versions := make(map[string]registry.Service)

	ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
	defer cancel()

	rsp, err := e.client.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	if len(rsp.Kvs) <= 0 {
		return []registry.Service{}, nil
	}

	for _, n := range rsp.Kvs {
		sn := decode(n.Value)
		if sn == nil {
			continue
		}
		v, ok := versions[sn.Name+sn.Version]
		if !ok {
			versions[sn.Name+sn.Version] = *sn
			continue
		}
		// append to service:version nodes
		v.Nodes = append(v.Nodes, sn.Nodes...)
	}

	services := make([]registry.Service, 0, len(versions))
	for _, service := range versions {
		services = append(services, service)
	}

	// sort the services
	sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })

	return services, nil
}

func (e *_EtcdRegistry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return newEtcdWatcher(ctx, e, e.options.Timeout, serviceName)
}

func (e *_EtcdRegistry) configure() clientv3.Config {
	if e.options.EtcdConfig != nil {
		return *e.options.EtcdConfig
	}

	config := clientv3.Config{
		Endpoints:   e.options.Endpoints,
		DialTimeout: e.options.Timeout,
		Username:    e.options.Username,
		Password:    e.options.Password,
		LogConfig:   e.options.ZapConfig,
	}

	if e.options.Secure || e.options.TLSConfig != nil {
		tlsConfig := e.options.TLSConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		config.TLS = tlsConfig
	}

	return config
}

func (e *_EtcdRegistry) registerNode(ctx context.Context, s registry.Service, node registry.Node, ttl time.Duration) error {
	if len(s.Nodes) <= 0 {
		return errors.New("require at least one node")
	}

	np := nodePath(s.Name, node.Id)

	// check existing lease cache
	e.RLock()
	leaseID, ok := e.leases[np]
	e.RUnlock()

	if !ok {
		// missing lease, check if the key exists
		ctx, cancel := context.WithTimeout(ctx, e.options.Timeout)
		defer cancel()

		// look for the existing key
		rsp, err := e.client.Get(ctx, np, clientv3.WithSerializable())
		if err != nil {
			return err
		}

		// get the existing lease
		for _, kv := range rsp.Kvs {
			if kv.Lease > 0 {
				leaseID = clientv3.LeaseID(kv.Lease)

				// decode the existing node
				srv := decode(kv.Value)
				if srv == nil || len(srv.Nodes) <= 0 {
					continue
				}

				// create hash of service; uint64
				h, err := hash.Hash(srv.Nodes[0], nil)
				if err != nil {
					continue
				}

				// save the info
				e.Lock()
				e.leases[np] = leaseID
				e.register[np] = h
				e.Unlock()

				break
			}
		}
	}

	var leaseNotFound bool

	// renew the lease if it exists
	if leaseID > 0 {
		logger.Tracef(e.serviceCtx, "renewing existing lease for %s %d", s.Name, leaseID)

		if _, err := e.client.KeepAliveOnce(ctx, leaseID); err != nil {
			if err != rpctypes.ErrLeaseNotFound {
				return err
			}

			logger.Tracef(e.serviceCtx, "lease not found for %s %d", s.Name, leaseID)
			// lease not found do register
			leaseNotFound = true
		}
	}

	// create hash of service; uint64
	h, err := hash.Hash(node, nil)
	if err != nil {
		return err
	}

	// get existing hash for the service node
	e.Lock()
	v, ok := e.register[s.Name+node.Id]
	e.Unlock()

	// the service is unchanged, skip registering
	if ok && v == h && !leaseNotFound {
		logger.Tracef(e.serviceCtx, "service %s node %s unchanged skipping registration", s.Name, node.Id)
		return nil
	}

	service := &registry.Service{
		Name:      s.Name,
		Version:   s.Version,
		Metadata:  s.Metadata,
		Endpoints: s.Endpoints,
		Nodes:     []registry.Node{node},
	}

	ctx, cancel := context.WithTimeout(context.Background(), e.options.Timeout)
	defer cancel()

	var lgr *clientv3.LeaseGrantResponse
	if ttl.Seconds() > 0 {
		// get a lease used to expire keys since we have a ttl
		lgr, err = e.client.Grant(ctx, int64(ttl.Seconds()))
		if err != nil {
			return err
		}
	}

	logger.Tracef(e.serviceCtx, "registering %s id %s with lease %v and leaseID %v and ttl %v", service.Name, node.Id, lgr, lgr.ID, ttl)
	// create an entry for the node
	if lgr != nil {
		_, err = e.client.Put(ctx, np, encode(service), clientv3.WithLease(lgr.ID))
	} else {
		_, err = e.client.Put(ctx, np, encode(service))
	}
	if err != nil {
		return err
	}

	e.Lock()
	// save our hash of the service
	e.register[np] = h
	// save our leaseID of the service
	if lgr != nil {
		e.leases[np] = lgr.ID
	}
	e.Unlock()

	return nil
}

func encode(s *registry.Service) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func decode(ds []byte) *registry.Service {
	var s *registry.Service
	json.Unmarshal(ds, &s)
	return s
}

func nodePath(s, id string) string {
	service := strings.ReplaceAll(s, "/", "-")
	node := strings.ReplaceAll(id, "/", "-")
	return path.Join(prefix, service, node)
}

func servicePath(s string) string {
	return path.Join(prefix, strings.ReplaceAll(s, "/", "-"))
}
