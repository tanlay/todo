package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/tanlay/todo/pkg/constant"
	"github.com/tanlay/todo/pkg/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type ToDoDaoInterface interface {
	CreateToDo(ctx context.Context, req *model.CreateToDoRequest) (*model.ToDo, error)
	DescribeToDo(ctx context.Context, req *model.DescribeToDoRequest) (*model.ToDo, error)
	QueryToDo(ctx context.Context, req *model.QueryToDoRequest) (*model.ToDoSet, error)
	UpdateToDo(ctx context.Context, req *model.UpdateToDoRequest) (*model.ToDo, error)
	UpdateToDoStatus(ctx context.Context, req *model.UpdateToDoStatusRequest) (*model.ToDo, error)
}

type TodoDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

var NewTodoDaoImpl = func(db *gorm.DB) ToDoDaoInterface {
	return &TodoDaoImpl{
		db:     db,
		logger: zap.L().Named("dao"),
	}
}

func (t *TodoDaoImpl) DB() *gorm.DB {
	return t.db.Table("todo")
}

func (t *TodoDaoImpl) CreateToDo(ctx context.Context, req *model.CreateToDoRequest) (*model.ToDo, error) {
	//校验对象合法性
	if err := req.Validate(); err != nil {
		t.logger.Error("创建Todo对象校验错误，", zap.Error(err))
		return nil, err
	}

	ins := model.NewCreateToDo(req)
	if err := t.DB().WithContext(ctx).Create(&ins).Error; err != nil {
		t.logger.Error("创建Todo错误，", zap.Error(err))
		return nil, err
	}
	return ins, nil
}

func (t *TodoDaoImpl) DescribeToDo(ctx context.Context, req *model.DescribeToDoRequest) (*model.ToDo, error) {
	ins := model.NewCreateToDo(model.NewCreateToDoRequest())
	if err := t.DB().WithContext(ctx).Where("id=?", req.Id).Find(&ins).Error; err != nil {
		t.logger.Error("获取Todo记录错误，", zap.Error(err))
		return nil, err
	}
	if ins.Id == 0 {
		t.logger.Info(fmt.Sprintf("未查询到Todo记录：%d", req.Id))
		return nil, errors.New(fmt.Sprintf("未查询到Todo记录：%d", req.Id))
	}
	return ins, nil
}

func (t *TodoDaoImpl) QueryToDo(ctx context.Context, req *model.QueryToDoRequest) (*model.ToDoSet, error) {
	set := model.NewToDoSet()
	query := t.DB().WithContext(ctx)
	//支持关键字查询
	if req.Keyword != "" {
		query = query.Where("task like ? or category like ?",
			"%"+req.Keyword+"%",
			"%"+req.Keyword+"%",
		)
	}
	//查询总条数
	if err := query.Count(&set.Total).Error; err != nil {
		t.logger.Error("查询Todo总记录数错误，", zap.Error(err))
		return nil, err
	}
	//查询未完成的条数
	if err := t.DB().WithContext(ctx).Where("status=?", "0").Count(&set.NoTotal).Error; err != nil {
		t.logger.Error("查询Todo未完成的记录数错误，", zap.Error(err))
		return nil, err
	}

	//支持分页
	if err := query.Offset(req.Offset()).Limit(req.PageSize).Order("create_at desc").
		Scan(&set.Items).Error; err != nil {
		t.logger.Error("查询Todo列表错误，", zap.Error(err))
		return nil, err
	}

	return set, nil
}

func (t *TodoDaoImpl) UpdateToDo(ctx context.Context, req *model.UpdateToDoRequest) (*model.ToDo, error) {
	ins, err := t.DescribeToDo(ctx, model.NewDescribeToDoRequest(req.Id))
	if err != nil {
		return nil, err
	}
	//修改
	if err := mergo.MergeWithOverwrite(ins.CreateToDoRequest, req.CreateToDoRequest); err != nil {
		t.logger.Error("更新todo错误", zap.Error(err))
		return nil, err
	}
	//ins.CreateToDoRequest = req.CreateToDoRequest
	//校验对象合法性
	if err := req.CreateToDoRequest.Validate(); err != nil {
		t.logger.Error("更新Todo对象校验错误，", zap.Error(err))
		return nil, err
	}

	//入库
	if err := t.DB().Save(&ins).Error; err != nil {
		t.logger.Error("更新Todo记录错误，", zap.Error(err))
		return nil, err
	}

	return ins, nil
}

func (t *TodoDaoImpl) UpdateToDoStatus(ctx context.Context, req *model.UpdateToDoStatusRequest) (*model.ToDo, error) {
	//判断对象是否存在
	ins, err := t.DescribeToDo(ctx, model.NewDescribeToDoRequest(req.Id))
	if err != nil {
		return nil, err
	}
	//判断状态是预期状态
	if ins.Status == req.Status {
		t.logger.Error(fmt.Sprintf("todo对象状态已是预期的状态: %d", ins.Status))
		return nil, errors.New("todo对象状态已是预期的状态")

	}
	//修改状态
	ins.Status = req.Status
	if ins.Status == constant.StatusDone {
		ins.CompletedAt = time.Now().Unix()
	} else if ins.Status == constant.StatusDoing {
		ins.CompletedAt = 0 //如果状态设置为doing，清除时间
	} else {
		t.logger.Error(fmt.Sprintf("不支持的状态: %d", req.Status))
		return nil, errors.New("不支持的状态")
	}
	//入库保存
	//if err := t.DB().WithContext(ctx).Updates(&ins).Error; err != nil { //Updates()不更新字段为0的字段
	if err := t.DB().WithContext(ctx).Save(&ins).Error; err != nil {
		t.logger.Error("更新todo对象状态失败", zap.Error(err))
		return nil, errors.New("更新todo对象状态失败")
	}
	return ins, nil
}
