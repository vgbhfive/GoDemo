package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Fund struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Code                  string  `json:"code"`
		Name                  string  `json:"name"`
		NetWorthDate          string  `json:"netWorthDate"`
		NetWorth              float64 `json:"netWorth"`
		DayGrowth             string  `json:"dayGrowth"`
		ExpectWorthDate       string  `json:"expectWorthDate"`
		ExpectWorth           float64 `json:"expectWorth"`
		ExpectGrowth          string  `json:"expectGrowth"`
		LastWeekGrowth        string  `json:"lastWeekGrowth"`
		LastMonthGrowth       string  `json:"lastMonthGrowth"`
		LastThreeMonthsGrowth string  `json:"lastThreeMonthsGrowth"`
		LastSixMonthsGrowth   string  `json:"lastSixMonthsGrowth"`
		LastYearGrowth        string  `json:"lastYearGrowth"`
	} `json:"data"`
	Meta string `json:"meta"`
}

type FundInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Title     string     `json:"title"`
		Date      string     `json:"date"`
		Stock     string     `json:"stock"`
		Bond      string     `json:"bond"`
		Cash      string     `json:"cash"`
		Total     string     `json:"total"`
		StockList [][]string `json:"stockList"`
	} `json:"data"`
	Meta interface{} `json:"meta"`
}

func main() {
	var a = flag.Bool("a", false, "Export all data of the fund")
	var f = flag.String("f", "", "Enter the numbers of the funds to be queried (separated by , )")
	var i = flag.Bool("i", false, "Export fund details")
	flag.Parse()

	fmt.Println("Hello Fund!")

	var code string = *f
	if len(code) == 0 || (len(code)+1)%7 != 0 {
		fmt.Println("Please check your fundcode!")
		return
	}
	resp, err := http.Get("https://api.doctorxiong.club/v1/fund?code=" + code)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res Fund
	_ = json.Unmarshal(body, &res)

	codes := strings.Split(code, ",")
	for idx, FundCode := range codes {
		if res.Code != 200 {
			fmt.Println(string(body))
			return
		}
		if *a {
			fmt.Printf("%d - %+v \n", idx, res.Data[idx])
		} else {
			fmt.Printf("%d - FundCode:%#v, FundName:%#v, ExpectGrowth:%#v\n",
				idx, res.Data[idx].Code, res.Data[idx].Name, res.Data[idx].ExpectGrowth)
		}

		if *i {
			fundInfo(FundCode, idx)
		}

		fmt.Printf("------------------------------------ \n\n")
	}

}

func fundInfo(code string, index int) {
	resp, err := http.Get("https://api.doctorxiong.club/v1/fund/position?code=" + code)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	var res FundInfo
	_ = json.Unmarshal(body, &res)

	fmt.Printf("%s - Date:%#v, Stock:%#v, Bound:%#v, Cash:%#v, \n", code, res.Data.Date, res.Data.Stock, res.Data.Bond, res.Data.Cash)

	if len(res.Data.StockList) != 0 {
		for i := 0; i < len(res.Data.StockList); i++ {
			fmt.Printf("%d-%d - %+v \n", index, i, res.Data.StockList[i])
		}
	}
}
