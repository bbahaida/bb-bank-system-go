package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	} else if err := store.Init(); err != nil {
		log.Fatal(err)
	} else {
		server := NewAPIServer(":3000", store)
		server.Run()
	}

}
