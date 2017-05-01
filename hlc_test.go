package hlc

import (
    "testing"
)

func TestComparison(t *testing.T) {
    t1 := &HybridLogicalClock{LocalTime:1, CausalTime:0}
    t2 := &HybridLogicalClock{LocalTime:1, CausalTime:0}
    t3 := &HybridLogicalClock{LocalTime:2, CausalTime:0}
    t4 := &HybridLogicalClock{LocalTime:2, CausalTime:1}

    if !Equal(t1, t2) {
        t.Error("Should be equal", t1, t2)
    }

    if !Greater(t3, t1) {
        t.Error("Should be greater", t3, t1)
    }

    if !Smaller(t3, t4) {
        t.Error("Should be smaller", t3, t4)
    }
}
