package response

type DspResponse struct {
	ID     string `json:"id"`
	DspID  string `json:"dsp_id"`
	BidID  string `json:"bid_id"`
	AdName string `json:"ad_name"`
	Price  int64  `json:"price"`
}
