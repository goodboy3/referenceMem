package examples

import (
	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/plugin/reference"
)

//cache example
func Reference_run() {
	bar1 := "bar1"
	err1 := reference.GetInstance_("ref1").Set("foo1", &bar1, 10)
	if err1 != nil {
		basic.Logger.Errorln(err1)
	}
	v, _, exist := reference.GetInstance_("ref1").Get("foo1")
	if exist {
		basic.Logger.Debugln(*(v.(*string)))
	}

	bar2 := "bar2"
	err2 := reference.GetInstance_("ref1").Set("foo2", &bar2, 10)
	if err2 != nil {
		basic.Logger.Errorln(err2)
	}
	v, _, exist = reference.GetInstance_("ref1").Get("foo2")
	if exist {
		basic.Logger.Debugln(*(v.(*string)))
	}
}
