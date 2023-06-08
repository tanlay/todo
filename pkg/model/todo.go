package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/tanlay/todo/pkg/constant"
	"time"
)

var (
	// 使用参考: https://raw.githubusercontent.com/go-playground/validator/master/_examples/simple/main.go
	validate = validator.New()
)

func NewCreateToDo(req *CreateToDoRequest) *ToDo {
	return &ToDo{
		CreateToDoRequest: req,
		Status:            constant.StatusDoing,
		CreateAt:          time.Now().Unix(),
	}
}

type ToDo struct {
	Id int `json:"id" db:"id"`
	*CreateToDoRequest
	Status      constant.Status `json:"status" db:"status"`             //是否完成
	CreateAt    int64           `json:"create_at" db:"create_at"`       //创建时间
	CompletedAt int64           `json:"completed_at" db:"completed_at"` //完成时间
}

func (req *ToDo) String() string {
	jd, _ := json.Marshal(req)
	return string(jd)
}

func (req *ToDoSet) String() string {
	jd, _ := json.Marshal(req)
	return string(jd)
}

type ToDoSet struct {
	Total   int64   `json:"total"`
	NoTotal int64   `json:"no_total"`
	Items   []*ToDo `json:"items"`
}

func NewToDoSet() *ToDoSet {
	return &ToDoSet{
		Items: []*ToDo{},
	}
}

func NewCreateToDoRequest() *CreateToDoRequest {
	return &CreateToDoRequest{}
}

type CreateToDoRequest struct {
	Task     string `json:"task" db:"task" validate:"required"`         //任务名
	Category string `json:"category" db:"category" validate:"required"` //分类
}

func (req *CreateToDoRequest) Validate() error {
	return validate.Struct(req)
}

type DescribeToDoRequest struct {
	Id int
}

func NewDescribeToDoRequest(id int) *DescribeToDoRequest {
	return &DescribeToDoRequest{
		Id: id,
	}
}

type QueryToDoRequest struct {
	Keyword  string
	PageNum  int
	PageSize int
}

func (req *QueryToDoRequest) Offset() int {
	return (req.PageNum - 1) * req.PageSize
}

func NewQueryToDoRequest() *QueryToDoRequest {
	return &QueryToDoRequest{
		PageSize: 10,
		PageNum:  1,
	}
}

type UpdateToDoRequest struct {
	Id int
	*CreateToDoRequest
}

func NewUpdateToDoRequest(id int) *UpdateToDoRequest {
	return &UpdateToDoRequest{
		Id:                id,
		CreateToDoRequest: NewCreateToDoRequest(),
	}
}

type UpdateToDoStatusRequest struct {
	Id     int
	Status constant.Status
}

func NewUpdateToDoStatusRequest(id int, status constant.Status) *UpdateToDoStatusRequest {
	return &UpdateToDoStatusRequest{
		Id:     id,
		Status: status,
	}
}
