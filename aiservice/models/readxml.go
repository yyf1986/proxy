package models

import (
	"encoding/xml"
	"io/ioutil"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type InputPara struct {
	IName string `xml:"name,attr"`
	IType string `xml:"type,attr"`
	IFlag string `xml:"flag,attr"`
}

type OutputPara struct {
	OName string `xml:"name,attr"`
	OType string `xml:"type,attr"`
}

type Input struct {
	Inputs []InputPara `xml:"para"`
}

type Output struct {
	Outputs []OutputPara `xml:"para"`
}

type Service struct {
	//XMLName  xml.Name   `xml:"serviceName"`
	Value string `xml:"value,attr"`
	Type string `xml:"type,attr"`
	Url string `xml:"requestUrl"`
	Input Input `xml:"input"`
	Output Output `xml:"output"`
}

type Services struct {
	XMLName  xml.Name   `xml:"serviceNames"`
	Service []Service `xml:"serviceName"`
}

var ServiceRets []Service

func read(){
	xmlpath := beego.AppConfig.String("xmlPath")
	content, err := ioutil.ReadFile(xmlpath)
    if err != nil {
        logs.Error(err)
    }
    var result Services
    err = xml.Unmarshal(content, &result)
    if err != nil {
        logs.Error(err)
    }
    //fmt.Println(result.Service)
	ServiceRets = result.Service
}

func read2() []Service{
	xmlpath := beego.AppConfig.String("xmlPath")
	content, err := ioutil.ReadFile(xmlpath)
    if err != nil {
        logs.Error(err)
    }
    var result Services
    err = xml.Unmarshal(content, &result)
    if err != nil {
        logs.Error(err)
    }
    //logs.Info(result.Service)
	return result.Service
}

func init() {
	read()
}