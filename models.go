package main

type LastRechargeRequest struct {
	RequestId string `validate:"required"`
	Msisdn    string `validate:"required"`
	Channel   string
}

type MobileLastRechargeResponse struct {
	RequestId string
	Msisdn    string
	Recharge  string
	Date      string
}

type FixedLastRechargeResponse struct {
	RequestId string
	Msisdn    string
	Mmsisdn   string
	Recharge  string
	Date      string
	Channel   string
}
