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

type Req struct {
	Lat    float64 `json:"Lat"`
	Lng    float64 `json:"Lng"`
	Radius float64 `json:"radius"`
}

type TripAnswerData struct {
	TripID string `json:"trip_id"`
}

type TripAnswer struct {
	ID              string         `json:"id"`
	Source          string         `json:"source"`
	Type            string         `json:"type"`
	DataContentType string         `json:"datacontenttype"`
	Time            time.Time      `json:"time"`
	Data            TripAnswerData `json:"data"`
}
