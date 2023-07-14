package app

import (
	"fmt"
	"kqdk/utils"
	"strconv"
	"strings"
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

	directionStr := "入"
	direction := 1             // in
	index := settings.In.Index // in
	start := settings.In.Start // in
	end := settings.In.End     // in

	if isIn == false {
		directionStr = "出"
		direction = 0              // out
		index = settings.Out.Index // out
		start = settings.Out.Start // out
		end = settings.Out.End     // out
	}

	dbList := getDataModel(direction, index, start, end)

	appLogger.Info(fmt.Sprintf("方向：%s，时间段：%s-%s，起始索引：%d，数量：%d", directionStr, start, end, index, len(dbList)))

	// 存在数据的时候才提交数据
	if len(dbList) > 0 {

		dbLast := dbList[0]

		iotMessage := getIotMessage(dbList)

		token := appClient.Publish("/v1/devices/SN-HT-XT-BenBuDaLou-RLSB/datas", settings.Mqtt.Qos, false, iotMessage)

		appLogger.Info(fmt.Sprintf("正在推送数据，方向：%s,终止索引：%d，数量：%d", directionStr, dbLast.RecordId, len(dbList)))
		isDone := token.WaitTimeout(30 * time.Second) // Can also use '<-t.Done()' in releases > 1.2.0

		if isDone == true && appClient.IsConnectionOpen() == true {

			appLogger.Info(fmt.Sprintf("方向：%s,推送成功", directionStr))

			if isIn == true {
				settings.In.Index = dbLast.RecordId
			} else {
				settings.Out.Index = dbLast.RecordId
			}
			saveSettings()

		} else {

			appLogger.Error(fmt.Sprintf("方向：%s,推送失败", directionStr))

		}
	}

}

// inOrOut in:1 out:0
// index: 进入或者出去的索引位置，两个索引位置不同
// start,end: 00:00:00 时间段的起止
func getDataModel(inOrOut int, index string, start string, end string) []Record {
	rs := []Record{}

	inOrOutStr := "入口"
	if inOrOut == 0 {
		inOrOutStr = "出口"
	}

	appMysql.Table("aitcp_snapshot_t AS shot").
		Select(`
		shot.device_name,
		shot.snap_time,
		emp.employee_name,
		emp.department_name,
		emp.card_no`).
		Joins("LEFT JOIN aitcp_employee_t AS emp ON shot.matched_person_id = emp.employee_code").
		Where("shot.matched_person_id IS NOT NULL").
		Where("emp.department_name IS NOT NULL").
		Where("LENGTH( card_no ) = 8 ").
		Where("shot.snap_time >= ?", index).
		Where("LEFT( shot.device_name,2 ) = ?", inOrOutStr).
		Where("time( shot.snap_time ) between ? and ?", start, end).
		Order("shot.snap_time desc").
		Scan(&rs)

	for k := range rs {
		rs[k].RecordId = rs[k].AuthTime.Format("2006-01-02 15:04:05")
	}

	return rs
}

// 进入或出入的sql查询数据变成 物联推送数据
func getIotMessage(inOrOutData []Record) []byte {
	iotData := make([]iotService, len(inOrOutData))
	eventTime := time.Now().Format("20060102T150405Z")
	for k := range iotData {
		iotData[k].Data.PassTime = strconv.FormatInt(inOrOutData[k].AuthTime.Add(-8*time.Hour).UnixMilli(), 10)[:13]
		iotData[k].Data.UserName = inOrOutData[k].Name
		iotData[k].Data.Department = inOrOutData[k].DepartMentName
		iotData[k].Data.UserId = inOrOutData[k].IdentityNo
		iotData[k].Data.DeviceSn = settings.Devices.DevicesSn

		iotData[k].EventTime = eventTime
		iotData[k].ServiceId = settings.Devices.ServiceId

		// 2：出方向，1：入方向
		// 出入替换 0->2 1->1
		if strings.Contains(inOrOutData[k].DeviceInout, "出口") == true {
			iotData[k].Data.Type = "2"
		} else {
			iotData[k].Data.Type = "1"
		}
	}

	devices := make([]iotDeviceUpdate, 1)
	devices[0] = iotDeviceUpdate{
		DeviceId: settings.Mqtt.ClientID,
		Services: iotData,
	}
	iot := iotDevices{
		Devices: devices,
	}
	bodyStr := utils.StructToJsonByte(iot)
	return bodyStr
}
