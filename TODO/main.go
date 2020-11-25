package main

import (
	"time"

	"github.com/arijitnayak92/taskAfford/TODO/app"
	"github.com/arijitnayak92/taskAfford/TODO/config"
	"github.com/arijitnayak92/taskAfford/TODO/domain"
)

func main() {
	config.Load()
	for i := 0; i < 4; i++ {
		time.Sleep(3 * time.Second)
		postgres := domain.InitDB()
		if postgres == nil {
			panic("postgres is nil")
		}
	}
	app.StartApp()
}
