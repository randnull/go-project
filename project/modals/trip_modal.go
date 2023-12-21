package modals

type Trip struct {
	ID       string        `json:"id"`
	DriverID string        `json:"driver_id"`
	UserId   string        `json:"user_id"`
	From     Latlngtiteral `json:"from"`
	To       Latlngtiteral `json:"to"`
	Price    Money         `json:"price"`
	Status   string        `json:"status"`
}
