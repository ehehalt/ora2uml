package ora2uml

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/godror/godror"
)

const (
	sqlAllTables = `
	select 
		owner, table_name
	from
		all_tables
	`
)

func readTablesSql(tables []ConfigTable) string {
	sql := "select "
	sql += "t.owner, t.table_name, c.comments "
	sql += "from all_tables t "
	sql += "left outer join all_tab_comments c on t.owner = c.owner and t.table_name = c.table_name "
	sql += "where (t.owner, t.table_name) in ("

	for idx, table := range tables {
		if idx > 0 {
			sql += ","
		}
		sql += fmt.Sprintf("('%s','%s')", table.Owner, table.Name)
	}

	sql += ")"
	return sql
}

func ReadTables(config Config) (Model, error) {
	query := readTablesSql(config.Tables)

	db, err := sql.Open("godror", config.ConnectionString())
	if err != nil {
		fmt.Println("ReadTables:", err)
		os.Exit(0)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ReadTables: Error running query")
		fmt.Println(err)
		os.Exit(0)
	}

	model := &Model{}

	for rows.Next() {
		table := &ModelTable{}
		rows.Scan(&table.Owner, &table.TableName, &table.Comments)
		fmt.Println(table.Owner, table.TableName)
		model.AddTable(*table)
	}

	return *model, nil
}
