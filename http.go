package crud

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type crud[T any] struct {
	db        *gorm.DB
	handler   *Handler[T]
	routeName string
}

func NewCRUD[T any](db *gorm.DB, routeName string) *crud[T] {
	handler := NewHandler[T](db)
	return &crud[T]{db: db, routeName: routeName, handler: handler}
}

func (c *crud[T]) RegisterRoutes(router *gin.RouterGroup) {

	router.GET(fmt.Sprintf("/%s", c.routeName), c.handler.List)
	router.POST(fmt.Sprintf("/%s", c.routeName), c.handler.Create)
	router.GET(fmt.Sprintf("/%s/:id", c.routeName), c.handler.Get)
	router.PUT(fmt.Sprintf("/%s/:id", c.routeName), c.handler.Update)
	router.DELETE(fmt.Sprintf("/%s/:id", c.routeName), c.handler.Delete)
}
