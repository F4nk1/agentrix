package main 

import (
	"context"
	"fmt"
	"net/http"

	"github.com/F4nk1/agentrix/backend/source/config"
	"github.com/F4nk1/agentrix/backend/source/database"
	"github.com/F4nk1/agentrix/backend/source/server"
	"github.com/F4nk1/agentrix/backend/source/service"
	"github.com/F4nk1/agentrix/backend/source/tracer"
	
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfiguration()
	tracer.Infof(ctx, "Starting database host '%s:%s'", cfg.Database.Host, cfg.Database.Port)

	conn, err := database.NewDataBase(ctx, cfg)
	if err != nil {
		tracer.Fatal(ctx, err)
	}
	defer conn.Close()
	
	service := service.NewService(conn)
	server := server.NewServer(service)

	serverHost := fmt.Sprintf(":%s",cfg.Server.Port)
	tracer.Infof(ctx, "Starting server on %s", serverHost)
	tracer.Fatal(ctx, http.ListenAndServe(serverHost,server.Handler))

}