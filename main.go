package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strconv"
	"time"
)

func main() {
	t2, err := time.Parse(time.RFC3339, "2021-01-29T13:56:15-08:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t2)
	// t2.MarshalJSON()

	// Create a Resty Client
	client := resty.New()

	// token := os.Getenv("CRM_API_TOKEN")

	// ref: https://developer.freshsales.io/api/#view_a_lead
	// https://appscode.freshsales.io/leads/5022967942
  	//  /api/leads/[id]
  	/*
  	curl -H "Authorization: Token token=sfg999666t673t7t82" -H "Content-Type: application/json" -X GET "https://domain.freshsales.io/api/leads/1"
  	*/

  	type LeadResponse struct {
		Lead Lead `json:"lead"`
	}

	resp, err := client.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit": "20",
			"sort":"name",
			"order": "asc",
			"random":strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Token token=%s", os.Getenv("CRM_API_TOKEN"))).
		// SetAuthToken(os.Getenv("CRM_API_TOKEN")).
		SetResult(LeadResponse{}).
		Get("https://appscode.freshsales.io/api/leads/5022967942")

	rs2 := resp.Result().(*LeadResponse)
	ldata, err := json.MarshalIndent(rs2.Lead, "", "  ")
	fmt.Println(string(ldata))


	//resp, err := client.R().
	//	EnableTrace().
	//	Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

}
