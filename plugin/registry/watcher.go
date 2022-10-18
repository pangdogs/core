package registry

type Watcher interface {
	// Next is a blocking call
	Next() (*Result, error)
	Stop()
}
