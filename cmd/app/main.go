package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aristorinjuang/weight-go/internal/application/config"
	"github.com/aristorinjuang/weight-go/internal/infrastructure/database/mysql"
	grpcServer "github.com/aristorinjuang/weight-go/internal/presentation/grpc"
	httpServer "github.com/aristorinjuang/weight-go/internal/presentation/http"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	))
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	weightRepository := mysql.NewWeightRepositoryMySQL(db)

	go func() {
		log.Panic(grpcServer.New(c.Port.Grpc, weightRepository).Run())
	}()
	log.Panic(httpServer.New(c.Port.Grpc, c.Port.Http, weightRepository).Run())
}
