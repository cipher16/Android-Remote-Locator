package src

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"strings"
	"time"
)

type DataSet struct {
	Mail string
	Id   string
	Type string
	IMEI string
	Date time.Time
	Data map[string]string
}

//can't store map in db, keeping it in JSON
type DataSetDB struct {
	Mail string
	Id   string
	Type string
	IMEI string
	Date time.Time
	Data string
}

func StoreDataSet(c *appengine.Context, set *DataSetDB) interface{} {
	set.Date = time.Now()
	key, err := datastore.Put(*(c), datastore.NewIncompleteKey(*(c), "DataSetDB", nil), set)
	if err != nil {
		return false
	}
	return key
}

func get10LastType(ty string, mail string, context *appengine.Context) []DataSet {
	var data []DataSetDB
	var dataset []DataSet
	q := datastore.NewQuery("DataSetDB").Filter("Mail=", mail).Filter("Type =", ty).Order("-Date").Limit(10)
	q.GetAll(*(context), &data)
	for _, v := range data {
		dataset = append(dataset, DataSetDbToDataSet(v))
	}
	return dataset
}

func getPhoneForAccount(mail string, context *appengine.Context) []DataSet {
	//no distinct raaaaaaaa
	var data []DataSetDB
	var dataset []DataSet
	q := datastore.NewQuery("DataSetDB").Filter("Mail=", mail).Order("-Date").Limit(10)
	q.GetAll(*(context), &data)
	for _, v := range data {
		dataset = append(dataset, DataSetDbToDataSet(v))
	}
	return dataset
}

func StringToDataSet(s string) *DataSet {
	var data DataSet
	dec := json.NewDecoder(strings.NewReader(s))
	dec.Decode(&data)
	return &data
}

func DataSetToDB(data *DataSet) *DataSetDB {
	dastr, _ := json.Marshal(data.Data)
	return &DataSetDB{
		Mail: data.Mail,
		Id:   data.Id,
		Type: data.Type,
		IMEI: data.IMEI,
		Date: data.Date,
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
		IMEI: data.IMEI,
		Date: data.Date,
		Data: da,
	}
}
