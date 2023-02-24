package main

import (
	"github.com/andrelensanro/building_twitter_copy/handlers"
	"github.com/andrelensanro/building_twitter_copy/bd"
	"log"
)
func main(){
	if bd.CheckConnection() == 0{
		log.Fatal("Without a connection to the DB")
		return 
	}
	handlers.Handlers()
}
