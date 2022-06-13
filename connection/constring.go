package connection

import (
	"os"
)

func UserConString() string {
	UserID := os.Getenv("USERID")
	Pass := os.Getenv("USERPASS")
	Port := os.Getenv("USERPORT")
	Db := os.Getenv("USERDB")
	Host := os.Getenv("USERHOSTNAME")

	return string("user=" + UserID + " password=" + Pass + " host=" + Host + " port=" + Port + " database=" + Db)

}

func DriverConString() string {
	DriverID := os.Getenv("DRIVERID")
	Pass := os.Getenv("DRIVERPASS")
	Port := os.Getenv("DRIVERPORT")
	Db := os.Getenv("DRIVERDB")
	Host := os.Getenv("DRIVERHOSTNAME")
	return string("user=" + DriverID + " password=" + Pass + " host=" + Host + " port=" + Port + " database=" + Db)

}

func OrderConString() string {
	OrderID := os.Getenv("ORDERID")
	Pass := os.Getenv("ORDERPASS")
	Port := os.Getenv("ORDERPORT")
	Db := os.Getenv("ORDERDB")
	Host := os.Getenv("ORDERHOSTNAME")
	return string("user=" + OrderID + " password=" + Pass + " host=" + Host + " port=" + Port + " database=" + Db)
}
