package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/betterfor/BookDown/models"
	"github.com/betterfor/gologger"
	"github.com/betterfor/gorequest"
)

// RandomUA will return a random user agent.
func RandomUA() string {
	return gorequest.RandomUA()
}

// 从 GoPool （https://github.com/betterfor/GoPool）中获取IpPool
func GetProxies() (ips []models.Ippool, err error) {
	url := "http://localhost:9000/list"
	_, body, errs := gorequest.New().Get(url).End()
	if len(errs) != 0 {
		gologger.Errorf("[GetProxies] request %s error:%v", url, errs)
		return nil, errors.New(fmt.Sprintf("%v", errs))
	}
	err = json.Unmarshal([]byte(body), &ips)
	return
}
