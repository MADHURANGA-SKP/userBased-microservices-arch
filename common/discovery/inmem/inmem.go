package inmem

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Registry struct {
	sync.RWMutex
	addr map[string]map[string]*serviceInstance
}

type serviceInstance struct {
	hostPort string
	lastActive time.Time
}

func NewRegistry() *Registry {
	return &Registry{addr: map[string]map[string]*serviceInstance{}}
}

func (r *Registry) Register( ctx  context.Context, instanceID, serviceName, hostPort string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.addr[serviceName]; ok {
		r.addr[serviceName] = map[string]*serviceInstance{}
	}

	r.addr[serviceName][instanceID] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}

	return nil
}

func (r *Registry) DeRegister(ctx context.Context, instanceID, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.addr[serviceName]; !ok {
		return nil
	}

	delete(r.addr[serviceName], instanceID)

	return nil
}

func (r *Registry) HealthCheck(ctx context.Context, instanceID, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.addr[serviceName]; !ok {
		return errors.New("service not registered yet")
	}

	if _, ok := r.addr[serviceName][instanceID]; !ok {
		return errors.New("service instance not registered yet")
	}

	r.addr[serviceName][instanceID].lastActive =time.Now()

	return nil
}

func (r *Registry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()

	if len(r.addr[serviceName]) == 0 {
		return nil, errors.New("no service address found")
	}

	var res []string
	for _, i := range r.addr[serviceName] {
		res = append(res, i.hostPort)
	}

	return res, nil
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()

	if len(r.addr[serviceName]) == 0 {
		return nil, errors.New("no service address found")
	}

	var res []string
	for _, i := range r.addr[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)){
			continue
		}
		res = append(res, i.hostPort)
	}

	return res, nil
}


