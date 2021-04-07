docker-compose up -d  
docker-compose ps

docker exec -it kafka_kafka_1 bash
docker exec -it simulator bash

kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
