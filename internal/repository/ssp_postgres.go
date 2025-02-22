package repository

import (
	"database/sql"
	"log"

	"github.com/segmentio/kafka-go"
)

func SspAddDeals(db *sql.DB, reader *kafka.Reader) {
	for {
		msgResp, err := ReadKafka(reader)
		if err != nil {
			continue
		}
		log.Println(msgResp)
		_, err = db.Exec("INSERT INTO tracked (bid_id, price, event_type) VALUES ($1, $2, $3)", msgResp.BidID, msgResp.Price, msgResp.EventType)
		if err != nil {
			continue
		}
	}
}
