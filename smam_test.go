package variant

import (
	"testing"
)

func Test_smm_full(t *testing.T) {
	sma := NewSimpleMovingMedian("", 3)
	sma.Update(1)
	sma.Update(2)
	sma.Update(10)
	avg := sma.Value()
	if avg != 2.0 {
		t.Errorf("expected median of 2.0, got %f", avg)
	}
}

func Test_smm_empty(t *testing.T) {
	sma := NewSimpleMovingMedian("", 3)
	avg := sma.Value()

	if avg != 0.0 {
		t.Errorf("expected median of 2.0, got %f", avg)
	}
}
