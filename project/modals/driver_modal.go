package modals

type Driver struct {
	Lat  float64 `db:"lat" json:"lat"`
	Lng  float64 `db:"lng" json:"lng"`
	ID   string  `db:"id" json:"id"`
	Name string  `db:"name" json:"name"`
	Auto string  `db:"auto" json:"auto"`
} // убрать db + протестить, из какого-то примера были
