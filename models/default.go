package models

import (
	"fmt"
	"os"
	"strconv"

	"github.com/astaxie/beego/orm"
	/* _ "github.com/go-sql-driver/mysql"
	   _ "github.com/lib/pq" */
	_ "github.com/mattn/go-sqlite3"
)

var (
	o            orm.Ormer
	DEBUG        bool
	AUTO_MIGRATE bool
)

type MyModel interface {
	Valid() bool
	TableName() string
	String() string
}

func init() {
	fmt.Println("Initialing model's defaults")
	dbHome := os.Getenv("DB_HOME")
	if dbHome == "" {
		dbHome = os.Getenv("HOME")
	}

	AUTO_MIGRATE, _ = strconv.ParseBool(os.Getenv("MIGRATE"))

	fmt.Printf("Connecting to \"%s\"\n", dbHome+"/socdaalka.db")

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", dbHome+"/socdaalka.db")

	orm.RegisterModel(new(Post), new(User), new(Profile), new(Image), new(Category))

	if AUTO_MIGRATE {
		name := "default"                          // Database alias.
		force := true                              // Drop table and re-create.
		verbose := true                            // Print log
		err := orm.RunSyncdb(name, force, verbose) // Sync with database
		if err != nil {
			Verbose(err.Error())
		}
	}
	o = orm.NewOrm()
	o.Raw("PRAGMA foreign_keys = ON")
}

func Verbose(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format+"\n", args...)
	}
}
