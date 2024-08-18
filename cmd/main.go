package main

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/redis_all/config"
	"github.com/jasurxaydarov/redis_all/db"
	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {

	cfg:=config.Load()
	log:=logger.NewLogger("",logger.LevelDebug)
	redisCli, err := db.ConnRedis(log, context.Background(), cfg.RedisConfig)

	if err != nil {
		fmt.Println(err)

		return

	}

	fmt.Println(redisCli)
}
