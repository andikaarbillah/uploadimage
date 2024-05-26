package api

import (
	"images/config"
	"images/router"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	db := config.DB()

	router.Router(r, db)

	r.Run()
}
