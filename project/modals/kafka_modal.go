package modals

import "time"

type TripData struct {
	TripID   string `json:"trip_id"`
	DriverID string `json:"driver_id"`
}

type KafkaMessage struct {
	ID              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	DataContentType string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            TripData  `json:"data"`
}
