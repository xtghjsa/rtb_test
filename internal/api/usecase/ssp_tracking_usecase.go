package usecase

import (
	"log"
	"test_project/internal/encoding"
	"time"
)

type TrackingUsecaseInt interface {
	WriteKafka(decoded *encoding.Tracking, eventType string) error
}

type TrackingUsecase struct {
	Repo TrackingUsecaseInt
}

func (u *TrackingUsecase) TrackingExec(toDecode string, eventType string) bool {
	log.Println("Decoding data", time.Now())
	decoded, err := encoding.Decode(toDecode)
	if err != nil {
		return false
	}
	log.Println("Writing to kafka", time.Now())
	err = u.Repo.WriteKafka(decoded, eventType)
	if err != nil {
		log.Printf("Error writing to kafka: %v", err)
		return false
	}
	log.Println("Data written to kafka", time.Now())
	return true
}
