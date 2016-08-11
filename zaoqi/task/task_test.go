package task

import (
	"fmt"
	"testing"
	"time"
)

type err struct {
	Value string
}

func (this err) Error() string {
	return this.Value
}
func TestTask(t *testing.T) {
	a := []interface{}{0, "a"}
	b := []interface{}{1, "b"}
	er := err{"error"}
	fn := func(a []interface{}) (r []interface{}, e error) {
		r = a
		e = er
		time.Sleep(time.Second * 1)
		return
	}
	tk1 := NewTask(fn, a...)
	tk2 := NewTask(fn, b...)
	r1, e1 := tk1.Get()
	r2, e2 := tk2.Get()
	if !eq(r1, a) || !eq(r2, b) || e1 != er || e2 != er {
		t.Fatal("UNKNOWN")
	}
	p(r1, e1.Error(), r2, e2.Error(), "1.00s")
}
func eq(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func p(str ...interface{}) {
	for _, s := range str {
		fmt.Println(s)
	}
}
