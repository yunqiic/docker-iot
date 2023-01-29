```shell
https://www.codenong.com/b-kafka-connect-mqtt-mongodb/

cd docker
sudo docker-compose up -d

curl -d@<path-to-config-file>/connect-mqtt-source.json -H"Content-Type: application/json" -X POST http://localhost:8083/connectors

sudo docker run \
    -it --rm --name mqtt-publisher --network 04_custom_default \
    efrecon/mqtt-client \
    pub -h mosquitto -t"baeldung" -m"{"id":1234,"message":"This is a test"}"

sudo docker run \
    --rm \
    confluentinc/cp-kafka:5.1.0 \
    kafka-console-consumer --network 04_custom_default --bootstrap-server kafka:9092 --topic connect-custom --from-beginning

curl -d@<path-to-config file>/connect-mongodb-sink.json -H"Content-Type: application/json" -X POST http://localhost:8083/connectors

curl -X DELETE http://localhost:8083/connectors/mqtt-source
curl -X DELETE http://localhost:8083/connectors/mongodb-sink

mqtt -> kafka -> mongo
```