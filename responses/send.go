package responses

import (
	"reflect"

	"github.com/halorium/networks-example/logger"
)

func (res *Response) Send() {
	logger.Debug("response-send", res.State)

	if res.Sent {
		return
	}

	res.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if res.Head != nil {
		reflectValue := reflect.ValueOf(res.Head).Elem()
		for i := 0; i < reflectValue.NumField(); i++ {
			key := reflectValue.Type().Field(i).Name
			val := reflectValue.Field(i).String()
			if val != "" {
				key = FormatHeaderKey("IN", key)
				res.Writer.Header().Set(key, val)
			}
		}
	}

	res.Writer.WriteHeader(res.HTTPStatusCode)

	if res.Body != nil {
		body := res.Serializer(res.Body)

		_, err := res.Writer.Write(body)

		if err != nil {
			res.Body = err
			logger.Panic("response-send", res.State)
		}
	}
}
