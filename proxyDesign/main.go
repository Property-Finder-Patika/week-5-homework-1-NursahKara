package main

import "fmt"

type Licence struct {
	Users int
}

type LicenceProxy struct {
	users *Licence
}

func (c *LicenceProxy) Control() {
	c.users.Users = 18

	if c.users.Users <= 20 {
		fmt.Println("You can use this licence.")
	} else {
		fmt.Println("NOt enought capacity!")
	}
}

func main() {
	var object LicenceProxy = LicenceProxy{}

	object.Control()
}
