package main

import "fmt"

type Users struct {
	Status string
}

func (user *Users) UpdateStatus(str string) {
	user.Status = str

}

func main() {
	user := &Users{
		Status: "active",
	}
	fmt.Printf(user.Status)

	user.UpdateStatus("inactive")
	fmt.Println(user.Status)
}
