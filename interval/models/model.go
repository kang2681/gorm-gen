package models

import (
	"fmt"
	"sync"

	"github.com/kang2681/common/drivers"
	"github.com/kang2681/common/stringsext"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConnectMap struct {
	sync.Mutex
	cache map[string]*gorm.DB
}

type cacheData struct {
	DB   *gorm.DB
	Conf *drivers.MysqlConfig
}

func NewConnectMap() *ConnectMap {
	return &ConnectMap{
		cache: make(map[string]*gorm.DB),
	}
}

func (c *ConnectMap) Connect(conf *drivers.MysqlConfig) error {
	c.Lock()
	defer c.Unlock()
	key := c.getUniqueKey(conf)
	if _, ok := c.cache[key]; ok {
		return fmt.Errorf("已连接")
	}
	return nil
}

func (c *ConnectMap) getUniqueKey(conf *drivers.MysqlConfig) string {
	str := fmt.Sprintf("%s-%s-%s-%s-%d", conf.Username, conf.Password, conf.DBName, conf.Host, conf.Port)
	return stringsext.MD5(str)
}
