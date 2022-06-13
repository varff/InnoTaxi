package models

type Analyst struct {
	username, password string
}

func LoginAnalyst(uname string, pass string) {
	//TODO connection check and return token
}
func CheckOrdersByMonth(month int8) []Order {
	var orderList []Order
	//TODO get orders from orders connection(by month)
	return orderList
}

func CheckOrdersByDay(day int8) []Order {
	var orderList []Order
	//TODO get orders from connection(by month)
	return orderList
}

func CheckDriverRate(username string) int8 {
	var driver Driver
	//TODO get driver from driver connection(by username)
	return driver.rate
}

func GetUserRate(username string) User {
	var user User
	//TODO get user from user connection(by username)
	return user
}
