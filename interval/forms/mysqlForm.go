package forms

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kang2681/common/log"
)

type MysqlForm struct {
	l *log.Logger
}

func NewMysqlForm(l *log.Logger) *MysqlForm {
	return &MysqlForm{
		l: l,
	}
}

type MysqlConnectParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
}

func (m *MysqlForm) CheckConnect(c *gin.Context) (*MysqlConnectParam, error) {
	data, err := c.GetRawData()
	if err != nil {
		m.l.Errorf("get raw data error:%s", err.Error())
		return nil, err
	}
	rs := MysqlConnectParam{}
	if err := json.Unmarshal(data, &rs); err != nil {
		m.l.Errorf("json unmarshal error %s, data:%s", err.Error(), data)
		return nil, err
	}
	return &rs, nil
}
