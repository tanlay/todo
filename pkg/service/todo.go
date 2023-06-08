package service

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"todo/pkg/dao"
	"todo/pkg/lib/table"
	"todo/pkg/model"
)

type TodoServiceInterface interface {
	CreateTodo(ctx context.Context, req *model.CreateToDoRequest) (*model.ToDo, error)
	DescribeToDo(ctx context.Context, req *model.DescribeToDoRequest) error
	QueryToDo(ctx context.Context, req *model.QueryToDoRequest) error
	UpdateToDo(ctx context.Context, req *model.UpdateToDoRequest) error
	UpdateToDoStatus(ctx context.Context, req *model.UpdateToDoStatusRequest) error
}

type TodoServiceImpl struct {
	db     *gorm.DB
	logger *zap.Logger
	ctx    context.Context
}

var NewTodoServiceImpl = func(db *gorm.DB) TodoServiceInterface {
	return &TodoServiceImpl{
		db:     db,
		logger: zap.L().Named("service"),
	}
}

func (t *TodoServiceImpl) CreateTodo(ctx context.Context, req *model.CreateToDoRequest) (*model.ToDo, error) {
	impl := dao.NewTodoDaoImpl(t.db)
	ins, err := impl.CreateToDo(ctx, req)
	if err != nil {
		t.logger.Error("创建todo错误", zap.Error(err))
		return nil, err
	}
	t.logger.Info("创建todo成功")
	table.PrintOneToConsole(ins) //终端打印
	return ins, nil
}

func (t *TodoServiceImpl) DescribeToDo(ctx context.Context, req *model.DescribeToDoRequest) error {
	impl := dao.NewTodoDaoImpl(t.db)
	ins, err := impl.DescribeToDo(ctx, req)
	if err != nil {
		t.logger.Error("查询todo项错误", zap.Error(err))
		return err
	}
	t.logger.Info("查询todo项成功")
	table.PrintOneToConsole(ins) //终端打印
	return nil
}

func (t *TodoServiceImpl) QueryToDo(ctx context.Context, req *model.QueryToDoRequest) error {
	impl := dao.NewTodoDaoImpl(t.db)
	set, err := impl.QueryToDo(ctx, req)
	if err != nil {
		t.logger.Error("查询todo列表错误", zap.Error(err))
		return err
	}

	t.logger.Info("查询todo列表成功")
	table.PrintToConsole(set) //终端打印
	return nil
}

func (t *TodoServiceImpl) UpdateToDo(ctx context.Context, req *model.UpdateToDoRequest) error {
	impl := dao.NewTodoDaoImpl(t.db)
	ins, err := impl.UpdateToDo(ctx, req)
	if err != nil {
		t.logger.Error("更新todo项错误", zap.Error(err))

		return err
	}
	t.logger.Info("更新todo项成功")
	table.PrintOneToConsole(ins) //终端打印
	return nil
}

func (t *TodoServiceImpl) UpdateToDoStatus(ctx context.Context, req *model.UpdateToDoStatusRequest) error {
	impl := dao.NewTodoDaoImpl(t.db)
	ins, err := impl.UpdateToDoStatus(ctx, req)
	if err != nil {
		t.logger.Error("更新todo状态错误", zap.Error(err))

		return err
	}
	t.logger.Info("更新todo状态成功")
	table.PrintOneToConsole(ins) //终端打印
	return nil
}
