package serve

import (
	"framework/app/task"
	"framework/ay"
	"framework/middleware"
	"framework/routers"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	r *gin.Engine
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "开启服务",

		Run: func(cmd *cobra.Command, args []string) {
			var err error
			ay.Init()

			// 定时任务
			go func() {
				c := cron.New()
				c.AddFunc("@every 3m", task.Instance)
				c.Start()
			}()

			// 加载gin
			gin.SetMode(gin.DebugMode)
			r = gin.Default()
			middleware.Instance(r)
			r.StaticFS("/static/", http.Dir("./static"))
			r.StaticFS("/root", http.Dir("./admin"))
			routers.Instance(r)

			err = r.Run(":8080")

			if err != nil {
				ay.Logger.Error(err.Error())
			}
			select {}
		},
	}
	return cmd
}
