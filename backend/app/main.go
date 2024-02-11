package main

import (
	"backend/module/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Errorln("Error loading .env file ", err)
	}

	r := gin.Default()

	tracer, err := apm.NewTracerOptions(apm.TracerOptions{
		ServiceName:        "Backend Nanda",
		ServiceVersion:     "V.1.2.3",
		ServiceEnvironment: "Misal DEV",
	})

	opts := apmgin.WithTracer(tracer)

	apmgin.WithPanicPropagation()
	r.Use(apmgin.Middleware(r, opts))

	repo := user.NewRepository()
	service := user.NewUsecase(repo)

	user.NewController(service, r)

	r.Run(":9090")
}
