package encoding

import (
	"encoding/base64"
	"test_project/internal/api/response"

	"google.golang.org/protobuf/proto"
)

func Encode(resp *response.DspResponse) (string, error) {
	data := &Tracking{
		DspId:  resp.DspID,
		AdName: resp.AdName,
		BidID:  resp.BidID,
		Price:  resp.Price,
	}

	dataMarshalled, err := proto.Marshal(data)
	if err != nil {
		return "", err
	}
	dataString := base64.URLEncoding.EncodeToString(dataMarshalled)
	return dataString, nil

}
