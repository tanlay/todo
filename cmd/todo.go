package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tanlay/todo/config"
	"github.com/tanlay/todo/pkg/constant"
	"github.com/tanlay/todo/pkg/lib/db"
	"github.com/tanlay/todo/pkg/model"
	"github.com/tanlay/todo/pkg/service"
	"go.uber.org/zap"
)

func CreateCMD() *cobra.Command {
	var (
		task     string
		category string
	)
	Cmd := &cobra.Command{
		Use:   "create",
		Short: "创建todo项",
		Long:  "创建todo项",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(config config.Config) error{
				SetupGlobalDB,
			}
			for _, task := range tasks {
				if err := task(*config.C); err != nil {
					zap.L().Error("setup err: ", zap.Error(err))
					return err
				}
			}
			svc := service.NewTodoServiceImpl(db.GlobalDB)
			req := model.NewCreateToDoRequest()
			req.Task = task
			req.Category = category
			return svc.CreateTodo(context.Background(), req)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.LoadConfigFromToml(cfgFile)
		},
	}
	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "配置文件路径")
	Cmd.PersistentFlags().StringVarP(&task, "task", "t", "", "Todo任务名")
	Cmd.PersistentFlags().StringVarP(&category, "category", "g", "default", "Todo任务分类")
	return Cmd
}

func DescribeCMD() *cobra.Command {
	var (
		id int
	)
	Cmd := &cobra.Command{
		Use:   "describe",
		Short: "通过ID查询todo项",
		Long:  "通过ID查询todo项",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(config config.Config) error{
				SetupGlobalDB,
			}
			for _, task := range tasks {
				if err := task(*config.C); err != nil {
					zap.L().Error("setup err: ", zap.Error(err))
					return err
				}
			}
			svc := service.NewTodoServiceImpl(db.GlobalDB)
			req := model.NewDescribeToDoRequest(id)

			return svc.DescribeToDo(context.Background(), req)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.LoadConfigFromToml(cfgFile)
		},
	}
	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "配置文件路径")
	Cmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Todo任务ID")
	return Cmd
}

func QueryCMD() *cobra.Command {
	var (
		keyword  string
		pageSize int
		pageNum  int
	)
	Cmd := &cobra.Command{
		Use:   "query",
		Short: "查询todo列表",
		Long:  "查询todo列表",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(config config.Config) error{
				SetupGlobalDB,
			}
			for _, task := range tasks {
				if err := task(*config.C); err != nil {
					zap.L().Error("setup err: ", zap.Error(err))
					return err
				}
			}
			svc := service.NewTodoServiceImpl(db.GlobalDB)
			req := model.NewQueryToDoRequest()
			req.Keyword = keyword
			req.PageSize = pageSize
			req.PageNum = pageNum

			return svc.QueryToDo(context.Background(), req)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.LoadConfigFromToml(cfgFile)
		},
	}
	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "配置文件路径")
	Cmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "", "查询关键字")
	Cmd.PersistentFlags().IntVarP(&pageSize, "page_size", "s", 10, "每页的大小")
	Cmd.PersistentFlags().IntVarP(&pageNum, "page_num", "n", 1, "查询第几页")
	return Cmd
}

func UpdateCMD() *cobra.Command {
	var (
		task     string
		category string
		id       int
	)
	Cmd := &cobra.Command{
		Use:   "update",
		Short: "更新todo",
		Long:  "更新todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(config config.Config) error{
				SetupGlobalDB,
			}
			for _, task := range tasks {
				if err := task(*config.C); err != nil {
					zap.L().Error("setup err: ", zap.Error(err))
					return err
				}
			}
			svc := service.NewTodoServiceImpl(db.GlobalDB)
			req := model.NewUpdateToDoRequest(id)
			req.Task = task
			req.Category = category

			return svc.UpdateToDo(context.Background(), req)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.LoadConfigFromToml(cfgFile)
		},
	}
	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "配置文件路径")
	Cmd.PersistentFlags().StringVarP(&task, "task", "t", "", "Todo任务名")
	Cmd.PersistentFlags().StringVarP(&category, "category", "g", "default", "Todo任务分类")
	Cmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Todo任务ID")
	return Cmd
}

func StatusCMD() *cobra.Command {
	var (
		id     int
		status int
	)
	Cmd := &cobra.Command{
		Use:   "status",
		Short: "更新todo状态",
		Long:  "更新todo状态",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(config config.Config) error{
				SetupGlobalDB,
			}
			for _, task := range tasks {
				if err := task(*config.C); err != nil {
					zap.L().Error("setup err: ", zap.Error(err))
					return err
				}
			}
			svc := service.NewTodoServiceImpl(db.GlobalDB)
			req := model.NewUpdateToDoStatusRequest(id, constant.Status(status))

			return svc.UpdateToDoStatus(context.Background(), req)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.LoadConfigFromToml(cfgFile)
		},
	}
	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "配置文件路径")
	Cmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Todo任务ID")
	Cmd.PersistentFlags().IntVarP(&status, "status", "s", 1, "Todo任务是否完成")
	return Cmd
}
