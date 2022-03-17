package examples

import (
	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/configuration"
)

//example get complex config
func ComplexConfig_run() {
	provide_folder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		basic.Logger.Errorln(err)
	}
	for _, v := range provide_folder {
		basic.Logger.Debugln("path:", v.AbsPath, "size:", v.SizeGB)
	}
}
