package src

//https://github.com/googollee/go-gcm/blob/master/example/send.go
import (
	"appengine"
	"fmt"
	"github.com/googollee/go-gcm"
	"net/http"
)

func GCM(w http.ResponseWriter, r *http.Request) {
	_, c := getDefaultVar(r)
	if r.Method == "GET" {

		switch r.URL.Query().Get("action") {
		case "send":
			w.Header().Add("content-type", "text/html")
			fmt.Fprintf(w,
				"Sending a GCM : <br /><form method='post'>"+
					"<textarea name='key'></textarea><br />"+
					"<textarea name='message'></textarea><br />"+
					"<input type='submit' value='SubMeat' />"+
					"</form>")

		default:
			w.Header().Add("content-type", "text/html")
			fmt.Fprintf(w,
				"<form method='post'>"+
					"<textarea name='data'></textarea><br />"+
					"<input type='submit' value='SubMeat' />"+
					"</form>")
		}

	} else { //POST action
		switch r.URL.Query().Get("action") {
		case "send":
			SendGCM(r.FormValue("key"), map[string]string{"action": r.FormValue("message")}, "0", c)
			w.Header().Add("content-type", "text/html")
			fmt.Fprintf(w, "Sending "+r.FormValue("message")+" to : "+r.FormValue("key"))
		default:
			w.Header().Add("content-type", "text/html")

			data := DataSetToDB(StringToDataSet(r.FormValue("data")))
			StoreDataSet(c, data)
			fmt.Fprintf(w,
				"<form method='post'>"+
					"Mail : '"+data.Mail+"'<br />"+
					"Id : '"+data.Id+"'<br />"+
					"Type : '"+data.Type+"'<br />"+
					"Data : "+data.Data+"<br />"+
					"<textarea name='data'></textarea><br />"+
					"<input type='submit' value='Submeat' />"+
					"</form>")
		}
	}
}

func SendGCM(device_id string, data map[string]string, collapse string, context *appengine.Context) bool {
	client := gcm.New(getLastStoredKey(context))

	load := gcm.NewMessage(device_id)
	for k, v := range data {
		load.SetPayload(k, v)
	}
	load.CollapseKey = collapse
	load.DelayWhileIdle = true
	load.TimeToLive = 10

	_, err := client.Send(load)
	//context.Infof("ID" + string(resp.Success))

	/*	fmt.Printf("id: %+v\n", resp)
		fmt.Println("err:", err)
		fmt.Println("err index:", resp.ErrorIndexes())
		fmt.Println("reg index:", resp.RefreshIndexes())*/
	if err != nil {
		return false
	}
	return true
}
