package proto

//“{"chargetype":3,"flag":-1620723712,"userid":参数user ,"reason":1,"point":参数money,"uuid":-1620723712,"permillage":参数money,"action":"payment"}”
type HResp struct {
	Chargetype int    `json:"chargetype"`
	Flag       int    `json:"flag"`
	Userid     string `json:"userid"`
	Reason     int    `json:"reason"`
	Point      string `json:"point"`
	UUID       int    `json:"uuid"`
	Permillage string `json:"permillage"`
	Action     string `json:"action"`
}

//明文json数据“{"passport":"57da00080e","encoding":"b","server":142,"recode":0,"action":"r_getauth","gate":"sw"}”
type TResp struct {
	Passport string `json:"passport"`
	Encoding string `json:"encoding"`
	Server   int    `json:"server"`
	Recode   int    `json:"recode"`
	Action   string `json:"action"`
	Gate     string `json:"gate"`
}
