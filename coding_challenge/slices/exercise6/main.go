package main

import "fmt"

func main() {

	friends := []string{"Marry", "John", "Paul", "Diana"}

	newFriends := make([]string, len(friends))

	copy(newFriends, friends)

	newFriends[0] = "Dan"

	fmt.Printf("Friends:%q\nnewFriends: %q\n", friends, newFriends)
}
