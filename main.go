package main

import (
	"fmt"
	"main.go/config"
	"main.go/modules/resident/storage"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(storage.QueryCountry(cfg.VietNamXml))
}
