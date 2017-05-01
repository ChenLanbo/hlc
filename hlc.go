// Hybrid logical clock implementation
// https://www.cse.buffalo.edu/tech-reports/2014-04.pdf
package hlc

import (
    "sync"
    "time"
)

// Global variables
var lt int64 = 0  // local physical time
var ct int64 = 0  // causality time
var mu sync.Mutex

func GetHLCFromLocalEvent() *HybridLogicalClock {
    mu.Lock()
    defer mu.Unlock()
    oldLt := lt
    lt = time.Now().UnixNano()
    if lt < oldLt {
        lt = oldLt
    }

    if lt == oldLt {
        ct++
    } else {
        ct = 0
    }

    return &HybridLogicalClock{LocalTime:lt, CausalTime:ct}
}

func GetHLCFromRemoteEvent(c *HybridLogicalClock) *HybridLogicalClock {
    mu.Lock()
    defer mu.Unlock()
    oldLt := lt
    lt = time.Now().UnixNano()
    if lt < oldLt {
        lt = oldLt
    }
    if lt < c.GetLocalTime() {
        lt = c.GetLocalTime()
    }

    if lt == oldLt && lt == c.GetLocalTime() {
        if ct < c.GetCausalTime() {
            ct = c.GetCausalTime()
        }
        ct++
    } else if lt == oldLt {
        ct++
    } else if lt == c.GetLocalTime() {
        ct = c.GetCausalTime() + 1
    } else {
        ct = 0
    }

    return &HybridLogicalClock{LocalTime:lt, CausalTime:ct}
}

func Equal(c1, c2 *HybridLogicalClock) bool {
    if c1.GetLocalTime() != c2.GetLocalTime() {
        return false
    }
    if c1.GetCausalTime() != c2.GetCausalTime() {
        return false
    }
    return true
}

func Greater(c1, c2 *HybridLogicalClock) bool {
    if c1.GetLocalTime() < c2.GetLocalTime() {
        return false
    }
    if c1.GetLocalTime() > c2.GetLocalTime() {
        return true
    }
    return c1.GetCausalTime() > c2.GetCausalTime()
}

func Smaller(c1, c2 *HybridLogicalClock) bool {
    if c1.GetLocalTime() > c2.GetLocalTime() {
        return false
    }
    if c1.GetLocalTime() < c2.GetLocalTime() {
        return true
    }
    return c1.GetCausalTime() < c2.GetCausalTime()
}
