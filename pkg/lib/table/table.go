/*
定义命令行界面展示的格式
*/

package table

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"time"
	"todo/pkg/model"
)

//定义完成和未完成的符号
const (
	doingStatus = "\u2716"
	doneStatus  = "\u2714"
)

//CustomCell 定义表格格式
var (
	cells      [][]*simpletable.Cell
	CustomCell = []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: "#"},
		{Align: simpletable.AlignCenter, Text: "Category"},
		{Align: simpletable.AlignCenter, Text: "Task"},
		{Align: simpletable.AlignCenter, Text: "Status"},
		{Align: simpletable.AlignCenter, Text: "CreatedAt"},
		{Align: simpletable.AlignCenter, Text: "CompletedAt"},
	}
)

var (
	status      string
	createAt    string
	completedAt string
)

func PrintOneToConsole(todo *model.ToDo) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: CustomCell,
	}
	task := todo.Task
	category := todo.Category
	if todo.Status == 0 {
		status = doingStatus
	} else {
		status = doneStatus
	}

	createAt = time.Unix(todo.CreateAt, 0).Format("2006-01-02 15:04:06")
	if todo.CompletedAt != 0 {
		completedAt = time.Unix(todo.CompletedAt, 0).Format("2006-01-02 15:04:06")
	} else {
		completedAt = ""
	}
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", todo.Id)},
		{Text: category},
		{Text: task},
		{Text: status},
		{Text: createAt},
		{Text: completedAt},
	})
	table.Body = &simpletable.Body{
		Cells: cells,
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func PrintToConsole(todos *model.ToDoSet) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: CustomCell,
	}

	for _, todo := range todos.Items {
		task := todo.Task
		category := todo.Category
		if todo.Status == 0 {
			status = doingStatus
		} else {
			status = doneStatus
		}

		createAt = time.Unix(todo.CreateAt, 0).Format("2006-01-02 15:04:06")
		if todo.CompletedAt != 0 {
			completedAt = time.Unix(todo.CompletedAt, 0).Format("2006-01-02 15:04:06")
		} else {
			completedAt = ""
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", todo.Id)},
			{Text: category},
			{Text: task},
			{Text: status},
			{Text: createAt},
			{Text: completedAt},
		})
	}
	table.Body = &simpletable.Body{
		Cells: cells,
	}
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: ""},
			{Align: simpletable.AlignLeft, Span: 5, Text: fmt.Sprintf("一共有%d条todo，还有%d条Todo未完成", todos.Total, todos.NoTotal)},
		},
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
