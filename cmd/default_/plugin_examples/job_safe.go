package examples

import (
	"time"

	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/tools/errors"

	"github.com/coreservice-io/UJob"
	"github.com/coreservice-io/USafeGo"
)

//job and safego example
func Job_Safeo_run() {
	count := 0
	job := UJob.Start(
		//job process
		"exampleJob",
		func() {
			count++
			basic.Logger.Debugln("Schedule Job running,count", count)
		},
		//onPanic callback
		errors.PanicHandler,
		2,
		// job type
		// UJob.TYPE_PANIC_REDO  auto restart if panic
		// UJob.TYPE_PANIC_RETURN  stop if panic
		UJob.TYPE_PANIC_REDO,
		// check continue callback, the job will stop running if return false
		// the job will keep running if this callback is nil
		func(job *UJob.Job) bool {
			return true
		},
		// onFinish callback
		func(inst *UJob.Job) {
			basic.Logger.Debugln("finish", "cycle", inst.Cycles)
		},
	)

	//safeGo
	USafeGo.Go(
		//process
		func(args ...interface{}) {
			basic.Logger.Debugln("example of USafeGo")
			time.Sleep(10 * time.Second)
			job.SetToCancel()
		},
		//onPanic callback
		errors.PanicHandler)

	for i := 0; i < 10; i++ {
		basic.Logger.Debugln("running")
		time.Sleep(1 * time.Second)
	}
}
