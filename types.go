package main

import "time"

type CompanyType struct {
	Partial  bool   `json:"partial"`
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type Company struct {
	ID                int64       `json:"id"`
	Name              string      `json:"name"`
	Address           string      `json:"address"`
	City              string      `json:"city"`
	State             string      `json:"state"`
	Zipcode           string      `json:"zipcode"`
	Country           string      `json:"country"`
	NumberOfEmployees int         `json:"number_of_employees"`
	AnnualRevenue     int         `json:"annual_revenue"`
	Website           string      `json:"website"`
	Phone             string      `json:"phone"`
	IndustryTypeID    int64       `json:"industry_type_id"`
	IndustryType      CompanyType `json:"industry_type"`
	BusinessTypeID    int64       `json:"business_type_id"`
	BusinessType      CompanyType `json:"business_type"`
}

type Currency struct {
	Partial      bool        `json:"partial"`
	ID           int64       `json:"id"`
	IsActive     bool        `json:"is_active"`
	CurrencyCode string      `json:"currency_code"`
	ExchangeRate string      `json:"exchange_rate"`
	CurrencyType int         `json:"currency_type"`
	ScheduleInfo interface{} `json:"schedule_info"`
}

type Deal struct {
	ID                 int64       `json:"id"`
	Name               string      `json:"name"`
	Amount             float64     `json:"amount"`
	CurrencyID         int64       `json:"currency_id"`
	BaseCurrencyAmount float64     `json:"base_currency_amount"`
	ExpectedClose      *time.Time  `json:"expected_close"`
	DealProductID      int64       `json:"deal_product_id"`
	DealProduct        interface{} `json:"deal_product"`
	Currency           Currency    `json:"currency"`
	ProductID          int         `json:"product_id"`
}

type Links struct {
	Conversations        string `json:"conversations"`
	TimelineFeeds        string `json:"timeline_feeds"`
	DocumentAssociations string `json:"document_associations"`
	Notes                string `json:"notes"`
	Tasks                string `json:"tasks"`
	Appointments         string `json:"appointments"`
	Reminders            string `json:"reminders"`
	Duplicates           string `json:"duplicates"`
	Connections          string `json:"connections"`
}

type EmailInfo struct {
	ID        int64       `json:"id"`
	Value     string      `json:"value"`
	IsPrimary bool        `json:"is_primary"`
	Label     interface{} `json:"label"`
	Destroy   bool        `json:"_destroy"`
}
type CustomFields struct {
	Interest              interface{} `json:"cf_interest"`
	Github                interface{} `json:"cf_github"`
	KubernetesSetup       string      `json:"cf_kubernetes_setup"`
	CalendlyMeetingAgenda interface{} `json:"cf_calendly_meeting_agenda"`
}
type Lead struct {
	ID                             int64         `json:"id"`
	JobTitle                       string        `json:"job_title"`
	Department                     string        `json:"department"`
	Email                          string        `json:"email"`
	Emails                         []EmailInfo   `json:"emails"`
	WorkNumber                     string        `json:"work_number"`
	MobileNumber                   string        `json:"mobile_number"`
	Address                        string        `json:"address"`
	City                           string        `json:"city"`
	State                          string        `json:"state"`
	Zipcode                        string        `json:"zipcode"`
	Country                        string        `json:"country"`
	TimeZone                       string        `json:"time_zone"`
	DoNotDisturb                   bool          `json:"do_not_disturb"`
	DisplayName                    string        `json:"display_name"`
	Avatar                         string        `json:"avatar"`
	Keyword                        string        `json:"keyword"`
	Medium                         string        `json:"medium"`
	LastSeen                       *time.Time    `json:"last_seen"`
	LastContacted                  *time.Time    `json:"last_contacted"`
	LeadScore                      int           `json:"lead_score"`
	LeadQuality                    string        `json:"lead_quality"`
	StageUpdatedTime               *time.Time    `json:"stage_updated_time"`
	FirstName                      string        `json:"first_name"`
	LastName                       string        `json:"last_name"`
	Company                        Company       `json:"company"`
	Deal                           Deal          `json:"deal"`
	Links                          Links         `json:"links"`
	CustomField                    CustomFields  `json:"custom_field"`
	CreatedAt                      string        `json:"created_at"`
	UpdatedAt                      string        `json:"updated_at"`
	LastContactedSalesActivityMode string        `json:"last_contacted_sales_activity_mode"`
	HasAuthority                   bool          `json:"has_authority"`
	EmailStatus                    string        `json:"email_status"`
	LastContactedMode              string        `json:"last_contacted_mode"`
	RecentNote                     string        `json:"recent_note"`
	LastContactedViaChat           *time.Time    `json:"last_contacted_via_chat"`
	LastContactedViaSalesActivity  string        `json:"last_contacted_via_sales_activity"`
	CompletedSalesSequences        int           `json:"completed_sales_sequences"`
	ActiveSalesSequences           int           `json:"active_sales_sequences"`
	WebFormIds                     string        `json:"web_form_ids"`
	LastAssignedAt                 *time.Time    `json:"last_assigned_at"`
	Tags                           []interface{} `json:"tags"`
	Facebook                       string        `json:"facebook"`
	Twitter                        string        `json:"twitter"`
	Linkedin                       string        `json:"linkedin"`
	IsDeleted                      bool          `json:"is_deleted"`
	TeamUserIds                    interface{}   `json:"team_user_ids"`
	SubscriptionStatus             int           `json:"subscription_status"`
	PhoneNumbers                   []interface{} `json:"phone_numbers"`
}
