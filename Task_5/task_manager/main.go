package main

import (
	"task_manager/data"
	"task_manager/router"
)
func main() {
	data.InitUserCollection()
	r := router.SetupRouter()
	r.Run()
}