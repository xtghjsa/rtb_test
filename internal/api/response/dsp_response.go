package response

type DspResponse struct {
	ID          string `json:"id"`
	DspID       string `json:"dsp_id"`
	AdName      string `json:"ad_name"`
	AdCondition string `json:"ad_condition"`
	Price       int    `json:"price"`
}
