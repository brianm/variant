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

func Test_smm_90p(t *testing.T) {
	sma := NewSimpleMovingPercentile("", 0.9, 10)

	sma.Update(1)
	sma.Update(2)
	sma.Update(3)
	sma.Update(4)
	sma.Update(5)
	sma.Update(6)
	sma.Update(7)
	sma.Update(8)
	sma.Update(9)
	sma.Update(10)

	avg := sma.Value()

	if avg != 10.0 {
		t.Errorf("expected median of 2.0, got %f", avg)
	}
}
