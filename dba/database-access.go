package dba

import (
	"database/sql"
	"fmt"
	"github.com/zceuiv/gorest/configure"
	"github.com/zceuiv/gorest/godbc"
)

var (
	__DbOpt *godbc.Options
)

func Option() *godbc.Options {
	if __DbOpt == nil {
		__DbOpt = configure.SingletonConfig().DbOptions()
	}
	return __DbOpt
}

func Query(query string) (cols []string, result []map[string]*string) {
	db, err := sql.Open(godbc.SqlName(Option()), godbc.ConnStr(Option()))
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	cols, _ = rows.Columns()
	//fmt.Println(cols)

	valuePointRow := make([]interface{}, len(cols))

	for rows.Next() {
		valueRow := make([]string, len(cols))
		retRow := make(map[string]*string)
		for i := 0; i < len(cols); i++ {
			valuePointRow[i] = &valueRow[i]
			retRow[cols[i]] = &valueRow[i]
		}
		rows.Scan(valuePointRow...)
		result = append(result, retRow)
	}

	return cols, result
}

func CreateOrReplace(query string) (id int64) {
	db, err := sql.Open("postgres", godbc.ConnStr(Option()))
	defer db.Close()
	st, err := db.Prepare(query)
	rows, err := st.Query()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func Delete(query string) {
	db, err := sql.Open("postgres", godbc.ConnStr(Option()))
	defer db.Close()
	st, err := db.Prepare(query)
	ret, err := st.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}
	ret.RowsAffected()
}
