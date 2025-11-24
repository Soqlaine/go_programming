package main

import "fmt"

func main() {

	friends := []string{"Marry", "John", "Paul", "Diana"}

	yourFriends := []string{}

	yourFriends = append(yourFriends, friends...)

	// fmt.Println(yourFriends)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)
}
