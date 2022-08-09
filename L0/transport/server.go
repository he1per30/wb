package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"L0/data/pg"
	"L0/data/redis"
)

//

type reqstOt struct {
	ID string `json:"id"`
}

func NewServer() {
	http.HandleFunc("/get", GetByID)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func GetByID(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	b, _ := ioutil.ReadAll(req.Body)

	var r reqstOt

	_ = json.Unmarshal(b, &r)

	res, err := redis.Get(r.ID)
	//if err != nil {
	//	fmt.Println(err)
	//	resp.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	if res != "" {
		_, err = resp.Write([]byte(res))
		if err != nil {
			fmt.Println("error when write resp", err)
			return
		}
		return
	}

	order, err := pg.GetOrder(r.ID)
	if err != nil {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resB, _ := json.Marshal(order)
	_, err = resp.Write(resB)
	if err != nil {
		fmt.Println("error when write resp", err)
	}

	err = redis.Set(order.OrderUid, res)
	if err != nil {
		fmt.Println("error when saving to cache", err)
	}
}
