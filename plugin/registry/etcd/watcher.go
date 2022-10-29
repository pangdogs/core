package etcd

import (
	"context"
	"errors"
	"github.com/pangdogs/galaxy/plugin/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type _EtcdWatcher struct {
	stopChan  chan bool
	watchChan clientv3.WatchChan
	client    *clientv3.Client
	timeout   time.Duration
}

func newEtcdWatcher(ctx context.Context, r *_EtcdRegistry, timeout time.Duration, serviceName string) (registry.Watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan bool, 1)

	go func() {
		<-stop
		cancel()
	}()

	watchPath := prefix
	if serviceName != "" {
		watchPath = servicePath(serviceName) + "/"
	}

	return &_EtcdWatcher{
		stopChan:  stop,
		watchChan: r.client.Watch(ctx, watchPath, clientv3.WithPrefix(), clientv3.WithPrevKV()),
		client:    r.client,
		timeout:   timeout,
	}, nil
}

func (ew *_EtcdWatcher) Next() (*registry.Result, error) {
	for watchRsp := range ew.watchChan {
		if watchRsp.Err() != nil {
			return nil, watchRsp.Err()
		}
		if watchRsp.Canceled {
			return nil, errors.New("could not get next")
		}
		for _, ev := range watchRsp.Events {
			service := decode(ev.Kv.Value)
			var action string

			switch ev.Type {
			case clientv3.EventTypePut:
				if ev.IsCreate() {
					action = "create"
				} else if ev.IsModify() {
					action = "update"
				}
			case clientv3.EventTypeDelete:
				action = "delete"

				// get service from prevKv
				service = decode(ev.PrevKv.Value)
			}

			if service == nil {
				continue
			}
			return &registry.Result{
				Action:  action,
				Service: service,
			}, nil
		}
	}
	return nil, errors.New("could not get next")
}

func (ew *_EtcdWatcher) Stop() {
	select {
	case <-ew.stopChan:
		return
	default:
		close(ew.stopChan)
	}
}
