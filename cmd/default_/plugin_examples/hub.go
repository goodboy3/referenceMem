package examples

import (
	uhub "github.com/coreservice-io/UHub"
	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/plugin/hub"
)

//hub example
const testKind uhub.Kind = 1

type testEvent string

func (e testEvent) Kind() uhub.Kind {
	return testKind
}
func Hub_run() {
	hub.GetInstance().Subscribe(testKind, func(e uhub.Event) {
		basic.Logger.Debugln("hub callback")
		basic.Logger.Debugln(string(e.(testEvent)))
	})
	hub.GetInstance().Publish(testEvent("hub message"))
}
