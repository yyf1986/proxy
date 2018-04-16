package models

import (
	"io/ioutil"
	"encoding/base64"
	"net/http"
	"bytes"
	"fmt"
)

func sendpost(url,jsonstr string) (string,string) {
	var jsonStr = []byte(jsonstr)
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
    return resp.Status, string(body)
}

func imagetobs64(image_path string) string {
	img, _ := ioutil.ReadFile(image_path)
	img_base64 := base64.StdEncoding.EncodeToString(img)
	return img_base64
}

func save_base64toimage(img_base64, save_path string) {
	data, _ := base64.StdEncoding.DecodeString(img_base64)
	ioutil.WriteFile(save_path, data, 0666)
}

func CheckServiceName(name string) bool {
	//fmt.Println(ServiceRets)
	//aaa := read2()
	//for _, s := range aaa {
	for _, s := range ServiceRets {
		service_value := s.Value
		if name == service_value {
			return true
		}
	}
	return false
}

func CheckMatchServiceNameAndParms(serviceName string, requestParms map[string]interface{}) bool {
	fmt.Println(requestParms)
	for _, s := range ServiceRets {
		serviceNameFromXml := s.Value
		//serviceTypeFromXml := s.Type
		//urlFromXml := s.Url
		if serviceName == serviceNameFromXml {
			flag := true
			for _, a := range s.Input.Inputs {
				fmt.Println(a.IName,a.IType,a.IFlag)
				fmt.Println(requestParms[a.IName])
				if requestParms[a.IName] == "" {
					flag = false
				}
			}
			fmt.Println(flag)
			return flag
			/*	
			for _, b := range s.Output.Outputs {
				fmt.Println(b)
			}*/
		}
	}
	return false
}