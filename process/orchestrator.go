package process

import (
	"fmt"

	"github.com/Julien4218/http-load-tester/common"
	log "github.com/sirupsen/logrus"
)

func Execute(ctx *common.CommandContext) {
	log.Info(fmt.Sprintf("Start execution with rpm:%d, loop:%d, parallel:%d on URL:%s", ctx.RPM, ctx.Loop, ctx.InputConfig.MinParallel, ctx.InputConfig.HttpTest.URL))
}
