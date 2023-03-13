package users

import "fmt"

type User struct {
	Id      int
	Name    string
	Age     int
	friends []User
}

// AddFriend Funcion que agrega un nuevo amigo y tiene manejo de PUNTERO que es para el manejo de
// persistir en la structura User
func (u *User) AddFriend(friend User) {
	u.friends = append(u.friends, friend)
}

func (u User) ListFriends() {
	for i, f := range u.friends {
		fmt.Printf("%d, %s\n", i+1, f.Name)
	}
}

func (u *User) RemoveFriend(Id int) {
	index := u.findFriend(Id)
	if index < 0 {
		return
	}

	// Se elimina todo lo que tenemos a la derecha y a la izquierda y los agregamos a un nuevo arreglo
	u.friends = append(u.friends[:index], u.friends[index+1:]...)
}

func (u User) findFriend(Id int) int {
	for i, f := range u.friends {
		if f.Id == Id {
			return i
		}
	}
	return -1
}
