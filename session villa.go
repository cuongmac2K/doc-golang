package main

import "time"

func abc() {

}

type T struct {
	UrlType        string      `json:"url_type"`
	Origin         interface{} `json:"origin"`
	ClientType     interface{} `json:"client_type"`
	LogoUrl        interface{} `json:"logo_url"`
	BillingBalance int         `json:"billing_balance"`
	Balances       struct {
		Primary           float64 `json:"primary"`
		PromotionGeneral  float64 `json:"promotion_general"`
		CloudServer       float64 `json:"cloud_server"`
		SimpleStorage     float64 `json:"simple_storage"`
		MailInbox         float64 `json:"mail_inbox"`
		Cdn               float64 `json:"cdn"`
		LoadBalancer      float64 `json:"load_balancer"`
		Sysadmin          float64 `json:"sysadmin"`
		Vpn               float64 `json:"vpn"`
		Pentest           float64 `json:"pentest"`
		CloudStorage      float64 `json:"cloud_storage"`
		CallCenter        float64 `json:"call_center"`
		Security          float64 `json:"security"`
		Vod               float64 `json:"vod"`
		FiberP2P          float64 `json:"fiber_p2p"`
		ContainerRegistry float64 `json:"container_registry"`
		Ddos              float64 `json:"ddos"`
		AutoScaling       float64 `json:"auto_scaling"`
		KubernetesEngine  float64 `json:"kubernetes_engine"`
		CloudWatcher      float64 `json:"cloud_watcher"`
		Livestream        float64 `json:"livestream"`
		Dbaas             float64 `json:"dbaas"`
		ApiGateway        float64 `json:"api_gateway"`
		TrafficManager    float64 `json:"traffic_manager"`
		Lms               float64 `json:"lms"`
		AppEngine         float64 `json:"app_engine"`
	} `json:"balances"`
	PaymentMethod   string        `json:"payment_method"`
	BillingAccId    string        `json:"billing_acc_id"`
	Debit           interface{}   `json:"debit"`
	Email           string        `json:"email"`
	Phone           interface{}   `json:"phone"`
	Fullname        string        `json:"fullname"`
	VerifiedEmail   interface{}   `json:"verified_email"`
	VerifiedPhone   interface{}   `json:"verified_phone"`
	LoginAlert      interface{}   `json:"login_alert"`
	VerifiedPayment interface{}   `json:"verified_payment"`
	LastRegion      string        `json:"last_region"`
	LastProject     string        `json:"last_project"`
	Type            string        `json:"type"`
	Otp             interface{}   `json:"otp"`
	Services        []interface{} `json:"services"`
	Whitelist       []interface{} `json:"whitelist"`
	Token           string        `json:"token"`
	Expires         time.Time     `json:"expires"`
	TenantId        string        `json:"tenant_id"`
	TenantName      string        `json:"tenant_name"`
	KsUserId        string        `json:"ks_user_id"`
	Roles           []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"roles"`
	Iam struct {
		Token           string        `json:"token"`
		Expire          time.Time     `json:"expire"`
		TenantId        string        `json:"tenant_id"`
		TenantName      string        `json:"tenant_name"`
		VerifiedPhone   interface{}   `json:"verified_phone"`
		VerifiedEmail   interface{}   `json:"verified_email"`
		VerifiedPayment interface{}   `json:"verified_payment"`
		PaymentMethod   string        `json:"payment_method"`
		Whitelist       []interface{} `json:"whitelist"`
		Roles           []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"roles"`
	} `json:"iam"`
	OtpEnabled interface{} `json:"otp_enabled"`
	Owner      string      `json:"owner"`
	Domain     struct {
	} `json:"domain"`
	PaymentType string        `json:"payment_type"`
	Avatar      string        `json:"avatar"`
	Dob         string        `json:"dob"`
	Gender      string        `json:"_gender"`
	UserRegions []interface{} `json:"user_regions"`
	Ip          string        `json:"ip"`
	UserAgent   string        `json:"user-agent"`
}
