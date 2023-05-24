package app

import (
	"fmt"
	"kqdk/utils"
	"strconv"
	"time"
)

func cronIn() {
	getInData(true)
}

func cronOut() {
	getInData(false)
}

// inOrOut: ture:in out:false
func getInData(isIn bool) {

	defer func() {
		if e := recover(); e != nil {
			appLogger.Panic(fmt.Sprintf("%s", e))
		}
	}()

	direction := 1             // in
	index := settings.In.Index // in
	start := settings.In.Start // in
	end := settings.In.End     // in

	if isIn == false {
		direction = 0              // out
		index = settings.Out.Index // out
		start = settings.Out.Start // out
		end = settings.Out.End     // out
	}

	dbList := getDataModel(direction, index, start, end)

	if isIn {
		appLogger.Info(fmt.Sprintf("方向：入，时间段：%s-%s，起始索引：%d，数量：%d", start, end, index, len(dbList)))
	} else {
		appLogger.Info(fmt.Sprintf("方向：出，时间段：%s-%s，起始索引：%d，数量：%d", start, end, index, len(dbList)))
	}

	// 存在数据的时候才提交数据
	if len(dbList) > 0 {

		dbLast := dbList[0]

		iotMessage := getIotMessage(dbList)

		token := appClient.Publish("/v1/devices/SNs-HT-XT-BenBuDaLou-RLSB/datas", 2, false, iotMessage)

		// token是阻塞，需要启动多线程
		appLogger.Info(fmt.Sprintf("正在推送数据，终止索引：%d，数量：%d,", dbLast.RecordId, len(dbList)))
		isDone := token.WaitTimeout(30 * time.Second) // Can also use '<-t.Done()' in releases > 1.2.0
		if token.Error() != nil {
			appLogger.Error("推送失败")
			appLogger.Error(token.Error().Error()) // Use your preferred logging technique (or just fmt.Printf)
		} else {
			appLogger.Info("推送成功")
		}

		if isDone == true {
			if isIn == true {
				settings.In.Index = dbLast.RecordId
			} else {
				settings.Out.Index = dbLast.RecordId
			}
			saveSettings()
		}
	}

}

// inOrOut in:1 out:0
// index: 进入或者出去的索引位置，两个索引位置不同
// start,end: 00:00:00 时间段的起止
func getDataModel(inOrOut int, index int, start string, end string) []Record {
	rs := []Record{}
	//sqlite
	//appSqlite.
	//	Where("LENGTH(IdentityNo)=8").
	//	Where("Record_ID > ?", index).
	//	Where("Device_InOut = ?", inOrOut).
	//	Where("strftime('%H:%M:%S',AuthTime) BETWEEN ? and ?", start, end).
	//	Order("Record_ID DESC").
	//	Find(&rs)
	//sql server
	appMssql.
		Where("DATALENGTH(IdentityNo)=8").
		Where("Record_ID > ?", index).
		Where("Device_InOut = ?", inOrOut).
		Where("CONVERT(char(8), AuthTime, 108) BETWEEN ? and ?", start, end).
		Order("Record_ID DESC").Find(&rs)
	return rs
}

// 进入或出入的sql查询数据变成 物联推送数据
func getIotMessage(inOrOutData []Record) []byte {
	iotData := make([]iotService, len(inOrOutData))
	eventTime := time.Now().Format("20060102T150405Z")
	for k := range iotData {
		iotData[k].Data.PassTime = strconv.FormatInt(inOrOutData[k].AuthTime.UnixMilli(), 10)[:13]
		iotData[k].Data.UserName = inOrOutData[k].Name
		iotData[k].Data.Department = inOrOutData[k].DepartMentName
		iotData[k].Data.UserId = inOrOutData[k].IdentityNo
		iotData[k].Data.DeviceSn = settings.Devices.DevicesSn

		iotData[k].EventTime = eventTime
		iotData[k].ServiceId = settings.Devices.ServiceId

		// 出入替换 0->2 1->1
		if inOrOutData[k].DeviceInout == 0 {
			iotData[k].Data.Type = "2"
		} else {
			iotData[k].Data.Type = "1"
		}
	}

	devices := make([]iotDeviceUpdate, 1)
	devices[0] = iotDeviceUpdate{
		DevicesId: settings.Mqtt.ClientID,
		Services:  iotData,
	}
	iot := iotDevices{
		Devices: devices,
	}
	bodyStr := utils.StructToJsonByte(iot)
	return bodyStr
}
