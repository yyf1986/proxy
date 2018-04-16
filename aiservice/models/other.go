package models

import (
	"strings"
	"time"
)

func getImagePath(dd time) string {
	aaa := dd.Format("2006_01_02_15_04_05.99999999")
	bbb := strings.Replace(aaa,`.`,`_`,1)
	return bbb
}