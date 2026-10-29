package pubsub

import "sync"

type dispatcher struct {
	mutex sync.RWMutex
	subs  map[Subscriber]struct{}
}

func NewDispatcher() Dispatcher {
	return &dispatcher{subs: make(map[Subscriber]struct{}, 16)}
}

func (dp *dispatcher) Publish(event string, data any) {
	go dp.publish(event, data)
}

func (dp *dispatcher) AddSubscriber(sub Subscriber) {
	if sub == nil {
		return
	}

	dp.mutex.Lock()
	defer dp.mutex.Unlock()

	dp.subs[sub] = struct{}{}
}

func (dp *dispatcher) DelSubscriber(sub Subscriber) {
	if sub == nil {
		return
	}

	dp.mutex.Lock()
	defer dp.mutex.Unlock()

	delete(dp.subs, sub)
}

func (dp *dispatcher) Count() int {
	dp.mutex.RLock()
	defer dp.mutex.RUnlock()

	return len(dp.subs)
}

func (dp *dispatcher) publish(event string, data any) {
	for _, sub := range dp.snapshot() {
		if err := sub.OnEvent(event, data); err != nil {
			dp.DelSubscriber(sub)
		}
	}
}

func (dp *dispatcher) snapshot() []Subscriber {
	dp.mutex.RLock()
	defer dp.mutex.RUnlock()

	subs := make([]Subscriber, 0, len(dp.subs))
	for sub := range dp.subs {
		subs = append(subs, sub)
	}

	return subs
}
