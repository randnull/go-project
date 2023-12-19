package driver_app

func main() {
	rep, _ := NewRepository("mongodb://127.0.0.1:27017")
	serv, _ := NewDriverService()
}
