package models

import (
	"fmt"
	"time"
)

type Ippool struct {
	Id         int64     `json:"id"`
	Protocol   string    `json:"protocol"`
	Ip         string    `json:"ip"`
	Port       string    `json:"port"`
	CreateTime time.Time `json:"create_time"`
	Deleted    int64     `json:"deleted"`
}

func (ip *Ippool) String() string {
	return fmt.Sprintf("%s://%s:%s", ip.Protocol, ip.Ip, ip.Port)
}
