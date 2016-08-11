
package task

import (
	"runtime"
)

type Task struct{ p *task }
type task struct {
	r []interface{}
	e chan error
}
type Func func(a []interface{}) (r []interface{}, e error)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func NewTask(fn Func, a ...interface{}) Task {
	tk := task{e: make(chan error)}
	go func() {
		r, e := fn(a)
		tk.r = r
		tk.e <- e
	}()
	return Task{&tk}
}
func (self Task) Get() (r []interface{}, e error) {
	e = <-(*self.p).e
	r = (*self.p).r
	self.p = nil
	return
}
