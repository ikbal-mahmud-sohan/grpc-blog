package main

import (
	"myGo/Blogs/blog/storage/postgres"
	"log"
)


func main(){
	if err := postgres.Migrate();err !=nil{
		log.Fatal(err)
	}
}