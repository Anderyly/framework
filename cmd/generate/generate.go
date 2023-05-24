package generate

import (
	"framework/ay"
	"framework/dao/model"
	"github.com/spf13/cobra"
	"gorm.io/gen"
	"log"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "开启服务",

		Run: func(cmd *cobra.Command, args []string) {
			runGen()
		},
	}
	return cmd
}

func runGen() {

	// 配置文件
	ay.Yaml = ay.InitConfig()
	err := ay.GetDB()
	if err != nil {
		log.Println(err)
	}
	go ay.WatchConf()

	g := gen.NewGenerator(gen.Config{
		OutPath: "./dao/query",
	})
	list := []interface{}{
		model.Demo{},
	}

	ay.Db.DisableForeignKeyConstraintWhenMigrating = true
	if err := ay.Db.AutoMigrate(list...); err != nil {
		panic(err)
	}
	g.ApplyBasic(list...) // 绑定表
	g.Execute()
}
