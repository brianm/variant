package variant

import (
	"expvar"
	"fmt"
	"testing"
)

func TestFullBehavior(t *testing.T) {
	sma := NewSimpleMovingAverage("", 3)
	sma.Update(1)
	sma.Update(2)
	sma.Update(3)
	avg := sma.Value()
	if avg != 2.0 {
		t.Errorf("expected abg of 2.0, got %f", avg)
	}
}

func TestPartialBehavior(t *testing.T) {
	sma := NewSimpleMovingAverage("", 3)
	sma.Update(1)
	sma.Update(2)
	avg := sma.Value()
	if avg != 1.5 {
		t.Errorf("expected avg of 1.5, got %f", avg)
	}
}

func TestMovedPastInitialBehavior(t *testing.T) {
	sma := NewSimpleMovingAverage("", 3)
	sma.Update(1)
	sma.Update(2)
	sma.Update(2)
	sma.Update(2)

	avg := sma.Value()
	if avg != 2.0 {
		t.Errorf("expected avg of 2.0, got %f", avg)
	}
}

func TestAsVar(t *testing.T) {
	s := NewSimpleMovingAverage("", 1)
	var _ expvar.Var = s
}

func TestStringMatchesAverage(t *testing.T) {
	s := NewSimpleMovingAverage("", 1)
	s.Update(1.1)
	if st := s.String(); st != fmt.Sprintf("%f", 1.1) {
		t.Errorf("expected '1.1', got %s", st)
	}
}

func TestSliceBehavior(t *testing.T) {
	ary := make([]int, 0)
	ary = append(ary, 1)
	ary = append(ary, 2)
	if len(ary) != 2 {
		t.Errorf("expected len=2, got len=%d", len(ary))
	}
}
