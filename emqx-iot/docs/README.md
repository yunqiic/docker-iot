```shell
mysql -h127.0.0.1 -uroot -p -P3311
root

/etc/emqx/plugins/emqx_backend_mysql.conf
backend.mysql.server = 127.0.0.1:3306

backend.mysql.username = root

backend.mysql.password = 123456

backend.mysql.database = mqtt

emqx_ctl plugins load emqx_backend_mysql

emqx plugins

https://www.emqx.com/zh/blog/emqx-plugin-persistence-series-5-mysql-mqtt-data-storage

emqx tdengine
https://www.emqx.com/zh/blog/emqx-tdengine-grafana
emqx 4.1.1

## 拉取并启动容器
sudo docker run -d --name tdengine -p 6030-6041:6030-6041 tdengine/tdengine:latest

## 启动后检查容器运行状态
sudo docker ps|grep tdengine
sudo docker ps -a

docker stop tdengine
docker rm tdengine

sudo docker run -d --name=grafana -p 3000:3000 grafana/grafana
docker stop grafana
docker rm grafana

sudo docker exec -it tdengine bash
taos
create database test;

use test;
CREATE TABLE sensor_data (
  ts timestamp,
  temperature float,
  humidity float,
  volume float,
  PM10 float,
  pm25 float,
  SO2 float,
  NO2 float,
  CO float,
  sensor_id NCHAR(255), 
  area TINYINT,
  coll_time timestamp
);

# payload

SELECT
  payload
FROM
  "sensor/data"

Test

topic
sensor/data
{
  "payload": "{\"temperature\":30,\"humidity\":20,\"volume\":44.5,\"PM10\":23,\"pm2.5\":61,\"SO2\":14,\"NO2\":4,\"CO\":5,\"id\":\"10-c6-1f-1a-1f-47\",\"area\":1,\"ts\":1596157444170}"
}
```