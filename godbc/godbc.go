package godbc

import (
	_ "database/sql"
	_ "database/sql/driver"
	"fmt"
	_ "github.com/zceuiv/mysql"
	_ "github.com/zceuiv/pq"
	"strings"
)

type Options struct {
	SqlName  string
	Addr     string
	Database string
	User     string
	Password string
}

var (
	__SqlName        string
	__ConnStr        string
	__MapNameConnStr map[string](func(opt *Options) string)
)

/*
type pqDrv struct{}


func (d *pqDrv) Open(name string) (driver.Conn, error) {
	return pq.Open(name)
}

func init() {
	sql.Register("postgres", &pqDrv{})
}
*/
func init() {
	__MapNameConnStr = make(map[string](func(opt *Options) string))
	__MapNameConnStr["postgres"] = pqConnStr
	__MapNameConnStr["mysql"] = mysqlConnStr
}

func SqlName(opt *Options) string {
	if __SqlName == "" {
		__SqlName = opt.SqlName
	}
	return __SqlName
}

func ConnStr(opt *Options) string {
	return __MapNameConnStr[SqlName(opt)](opt)
}

func pqConnStr(opt *Options) string {
	if __ConnStr == "" {
		addrs := strings.Split(opt.Addr, ":")
		__ConnStr += "host=" + addrs[0] + " "
		__ConnStr += "port=" + addrs[1] + " "
		__ConnStr += "dbname=" + opt.Database + " "
		__ConnStr += "user=" + opt.User + " "
		__ConnStr += "password=" + opt.Password + " "
		__ConnStr += "sslmode=disable"
	}
	return __ConnStr
}

func mysqlConnStr(opt *Options) string {
	if __ConnStr == "" {
		__ConnStr += fmt.Sprintf("%s:%s@tcp(%s)/%s", opt.User, opt.Password, opt.Addr, opt.Database)
	}
	return __ConnStr
}
