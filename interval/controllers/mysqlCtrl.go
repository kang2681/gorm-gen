package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kang2681/common/log"
	"github.com/kang2681/modtool/interval/forms"
)

type MysqlCtrl struct {
	baseCtrl
}

func NewMysqlCtrl() *MysqlCtrl {
	return &MysqlCtrl{}
}

func (m *MysqlCtrl) Connect(c *gin.Context) {
	l := log.NewWithUUID()
	f := forms.NewMysqlForm(c)
	param, err := f.CheckConnect(c)
	if err != nil {
		l.Errorf("form check connect error %s", err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, param)
}
