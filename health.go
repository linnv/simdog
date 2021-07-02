package simdog

import (
	"encoding/json"
	"net/http"
	"time"
)

func HeadlCheckHandler(w http.ResponseWriter, r *http.Request) {
	type respStruct struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	now := time.Now()
	oneResp := respStruct{
		Code: http.StatusOK,
		Msg:  now.String() + ": running",
	}
	bs, _ := json.Marshal(oneResp)
	w.Write(bs)
}
