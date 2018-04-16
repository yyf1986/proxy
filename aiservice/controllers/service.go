package controllers

import (
	"aiservice/models"
	"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
)

// Operations about Service
type ServiceController struct {
	beego.Controller
}

type RBody struct {
	ServiceName string `json:"serviceName"`
	RequestParms map[string]interface{} `json:"requestParms"`
}

func (s *ServiceController) Post() {
	respId, _ := uuid.NewV4()
	var data RBody
	json.Unmarshal([]byte(s.Ctx.Input.RequestBody), &data)
	serviceName := data.ServiceName
	requestParms := data.RequestParms
	switch {
	//serviceName参数是否为空
	case serviceName == "":
		logs.Error("no serviceName")
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"Argument Required: serviceName", 
												"respInfo": ""}
		s.ServeJSON()
	//requestParms参数是否为空
	case len(requestParms) == 0:
		logs.Error("no requestParms")
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"Argument Required: requestParms", 
												"respInfo": ""}
		s.ServeJSON()
	//传入的serviceName是否在xml配置中
	case models.CheckServiceName(serviceName) == false:
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"serviceName not exist", 
												"respInfo": ""}
		s.ServeJSON()
	//检查对应的serviceName，对应的requestParms中传入的参数是否匹配
	case models.CheckMatchServiceNameAndParms(serviceName, requestParms) == false:
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"aaaa not exist", 
												"respInfo": ""}
		s.ServeJSON()
	}
	s.Data["json"] = map[string]interface{}{"respId": respId, 
											"respCode": 200, 
											"respMsg":"", 
											"respInfo": "ok"}
	s.ServeJSON()
}