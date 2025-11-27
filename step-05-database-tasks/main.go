package main

import "log"

func main() {
	db := ConnectDB()
	defer db.Close()

	log.Println("Database ready!")
}
