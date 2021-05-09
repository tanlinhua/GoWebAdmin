package main

import (
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/router"
	"github.com/tanlinhua/go-web-admin/server/cron"
)

func main() {
	model.InitDB()
	go cron.Work()
	go router.InitAdmServer()
	router.InitApiServer()
}
