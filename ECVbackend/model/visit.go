package model

import "time"

type Visit struct {
	StartTime   time.Time
	EndTime     time.Time
	LatencyTime time.Duration
	ReqMethod   string
	ReqUri      string
	StatusCode  int
	ClientIp    string
}

func (v Visit) Insert() {
	dbEngine.Insert(&v)
}

func (v Visit) Count() int64 {
	num, _ := dbEngine.Where("req_uri REGEXP '/api/'").Count(&v)
	return num
}
