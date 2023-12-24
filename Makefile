clean:
	rm -rf deployments/data/kafka{1,2,3} && rm -rf deployments/data/zoo{1,2,3} && mkdir -p deployments/data/kafka{1,2,3} && mkdir -p deployments/data/zoo{1,2,3}

start_app: clean
	docker compose -f ./project/docker-compose.yml up -d

make_topic: start_app
	cd project/cmd/topic && go run topic.go

start: make_topic
	cd project/cmd/kafka_consumer && go run customer.go

get_kafdrop:
	wget -c https://github.com/obsidiandynamics/kafdrop/releases/download/4.0.1/kafdrop-4.0.1.jar

kafdrop: get_kafdrop
	java -jar kafdrop-4.0.1.jar --kafka.brokerConnect=127.0.0.1:29092


#Код с семинара с изменением