package src

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"strings"
)

type DataSet struct {
	Mail string
	Id   string
	Type string
	Data map[string]string
}

//can't store map in db, keeping it in JSON
type DataSetDB struct {
	Mail string
	Id   string
	Type string
	Data string
}

func StoreDataSet(c appengine.Context, set DataSetDB) interface{} {
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DataSetDB", nil), &set)
	if err != nil {
		return false
	}
	return key
}

func getStoredData(c appengine.Context, ty string) {
	//c.
}

func StringToDataSet(s string) DataSet {
	var data DataSet
	dec := json.NewDecoder(strings.NewReader(s))
	dec.Decode(&data)
	return data
}

func DataSetToDB(data DataSet) DataSetDB {
	dastr, _ := json.Marshal(data.Data)
	return DataSetDB{
		Mail: data.Mail,
		Id:   data.Id,
		Type: data.Type,
		Data: string(dastr),
	}
}

func DataSetDbToDataSet(data DataSetDB) DataSet {
	var da map[string]string
	dec := json.NewDecoder(strings.NewReader(data.Data))
	dec.Decode(&da)
	return DataSet{
		Mail: data.Mail,
		Id:   data.Id,
		Type: data.Type,
		Data: da,
	}
}
