package presenters

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func StructToBson(structObject interface{}) (bson.M, error) {
	var bsonToReturn bson.M
	data, err := bson.Marshal(structObject)
	if err != nil {
		return bsonToReturn, err
	}

	errUnmarshaling := bson.Unmarshal(data, &bsonToReturn)
	if err != nil {
		return bsonToReturn, errUnmarshaling
	}
	return bsonToReturn, nil
}

func ArrayStructToBson(array, outArray interface{}) error {
	inStructArrData, err := bson.Marshal(array)
	if err != nil {
		return err
	}
	// kind 4 for array
	raw := bson.Raw{Kind: 4, Data: inStructArrData}

	return raw.Unmarshal(outArray)
}

func GetTimeNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func ReturnHttpPayload(object interface{}, w http.ResponseWriter) {
	payload, _ := json.Marshal(object)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "limit, includes, skip, sort, filter, select, self, destination, Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(payload))
}

func ReturnHttpError(err error, w http.ResponseWriter, httpErrType int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "limit, includes, skip, sort, filter, select, destination, self, Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token	")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErrType)
	w.Write([]byte("{\"success\":false,\"error\": \"" + err.Error() + "\"}"))
}

func GetRequestValue(key string, r *http.Request) string {
	params := mux.Vars(r)
	return params[key]
}

const (
	InstanceStatusActive   string = "ACTIVE"
	InstanceStatusInactive string = "INACTIVE"
	InstanceStatusDeleted  string = "DELETED"
)

type Instance struct {
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt int64  `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
}

func CreateInstance() Instance {
	var instance Instance
	instance.Status = InstanceStatusActive
	instance.ModifiedAt = GetTimeNow()
	instance.CreatedAt = GetTimeNow()

	return instance
}
