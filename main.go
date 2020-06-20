package main

import (
	_ "github.com/lib/pq"
	"github.com/sureeratend/finalexam/todo"
)

func main() {
	r := todo.SetupRouter()
	r.Run(":2019")
}
