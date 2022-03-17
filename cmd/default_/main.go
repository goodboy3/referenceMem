package default_

import (
	"strconv"
	"time"

	"github.com/coreservice-io/UJob"
	"github.com/fatih/color"
	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/cmd/default_/api"
	"github.com/goodboy3/referenceMem/configuration"
	"github.com/goodboy3/referenceMem/plugin/echoServer"
	"github.com/goodboy3/referenceMem/plugin/reference"
	"github.com/goodboy3/referenceMem/src/global"
	"github.com/goodboy3/referenceMem/tools/errors"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	color.Green(basic.Logo)
	//basic.Logger.Infoln("hello world , this cli app")

	//ini components and run example
	initComponent()

	//defer func() {
	//	//global.ReleaseResources()
	//}()

	global.StartTime = time.Now().Unix()
	global.ID = 0
	RunRefJob()

	start_http_sever()
}

//httpServer example
func start_http_sever() {
	httpServer := echoServer.GetInstance()
	api.DeclareApi(httpServer)
	http_api, _ := configuration.Config.GetBool("http_api", false)
	if http_api {
		api.ConfigApi(httpServer)
	} else {
		httpServer.StaticWeb()
	}
	httpServer.Start()
}

type ExampleUserModel struct {
	ID      uint64
	Status  string
	Name    string
	Email   string
	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}

func RunRefJob() {
	global.UJobInfo = UJob.Start("Reference memory test job", func() {
		for i := 0; i < 10000; i++ {
			global.ID++
			u := &ExampleUserModel{
				ID:      global.ID,
				Status:  "normal",
				Name:    "Tom",
				Email:   "tom@email.com",
				Updated: time.Now().Unix(),
				Created: time.Now().Unix(),
			}
			reference.GetInstance().Set(strconv.FormatUint(u.ID, 10), u, 60)
		}
	}, errors.PanicHandler, 1, UJob.TYPE_PANIC_REDO, nil, nil)
}
