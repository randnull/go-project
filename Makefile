clean:
	rm -rf deployments/data/kafka{1,2,3} && rm -rf deployments/data/zoo{1,2,3} && mkdir -p deployments/data/kafka{1,2,3} && mkdir -p deployments/data/zoo{1,2,3}

start_kafka: clean
	docker compose -f ./deployments/docker-compose.yml up -d

start_app: start_kafka
	docker compose -f ./project/docker-compose.yml up -d

#make_topic: start_app
#	go run project/cmd/topic/topic.go
#
#start: make_topic
#	go run project/cmd/kafka_consumer/customer.go

get_kafdrop:
	wget -c https://github.com/obsidiandynamics/kafdrop/releases/download/4.0.1/kafdrop-4.0.1.jar

kafdrop: get_kafdrop
	java -jar kafdrop-4.0.1.jar --kafka.brokerConnect=127.0.0.1:29092


#Код с семинара