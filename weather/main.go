package main

import (
	"fmt"
	"go-soa/weather/weatherService"

	"github.com/hooklift/gowsdl/soap"
	"github.com/xuri/excelize/v2"
)

func main() {
	client := soap.NewClient("http://ws.webxml.com.cn/WebServices/WeatherWebService.asmx")

	service := weatherService.NewWeatherWebServiceSoap(client)

	result, err := service.GetWeatherbyCityName(&weatherService.GetWeatherbyCityName{
		TheCityName: "长沙",
	})

	if err != nil {
		panic(err)
	}

	data := result.GetWeatherbyCityNameResult.Astring
	for _, v := range data {
		fmt.Println(*v)
	}

	// ------------- 保存数据到 Excel -------------
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置单元格的值
	for i, v := range data {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), *v)
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("weather.xlsx"); err != nil {
		fmt.Println(err)
	}
}
