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
type SimpleMovingStat struct {
	size int
	mutex *sync.Mutex
	values *ring.Ring
	updater func(*SimpleMovingStat, float64)
	valuer func(*SimpleMovingStat) float64
}

// Crate a new simple moving average expvar.Var. It will be
// published under `name` and maintain `size` values for
// calculating the average.
func NewSimpleMovingAverage(name string, size int) *SimpleMovingStat {
	sma := new(SimpleMovingStat)
	sma.size = size
	sma.mutex = new(sync.Mutex)
	sma.values = ring.New(size)
	sma.updater = func(s *SimpleMovingStat, val float64) {
		s.values.Value = val
		s.values = s.values.Next()
	}
	
	sma.valuer = func(s *SimpleMovingStat) float64 {
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
	
	if name != "" {
		expvar.Publish(name, sma)
	}
	return sma
}

// display the value as a string
func (s *SimpleMovingStat) String() string {
	return fmt.Sprintf("%f", s.Value())
}

// Append a new value to the stat
func (s *SimpleMovingStat) Update(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.updater(s, val)
}

// obtain the current value
func (s *SimpleMovingStat) Value() float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.valuer(s)
}
