package g

import (
	"log"
	"os"

	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var (
	APP_ID string
)

func InitEnv() {
	APP_ID = os.Getenv("APP_ID")
	if APP_ID == "" {
		APP_ID = "1234567890"
	}
	globalSessions, err := session.NewManager("redis", &session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600, ProviderConfig: "127.0.0.1:6379"})
	if err != nil {
		log.Fatalln(err)
	}
	go globalSessions.GC()
}
