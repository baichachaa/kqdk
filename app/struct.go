package app

import (
	"time"
)

type iotServiceData struct {
	PassTime   string `json:"passTime"`   // 通过时间
	UserName   string `json:"userName"`   // 姓名
	Department string `json:"department"` // 部门
	Type       string `json:"type"`       // 2:出，1入
	UserId     string `json:"userId"`     // 人资编码
	DeviceSn   string `json:"deviceSn"`
}

type iotService struct {
	Data      iotServiceData `json:"data"`
	EventTime string         `json:"eventTime"`
	ServiceId string         `json:"serviceId"`
}

type iotDeviceUpdate struct {
	DevicesId string       `json:"devicesId"` // clientId
	Services  []iotService `json:"services"`
}

type iotDevices struct {
	Devices []iotDeviceUpdate `json:"devices"`
}

type MyJsonTime time.Time

func (this MyJsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 21)
	b = append(b, '"')
	b = time.Time(this).AppendFormat(b, "2006-01-02 15:04:05")
	b = append(b, '"')
	return b, nil
}

func (this *MyJsonTime) UnmarshalJSON(data []byte) (err error) {
	t, _ := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*this = MyJsonTime(t)
	return
}
