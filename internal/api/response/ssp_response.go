package response

type SspResponse struct {
	DspID                 string `json:"dsp_id"`
	Price                 int64  `json:"price"`
	AdName                string `json:"ad_name"`
	TrackingClickURL      string `json:"tracking_click"`      //Ссылка на трекер+(инфа о размещении)Encoded(protobuff+(что нибудь)) Price/BidId/smth else
	TrackingImpressionURL string `json:"tracking_impression"` //Ссылка на трекер+(инфа о размещении)Encoded(protobuff+(что нибудь)) Price/BidId/smth else
}
