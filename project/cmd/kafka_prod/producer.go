package main

//func main() {
//	ctx := context.Background()
//
//	logger := log.Default()
//
//	writer := kafka.NewWriter(kafka.WriterConfig{
//		Brokers:     []string{"127.0.0.1:29092"},
//		Topic:       "demo",
//		Async:       true,
//		Logger:      kafka.LoggerFunc(logger.Printf),
//		ErrorLogger: kafka.LoggerFunc(logger.Printf),
//		BatchSize:   2000,
//	})
//	defer writer.Close()
//
//	//m
//
//	messageBytes, err := json.Marshal(m)
//	err = writer.WriteMessages(ctx, kafka.Message{Key: []byte(strconv.Itoa(0)), Value: messageBytes})
//	if err != nil {
//		log.Fatal(err)
//	}
//}
