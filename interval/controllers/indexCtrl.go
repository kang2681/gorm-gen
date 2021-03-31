package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kang2681/common/log"
)

type IndexCtrl struct {
	baseCtrl
}

func NewIndexCtrl() *IndexCtrl {
	return &IndexCtrl{}
}

// 首页
func (i *IndexCtrl) Index(c *gin.Context) {
	l := log.NewWithUUID()
	c.HTML(200, "", l)
}
