package main

import (
	user "solid-principle/server/users"
)

type app struct {
	user user.IUsers
}

func main() {
	// app := &app{
	// 	user: user.InitializeUsers(),
	// }
	// singleton pattern
	// for i := 0; i < 30; i++ {
	// 	go dbService.GetInstance()
	// }

	//app.user.SetUserProfile("Kamal Kant", 123, "16th Cross B town")
	// fmt.Println("this is another line", app.user.GetUserName())
	// app.user.GetUserProfile()
	//fmt.Scanln()

	// Prototype Pattern
	//file1 := &p.File{}

}
