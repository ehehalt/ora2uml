package ora2uml

import (
	"testing"
)

func TestModelTableFullName(t *testing.T) {
	var table = ModelTable{TableName: "tableName"}

	if table.FullName() != "tableName" {
		t.Errorf("table.FullName() should be 'tableName', but was '%v'", table.FullName())
	}

	table.Owner = "sys"

	if table.FullName() != "sys.tableName" {
		t.Errorf("table.FullName() should be 'sys.tableName', but was '%v'", table.FullName())
	}
}
