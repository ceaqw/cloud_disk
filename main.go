package main

import (
	"CouldDisk/conf"
	"CouldDisk/router"
	"fmt"
	"log"
)

func main() {
	appCfg := conf.GetAppCfg()
	r := router.Init()
	err := r.Run(fmt.Sprintf("%s:%s", appCfg.Host, appCfg.Port))

	if err != nil {
		log.Fatalf("Start server: %+v", err)
	} else {
		log.Fatalf("Start server: success")
	}
}
