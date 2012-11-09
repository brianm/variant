package variant

import (
	"fmt"
	"expvar"
	"sync"
	"container/ring")

type SimpleMovingAverage struct {
	size int
	mutex *sync.Mutex
	values *ring.Ring
}

func NewSimpleMovingAverage(name string, size int) *SimpleMovingAverage {
	sma := new(SimpleMovingAverage)
	sma.size = size
	sma.mutex = new(sync.Mutex)
	sma.values = ring.New(size)
	
	if name != "" {
		expvar.Publish(name, sma)
	}
	return sma
}

func (s *SimpleMovingAverage) String() string {
	return fmt.Sprintf("%f", s.Average())
}

func (s *SimpleMovingAverage) Update(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.values.Value = val
	s.values = s.values.Next()
}

func (s *SimpleMovingAverage) Average() float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	var sum float64 = 0.0
	var cnt int = 0

	s.values.Do(func(val interface{}) {
		if val != nil {		
			cnt++
			sum = sum + val.(float64)
		} 
	})
	return sum / float64(cnt)
}
