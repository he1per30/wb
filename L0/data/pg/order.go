package pg

import (
	"L0/data/model"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type orderDB struct {
	Order_uid          string
	Track_number       string
	Entry              string
	Delivery           string
	Payment            string
	Order_item         string
	Locale             string
	Internal_signature string
	Delivery_service   string
	Shardkey           string
	Sm_id              string
	Date_created       string
	Oof_shard          string
}

func AddOrder(dataFromNats model.Order) error {
	deliveryBt, _ := json.Marshal(dataFromNats.Delivery)

	_, err := db.Exec(`
	INSERT INTO list_order
	    (order_uid,      track_number,            entry, delivery, payment, order_item, 
	        locale,internal_signature, delivery_service, shardkey,   sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		dataFromNats.OrderUid, dataFromNats.TrackNumber, dataFromNats.Entry, deliveryBt,
		dataFromNats.Payment, dataFromNats.OrderItem, dataFromNats.Locale, dataFromNats.InternalSignature,
		dataFromNats.DeliveryService, dataFromNats.Shardkey, dataFromNats.SmId, dataFromNats.DateCreated,
		dataFromNats.OofShard)

	return err
}

func GetOrder(uuid string) (model.Order, error) {
	var orderData orderDB
	var orderModel model.Order

	s := fmt.Sprintf(`SELECT * FROM list_order WHERE order_uid = '%s'`, uuid)

	err := db.Get(&orderData, s)
	if err != nil {
		fmt.Printf("error while get from pg: %s", err)
	}

	byteData, _ := json.Marshal(orderData)

	err = json.Unmarshal(byteData, &orderModel)
	return orderModel, nil
}

// Вариант если нужно будет сразу много сохранять
//queryInsert := `INSERT INTO list_order (order_uid, track_number, entry, delivery, payment, order_item,
//locale,internal_signature, delivery_service, shardkey, sm_id, date_created, oof_shard)
//VALUES `
//
//insertparams := []interface{}{}
//
//for i, result := range dataFromNats {
//p1 := i * 13 // starting position for insert params
//
//queryInsert += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),",
//p1+1, p1+2, p1+3, p1+4, p1+5, p1+6, p1+7, p1+8, p1+9, p1+10, p1+11, p1+12, p1+13)
//
//insertparams = append(insertparams, result.OrderUid, result.TrackNumber, result.Entry, result.Delivery,
//result.Payment, result.OrderItem, result.Locale, result.InternalSignature,
//result.DeliveryService, result.Shardkey, result.SmId, result.DateCreated,
//result.OofShard)
//}
//
//queryInsert = queryInsert[:len(queryInsert)-1] // remove trailing ","
//
//db.MustExec(queryInsert, insertparams...)
