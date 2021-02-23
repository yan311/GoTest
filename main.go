package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type StrResources struct {
	XMLName xml.Name `xml:"food_menu"` //标签名称
	Food []Foods `xml:"food"` //获取food数组
}

type Foods struct{
	Name string `xml:"name"` //获取name配置项，并将结果保存到Name中
	Price string `xml:"price"`
	Description string `xml:"description"`
	Calories int `xml:"calories"`
}

func main(){
	//打开xml文件
	file, err := os.Open("FoodMenu.xml")
	if err != nil{
		fmt.Println("打开文件错误信息：",err)
		return
	}
	defer file.Close()

	//读取文件
	data, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println("读取文件错误信息：",err)
		return
	}
	//将文件转换为对象
	var result StrResources
	err = xml.Unmarshal(data,&result)
	if err != nil{
		fmt.Println("读文件内容错误信息：",err)
		return
	}

	fmt.Println("result：",result)
	fmt.Println()
	for _ ,o := range result.Food{
	    fmt.Println("name：",o.Name)
		fmt.Println("price：",o.Price)
		fmt.Println("description：",o.Description)
		fmt.Println("calories：",o.Calories)
		fmt.Println()
	}
	fmt.Println("读取完成！")
}