package src

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type ApiKey struct {
	Key  string
	Date time.Time
}

func StoreApiKey(key string, context *appengine.Context) bool {
	data := ApiKey{
		Key:  key,
		Date: time.Now(),
	}
	_, err := datastore.Put(*(context), datastore.NewIncompleteKey(*(context), "ApiKey", nil), &data)
	if err != nil {
		return false
	}
	return true
}

func getLastStoredKey(context *appengine.Context) string {
	var api []ApiKey
	ret := ""
	q := datastore.NewQuery("ApiKey").Order("-Date").Limit(1)
	q.GetAll(*(context), &api)
	if len(api) > 0 {
		ret = api[0].Key
	}
	return ret
}
