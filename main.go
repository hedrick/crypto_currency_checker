package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Coin struct {
	Cap24hrChange float64 `json:"cap24hrChange"`
	Long          string  `json:"long"`
	Mktcap        float64 `json:"mktcap"`
	Perc          float64 `json:"perc"`
	Price         float64 `json:"price"`
	Shapeshift    bool    `json:"shapeshift"`
	Short         string  `json:"short"`
	Supply        float64 `json:"supply"`
	UsdVolume     float64 `json:"usdVolume"`
	Volume        float64 `json:"volume"`
	VwapData      float64 `json:"vwapData"`
	VwapDataBTC   float64 `json:"vwapDataBTC"`
}

func main() {
	url := "http://coincap.io/front"

	cl := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println("Error:", err)
	}

	req.Header.Set("User-Agent", "golang http tutorial")

	res, getErr := cl.Do(req)
	if getErr != nil {
		fmt.Println("Get Error:", getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Read Error:", readErr)
	}

	coins := []Coin{}
	jsonErr := json.Unmarshal(body, &coins)
	if jsonErr != nil {
		fmt.Println("JSON Error:", jsonErr)
	}

	for c := range coins {
		fmt.Println(c.Long)
	}

	fmt.Println(coins)
}
