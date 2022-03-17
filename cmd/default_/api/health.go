package api

import (
	"runtime"
	"time"

	"github.com/coreservice-io/UJob"
	"github.com/goodboy3/referenceMem/plugin/echoServer"
	"github.com/goodboy3/referenceMem/plugin/reference"
	"github.com/goodboy3/referenceMem/src/global"
	"github.com/labstack/echo/v4"
)

// @Summary      health check
// @Description  health check
// @Tags         health
// @Produce      json
// @Success      200 {object} echoServer.RespBody{data=int64} "result"
// @Router       /api/health [get]
func healthCheck(ctx echo.Context) error {
	return echoServer.SuccessResp(ctx, 1, time.Now().Unix(), "")
}

// @Summary      status check
// @Description  status check
// @Tags         health
// @Produce      json
// @Success      200 {object} echoServer.RespBody "result"
// @Router       /api/status [get]
func statusHandler(ctx echo.Context) error {

	var m runtime.MemStats
	runtime.ReadMemStats(&m) //http://www.iargs.cn/?p=62

	s := Status{
		StartTime:     time.Unix(global.StartTime, 0).UTC().Format(time.RFC3339),
		TotalKeyCount: reference.GetInstance().GetLen(),
		CurrentID:     global.ID,
		Alloc_MB:      m.Alloc / (1024 * 1024),
		Sys_MB:        m.Sys / (1024 * 1024),
		MemStats:      m,
		JobInfo:       *global.UJobInfo,
	}

	return echoServer.SuccessResp(ctx, 1, s, "")
}

func config_health(httpServer *echoServer.EchoServer) {
	//health
	httpServer.GET("/api/health", healthCheck)
	httpServer.GET("/api/status", statusHandler)
}

type Status struct {
	StartTime     string
	TotalKeyCount int64
	CurrentID     uint64
	Alloc_MB      uint64
	Sys_MB        uint64
	MemStats      runtime.MemStats
	JobInfo       UJob.Job
}
