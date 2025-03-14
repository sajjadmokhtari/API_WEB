package api

import (
	"GOLANG_CLEAN_WEB_API/src/api/routers"
	"GOLANG_CLEAN_WEB_API/src/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initserver() {
    cfg := config.GetConfig()
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		heath := v1.Group("/health")
		routers.Health(heath)

	}

	r.Run(fmt.Sprintf(":%s",cfg.Server.Port))

}
