package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type MailData struct {
	Name    string
	Product string
}

func main() {
	opts := Options{
		MailgunDomain:        os.Getenv("MG_DOMAIN"),
		MailgunPrivateAPIKey: os.Getenv("MG_API_KEY"),
		MailSender:           "tamal@appscode.com",
		MailLicenseTracker:   "michael+bcc@appscodehq.com",
		MailReplyTo:          "michael@appscodehq.com",
	}

	recipient := "michael@appscodehq.com"
	subject := "mailgun tracker"

	src := `Hi {{.Name}},
Thanks for your interest in {{.Product}}. The license for [Kubernetes](https://kubernetes.io) cluster **Cluster** is attached with this email.

Please let us know if you have any questions.

Regards,
AppsCode Team`
	bodyText, bodyHtml, err := RenderMail(src, MailData{
		Name:    "John",
		Product: "kubedb-enterprise",
	})
	if err != nil {
		panic(err)
	}

	// fmt.Println(bodyHtml)
	// bodyHtml= html_email

	// recipient, subject, bodyText, bodyHtml string, attachments map[string][]byte) error {

	err = SendMail(opts, recipient, subject, bodyText, bodyHtml, nil)
	if err != nil {
		panic(err)
	}
	os.Exit(1)

	t2, err := time.Parse(time.RFC3339, "2021-01-29T13:56:15-08:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t2)
	// t2.MarshalJSON()

	// Create a Resty Client
	client := resty.New().
		EnableTrace().
		SetHostURL("https://appscode.freshsales.io").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Token token=%s", os.Getenv("CRM_API_TOKEN")))

	// lookup search
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":        "tamal.saha@gmail.com",
			"f":        "email",
			"entities": "lead,contact",
		}).
		SetResult(&LookupResult{}).
		Get("/api/lookup")
	if err != nil {
		panic(err)
	}

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	os.Exit(1)

	// add note

	resp, err = client.R().
		SetBody(APIObject{Note: &Note{
			Description:    "Issued license for cluster xyz",
			TargetableType: "Lead",
			TargetableID:   5023512614,
		}}).
		SetResult(&APIObject{}). // or SetResult(AuthSuccess{}).
		// SetError(&AuthError{}).       // or SetError(AuthError{}).
		Post("/api/notes")
	if err != nil {
		panic(err)
	}
	rs5 := resp.Result().(*APIObject)
	ndata, err := json.MarshalIndent(rs5, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ndata))

	os.Exit(1)

	// update lead

	// 5023512614

	resp, err = client.R().
		SetBody(APIObject{Lead: &Lead{
			JobTitle: "Servant",
		}}).
		SetResult(&APIObject{}). // or SetResult(AuthSuccess{}).
		// SetError(&AuthError{}).       // or SetError(AuthError{}).
		Put("/api/leads/5023512614")
	if err != nil {
		panic(err)
	}

	rs4 := resp.Result().(*APIObject)
	udata, err := json.MarshalIndent(rs4, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(udata))

	// search lead

	resp, err = client.R().
		SetQueryParams(map[string]string{
			"q":       "tamal.saha@gmail.com",
			"include": "lead",
		}).
		SetResult(SearchResults{}).
		Get("/api/search")

	rs3 := resp.Result().(*SearchResults)
	rdata, err := json.MarshalIndent(rs3, "", "  ")
	fmt.Println(string(rdata))

	// Get Lead by id

	// ref: https://developer.freshsales.io/api/#view_a_lead
	// https://appscode.freshsales.io/leads/5022967942
	//  /api/leads/[id]
	/*
		curl -H "Authorization: Token token=sfg999666t673t7t82" -H "Content-Type: application/json" -X GET "https://domain.freshsales.io/api/leads/1"
	*/

	resp, err = client.R().
		SetResult(APIObject{}).
		Get("/api/leads/5006838695")

	rs2 := resp.Result().(*APIObject)
	ldata, err := json.MarshalIndent(rs2.Lead, "", "  ")
	fmt.Println(string(ldata))

	// create lead

	rs2.Lead.Email = "kamal.saha@gmail.com"
	rs2.Lead.DisplayName = "Kamal Saha"
	rs2.Lead.FirstName = "Kamal"

	resp, err = client.R().
		SetBody(APIObject{Lead: rs2.Lead}).
		SetResult(&APIObject{}). // or SetResult(AuthSuccess{}).
		// SetError(&AuthError{}).       // or SetError(AuthError{}).
		Post("/api/leads")

	// :5023512614,

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
