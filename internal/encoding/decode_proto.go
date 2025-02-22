package encoding

import (
	"encoding/base64"

	"google.golang.org/protobuf/proto"
)

func Decode(encoded string) (*Tracking, error) {

	data, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	var decoded Tracking

	err = proto.Unmarshal(data, &decoded)
	if err != nil {
		return nil, err
	}

	return &decoded, nil
}
