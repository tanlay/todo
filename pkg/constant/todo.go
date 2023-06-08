package constant

import (
	"fmt"
)

type Status int

const (
	StatusDoing Status = iota //正在进行中
	StatusDone                //已完成
)

var (
	StatusMap = map[Status]string{
		StatusDoing: "doing",
		StatusDone:  "done",
	}
)

func (s Status) String() string {
	if v, ok := StatusMap[s]; ok {
		return v
	}
	return fmt.Sprintf("%d", s)
}
