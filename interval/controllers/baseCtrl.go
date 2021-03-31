package controllers

import "github.com/gin-gonic/gin"

type baseCtrl struct{}

func (b *baseCtrl) Html(c *gin.Context) {
	c.HTML()
}
