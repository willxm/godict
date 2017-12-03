package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/willxm/godict/utils"
)

type Youdao struct {
}

type RequestYoudao struct {
	P      string `json:"p,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	AppKey string `json:"app_key,omitempty"`
	Salt   string `json:"salt,omitempty"`
	Sign   string `json:"sign,omitempty"`
}

type ResponseYoudao struct {
	ErrorCode   string   `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	Basic       struct {
		Phonetic   string   `json:"phonetic"`
		UkPhonetic string   `json:"uk-phonetic"`
		UsPhonetic string   `json:"us-phonetic"`
		Explains   []string `json:"explains"`
	} `json:"basic"`
	Web []struct {
		Key   string   `json:"key"`
		Value []string `json:"value"`
	} `json:"web"`
	Dict struct {
		URL string `json:"url"`
	} `json:"dict"`
	Webdict struct {
		URL string `json:"url"`
	} `json:"webdict"`
	L string `json:"l"`
}

func (y *Youdao) Translate(q string) *ResponseYoudao {
	var ry ResponseYoudao
	salt := utils.Rand(1)
	sign := utils.MD5(utils.YOUDAO_APP_ID + q + salt + utils.YOUDAO_APP_KEY)
	queryUrl := "?q=" + q + "&from=EN&to=zh-CHS" + "&appKey=" + utils.YOUDAO_APP_ID + "&salt=" + salt + "&sign=" + sign
	resq, err := http.Get(utils.YOUDAO_TRANSLATE_API + queryUrl)
	if err != nil {
		panic(err)
	}
	defer resq.Body.Close()
	data, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		panic(err)
	}
	jsonErr := json.Unmarshal(data, &ry)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return &ry
}
