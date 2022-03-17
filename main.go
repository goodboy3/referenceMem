package main

import (
	"os"

	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/cmd"
)

func main() {
	basic.InitLogger()

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Panicln(errRun)
	}
}
