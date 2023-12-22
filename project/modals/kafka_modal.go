package modals

import "time"

type TripData struct {
	TripID  string        `json:"trip_id"`
	OfferID string        `json:"offer_id"`
	Price   Money         `json:"price"`
	Status  string        `json:"status"`
	From    Latlngtiteral `json:"from"`
	To      Latlngtiteral `json:"to"`
}

type KafkaMessage struct {
	ID              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	DataContentType string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            TripData  `json:"data"`
}
