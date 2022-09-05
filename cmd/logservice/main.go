package main

import (
	"context"
	"fmt"
	"github.com/karthick-workspace/distributed_app/log"
	"github.com/karthick-workspace/distributed_app/service"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"LogService",
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()

	fmt.Println("Shutting down log service")

}
