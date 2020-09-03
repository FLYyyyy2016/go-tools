package api

import (
	"encoding/json"
	"log"
)

type IPStatusResponse struct {
	Return  int          `json:"ret"`
	Data    IPStatusData `json:"data"`
	Message string       `json:"msg"`
}

type IPStatusData struct {
	ErrCode  int      `json:"err_code"`
	ErrMsg   string   `json:"err_msg"`
	IPStatus IPStatus `json:"data"`
}

type IPStatus struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
	Isp     string `json:"isp"`
}

// when ip is "", result is your ip by default.
func GetMyExtranetIP(ip string) (IPStatus, error) {
	queryUrl := appKey + GetIPStatusServiceName
	sign, err := getSign(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryUrl = myURL + "?" + "&app_key=" + appKey + "&s=" + GetIPStatusServiceName + "&sign=" + sign
	if ip != "" {
		queryUrl = queryUrl + "&ip=" + ip
	}
	log.Println(queryUrl)
	body := getRequest(queryUrl)
	var ipStatusResponse IPStatusResponse
	err = json.Unmarshal(body, &ipStatusResponse)
	if err != nil {
		log.Fatalln(err, ip)
	}
	if ipStatusResponse.Return != 200 {
		return ipStatusResponse.Data.IPStatus, &getMessageError{ipStatusResponse.Return, ipStatusResponse.Message}
	}
	return ipStatusResponse.Data.IPStatus, nil
}
