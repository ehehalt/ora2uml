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

func (table ModelTable) FullName() string {
	if len(table.Owner) > 0 {
		return fmt.Sprintf("%s.%s", table.Owner, table.TableName)
	}
	return table.TableName
}
