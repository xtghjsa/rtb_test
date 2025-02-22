package response

type SspResponse struct {
	DspID                 string `json:"dsp_id"`
	Price                 int64  `json:"price"`
	AdName                string `json:"ad_name"`
	TrackingClickURL      string `json:"tracking_click"`
	TrackingImpressionURL string `json:"tracking_impression"`
}
