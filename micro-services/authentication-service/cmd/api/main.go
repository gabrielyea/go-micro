package main

import (
	"auth/db"
	"fmt"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Errorf("something went wrong with db %s", err)
	}
	defer db.Close()

}
