package main

import (
	"fmt"

	"github.com/go_structures/users"
)

// Id: 1, name: martha, age: 20
func main() {

	martha := users.User{Id: 1, Name: "Martha", Age: 20}
	pedro := users.User{Id: 2, Name: "Pedro", Age: 21}
	john := users.User{Id: 3, Name: "John", Age: 22}
	jane := users.User{Id: 4, Name: "Jane", Age: 23}

	martha.AddFriend(pedro)
	martha.AddFriend(john)
	martha.AddFriend(jane)

	martha.ListFriends()

	martha.RemoveFriend(john.Id)
	fmt.Print("======================\n")
	martha.ListFriends()

}
