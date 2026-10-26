package selfexit

import (
	"os"
	"sync"
	"time"
)

type Exit struct {
	mutex sync.RWMutex
	timer *time.Timer
	dead  time.Time
}

func (e *Exit) Exit(du time.Duration, code int) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.timer != nil {
		e.timer.Stop()
	}
	e.dead = time.Now().Add(du)
	e.timer = time.AfterFunc(du, func() { os.Exit(code) })
}

func (e *Exit) Cancel() time.Time {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	var dead time.Time
	if e.timer != nil {
		e.timer.Stop()
		dead = e.dead
		e.timer = nil
		e.dead = time.Time{}
	}

	return dead
}

func (e *Exit) Deadline() time.Time {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.dead
}
