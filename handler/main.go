package handler

import (
	"errors"
	"kertas_kerja/contract"
	"kertas_kerja/middleware"
	"kertas_kerja/pkg/errs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller interface {
	// getPrefix returns the route prefix that the controller will use.
	getPrefix() string

	// initService initializes the necessary services for the controller.
	// This service typically contains the business logic required by the controller.
	initService(service *contract.Service)

	// initRoute sets up the routes for the controller within the given router group.
	initRoute(app *gin.RouterGroup)
}

func New(app *gin.Engine, service *contract.Service) {
	allController := []controller{
		&kertasKerjaController{},
		&authController{},
	}

	// do not modify the code below there
	for _, c := range allController {
		c.initService(service)
		group := app.Group(c.getPrefix())
		group.Use(middleware.CORSMiddleware()) // Global Middleware
		c.initRoute(group)
		log.Printf("initiate route %s\n", c.getPrefix())
	}
}

func handlerError(ctx *gin.Context, err error) {
	var messageErr errs.MessageErr
	if errors.As(err, &messageErr) {
		ctx.JSON(messageErr.Status(), messageErr)
	} else {
		ctx.Error(err).SetType(gin.ErrorTypePrivate)
		ctx.JSON(http.StatusInternalServerError, errs.ErrServer)
	}
}
