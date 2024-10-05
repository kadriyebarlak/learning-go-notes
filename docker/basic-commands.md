 1048  docker rmi alpine
 1049  docker images
 1050  docker run --rm -i -t nginx
 1051  docker run --rm -i -t -p 8080:80 nginx
 1052  docker run -d -p 8080:80 --restart=always --name=web nginx
 1053  docker ps
 1055  docker logs --help
 1056  docker logs -f --since 10s
 1057  docker logs -f --since 10s web
 1058  docker ps
 1059  docker exec -i -t 715e4d512e9a bash
 1060  docker ps
 1061  docker rm -f web
 1062  docker ps



 docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
 docker run -p 6379:6379 redislabs/rebloom:latest

* Kafka
 docker-compose up -d
 docker-compose stop
 docker exec -it docker-kafka-kafka-1 /bin/sh
 kafka-topics --describe --bootstrap-server localhost:9092 --topic myTopic (partition)




