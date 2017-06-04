package main

import (
	"fmt"
	"log"

	"github.com/luladjiev/hnews/item"
	"github.com/luladjiev/hnews/user"
)

func main() {
	item, err := item.Get(8863)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(item.Title)

	user, err := user.Get("luladjiev")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.ID)
}
