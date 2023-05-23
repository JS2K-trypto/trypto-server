package util

import (
	"encoding/json"
	"trypto-server/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ErrorHandler(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func GetJsonIdKeyValue(data []byte) (primitive.ObjectID, string, interface{}) {
	var unMarshared map[string]interface{}
	err := json.Unmarshal(data, &unMarshared)
	ErrorHandler(err)

	var id primitive.ObjectID
	var key string
	var value interface{}

	// id로 구분하려면, primitive.ObjectIDFromHex()함수를 사용해 형변환을 해줘야한다.
	if _, exist := unMarshared["id"]; exist {
		id, _ = primitive.ObjectIDFromHex(unMarshared["id"].(string))
	}
	if _, exist := unMarshared["key"]; exist {
		key = unMarshared["key"].(string)
	}
	if _, exist := unMarshared["value"]; exist {
		value = unMarshared["value"]
	}

	return id, key, value
}

func ConvertStringToObjectId(sId string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(sId)
	ErrorHandler(err)

	return id
}
