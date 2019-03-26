package main

import "testing"

const testJson1 = `{"Username":"<username>","Password":"<password>","Url":"https://st-ccsd.accruenttest.net","Endpoint":"/MobileWebServices/apis/360facility/v1/assets","Record":{"Name":"NY3-0110BMS -01-BMS","SerialNumber":"46171613","StatusComment":"Detection Method: Semiconductor\r\nPower Source: 120 VAC\r\nAudible Alert: 85dB @ 10ft\r\nOperating Temp: -40 to 150 F\r\n","BarcodeNumber":"1","SpaceId":598,"ActiveFlag":false,"MakeId":413,"ModelId":1033,"AssetStatusId":1,"AssetRankId":3,"AssetClassId":107,"AssetNumber":"134"}}`
const testJson2 = `{"Username":"<username>","Password":"<password>","Url":"https://st-ccsd.accruenttest.net","Endpoint":"/MobileWebServices/apis/360facility/v1/assets","Record":{"SerialNumber":"46171613","StatusComment":"Detection Method: Semiconductor\r\nPower Source: 120 VAC\r\nAudible Alert: 85dB @ 10ft\r\nOperating Temp: -40 to 150 F\r\n"},"Id":8799}`

func TestCreateRecordJson(test *testing.T) {
	create := CreateRecord{}
	data, err := create.ExecuteJson(testJson1)
	if err != nil {
		panic(err)
	}
	if jsonData, ok := data.(JsonMap); ok {
		println(jsonData.String())
	} else {
		println(data.(string))
	}
}

func TestUpdateRecordJson(test *testing.T) {
	update := UpdateRecord{}
	data, err := update.ExecuteJson(testJson2)
	if err != nil {
		panic(err)
	}
	if jsonData, ok := data.(JsonMap); ok {
		println(jsonData.String())
	} else {
		println(data.(string))
	}
}
