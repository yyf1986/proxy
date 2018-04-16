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
	case serviceName == "":
		logs.Error("no serviceName")
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"json string has no serviceName", 
												"respInfo": ""}
		s.ServeJSON()
	case len(requestParms) == 0:
		logs.Error("no requestParms")
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"json string has no requestParms", 
												"respInfo": ""}
		s.ServeJSON()
	case models.CheckServiceName(serviceName) == false:
		s.Data["json"] = map[string]interface{}{"respId": respId, 
												"respCode": 200, 
												"respMsg":"serviceName not exist", 
												"respInfo": ""}
		s.ServeJSON()
	}
	s.Data["json"] = map[string]interface{}{"respId": respId, 
											"respCode": 200, 
											"respMsg":"", 
											"respInfo": "ok"}
	s.ServeJSON()
}