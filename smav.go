/* 

 Package variant provides some (well, at the time of this writing)
 useful implementations of the expvar.Var interface (which is just a
 Stringer).

*/
package variant

import (
	"fmt"
	"expvar"
	"sync"
	"container/ring")

// represents a size bounded simple moving average
// it is thread/goroutine safe
type SimpleMovingAverage struct {
	size int
	mutex *sync.Mutex
	values *ring.Ring
}

// Crate a new simple moving average expvar.Var. It will be
// published under `name` and maintain `size` values for
// calculating the average.
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

// display the value as a string
func (s *SimpleMovingAverage) String() string {
	return fmt.Sprintf("%f", s.Average())
}

// Append a new value to the stat
func (s *SimpleMovingAverage) Update(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.values.Value = val
	s.values = s.values.Next()
}

// obtain the current value
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
