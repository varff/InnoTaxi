package connection

import (
	"InnoTaxi/pkg/helper"
	"fmt"
)

func UserConString() string {
	UserID := helper.GetEnvDefault("USERID", "user")
	Pass := helper.GetEnvDefault("USERPASS", "secret")
	Port := helper.GetEnvDefault("USERPORT", "5432")
	Db := helper.GetEnvDefault("USERDB", "postgres")
	Host := helper.GetEnvDefault("USERHOSTNAME", "localhost")

	return fmt.Sprintf("user=" + UserID + " password=" + Pass + " host=" + Host + " port=" + Port + " database=" + Db)
}
