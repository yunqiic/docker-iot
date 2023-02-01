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

t/1
{
  "bar": "0.2"
}

INSERT INTO test.sensor_data VALUES(
  now,
  ${payload.temperature},
  ${payload.humidity},
  ${payload.volume},
  ${payload.PM10},
  ${payload.pm25},
  ${payload.SO2},
  ${payload.NO2},
  ${payload.CO},
  '${payload.id}',
  ${payload.area},
  ${payload.ts}
)
select * from sensor_data;

tdengnine root 密码
SHOW USERS;
ALTER USER <user_name> PRIVILEGE <write|read>;

alter user root pass '123456';
create user test pass '123456';
drop user test;

tdengine
http://172.21.16.11:6041/rest/sql
POST
Authorization 的值为 Basic + TDengine 的 {username}:{password} 经过 Base64 编码之后的字符串, 例如 root:taosdata 编码后实际填入的值为：Basic cm9vdDp0YW9zZGF0YQ==
Authorization Basic cm9vdDp0YW9zZGF0YQ==
```