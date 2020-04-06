package ora2uml

import (
	"fmt"
)

type ModelColumn struct {
}

type ModelTable struct {
	Owner     string
	TableName string
	Comments  string
	Columns   []ModelColumn
}

type Model struct {
	Tables []ModelTable
}

func (table ModelTable) FullName() string {
	if len(table.Owner) > 0 {
		return fmt.Sprintf("%s.%s", table.Owner, table.TableName)
	}
	return table.TableName
}

func (model *Model) AddTable(table ModelTable) {
	model.Tables = append(model.Tables, table)
}
