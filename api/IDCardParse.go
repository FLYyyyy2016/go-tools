package api

import (
	"encoding/json"
	"log"
)

type IDCardResponse struct {
	Return  int        `json:"ret"`
	Data    IDCardData `json:"data"`
	Message string     `json:"msg"`
}
type IDCardData struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Info    IDCardInfo
}

type IDCardInfo struct {
	IsLegal  bool `json:"is_legal"`
	Birthday string
	Age      int
	Gender   string
	Code     int
	Province string
	City     string
	Area     string
	Address  string
}

func GetIDInfoByID(idNumber string) (IDCardInfo, error) {
	queryUrl := appKey + idNumber + IDCardServiceName
	sign, err := getSign(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryUrl = myURL + "?" + "&app_key=" + appKey + "&service=" + IDCardServiceName + "&id_number=" + idNumber + "&sign=" + sign
	log.Println(queryUrl, "request")
	body := getRequest(queryUrl)
	var response IDCardResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}
	if response.Return != 200 {
		return response.Data.Info, &getMessageError{response.Return, response.Message}
	}
	return response.Data.Info, err
}
