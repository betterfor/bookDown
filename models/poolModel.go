package models

import (
	"fmt"
	"github.com/betterfor/BookDown/db"
	"github.com/betterfor/GoLogger/logger"
	"time"
)

type Ippool struct {
	Id         int64     `json:"id" 			xorm:"id" 					form:"id"`
	Protocol   string    `json:"protocol" 		xorm:"protocol" 			form:"protocol"`
	Ip         string    `json:"ip" 			xorm:"ip" 					form:"ip"`
	Port       string    `json:"port" 			xorm:"port" 				form:"port"`
	CreateTime time.Time `json:"create_time" 	xorm:"create_time created" 	form:"create_time"`
	Deleted    int64     `json:"deleted"		xorm:"deleted"				form:"deleted"`
}

func (p *Ippool) TableName() string {
	return "ippool"
}

// 插入一条记录
func (p *Ippool) InsertOne(rel Ippool) error {
	e := db.GetEngine()
	_, err := e.InsertOne(rel)
	if err != nil {
		logger.Errorf("[InsertOne] Insert %+v into db error: %s", rel, err.Error())
		return err
	}
	return err
}

// 插入多条记录
func (p *Ippool) InsertMany(rels []*Ippool) error {
	e := db.GetEngine()
	_, err := e.Insert(rels)
	return err
}

// 获取一条记录
func (p *Ippool) GetById(id int64) (*Ippool, error) {
	e := db.GetEngine()
	Rel := new(Ippool)
	err := db.GetCheck(e.ID(id).Get(Rel))
	return Rel, err
}

// 获取分页列表
func (p *Ippool) GetList(offset, limit int64) (rels []*Ippool, total int64, err error) {
	e := db.GetEngine()
	sql := "select * from ippool"

	pool := new(Ippool)
	total, err = e.SQL(sql).Count(pool)
	if err != nil {
		return nil, 0, err
	}
	sql = fmt.Sprintf("%s ?,?", sql)
	err = e.SQL(sql, offset, limit).Find(&rels)
	return
}

// 删除一条记录
func (p *Ippool) Delete(id int64) error {
	e := db.GetEngine()
	session := e.NewSession()
	defer session.Close()
	pool := new(Ippool)
	pool.Id = id
	_, err := session.Where("id=?", id).Get(pool)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Where("delete=?", pool.Deleted+1).Update(pool)
	if err != nil {
		session.Rollback()
		return err
	}
	return nil
}

// 随机获取一条记录
func (p *Ippool) GetOneRandom() (*Ippool, error) {
	//sql := "select * from ippool order by rand() limit 1"
	var rels []*Ippool
	e := db.GetEngine()
	err := e.OrderBy("rand()").Limit(1).Find(&rels)
	return rels[0], err
}

// 获取总数
func (p *Ippool) GetNums() (int64, error) {
	e := db.GetEngine()
	pool := new(Ippool)
	total, err := e.Where("id > ?", -1).Count(pool)
	return total, err
}

func (p *Ippool) String() string {
	return p.Protocol + "://" + p.Ip + ":" + p.Port
}
