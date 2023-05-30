package serve

import (
	"framework/app/services"
	"framework/ay"
	"framework/middleware"
	"framework/routers"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"log"
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
			// 配置文件
			ay.Yaml = ay.InitConfig()
			err := ay.GetDB()
			if err != nil {
				log.Println(err)
			}
			go ay.WatchConf()

			// 定时任务
			c := cron.New()
			_, err = c.AddFunc("@every 3m", services.Instance)
			if err != nil {
				log.Println(err.Error())
				return
			}
			c.Start()

			// 加载gin
			gin.SetMode(gin.DebugMode)
			r = gin.Default()

			middleware.Instance(r)
			r.StaticFS("/static/", http.Dir("./static"))
			r.StaticFS("/root", http.Dir("./admin"))
			routers.Instance(r)
			err = r.Run(":" + ay.Yaml.GetString("port"))

			if err != nil {
				panic(err.Error())
			}
			select {}
		},
	}
	return cmd
}
