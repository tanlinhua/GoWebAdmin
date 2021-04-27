package main

import (
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/router"
)

func main() {
	model.InitDB()
	go router.InitAdmServer()
	router.InitApiServer()
}
