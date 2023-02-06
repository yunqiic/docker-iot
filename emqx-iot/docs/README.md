```shell
sudo docker pull emqx/emqx-ee:4.4.14
sudo docker run -d --name emqx-ee -p 1883:1883 -p 8081:8081 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx-ee:4.4.14

sudo docker pull emqx/emqx:5.0.15
sudo docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:5.0.15

https://www.taosdata.com/docs/cn/v2.0/connector/java
https://github.com/taosdata/taos-connector-jdbc

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

SELECT
  payload
FROM
  "sensor/data"
{
  "temperature":30,
  "humidity":20,
  "volume":44.5,
  "PM10":23,
  "pm25":61,
  "SO2":14,
  "NO2":4,
  "CO":5,
  "id":"10-c6-1f-1a-1f-47",
  "area":1,
  "ts":1596157444170
}

t/1
{
  "bar": "0.2"
}

https://www.emqx.io/docs/zh/v5/
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
);
INSERT INTO test.sensor_data VALUES(
  now,
  1,
  1,
  1,
  1,
  1,
  1,
  1,
  1,
  '1',
  1,
  now
);
INSERT INTO test.sensor_data VALUES(
  now,
  1,
  1,
  1,
  1,
  1,
  1,
  1,
  1,
  '${payload.id}',
  1,
  now
);
select * from sensor_data;

tdengnine root 密码
SHOW USERS;
ALTER USER <user_name> PRIVILEGE <write|read>;

root:taosdata
cm9vdDp0YW9zZGF0YQ==
test:123456
dGVzdDoxMjM0NTY=

curl -L -H "Authorization: Basic cm9vdDp0YW9zZGF0YQ==" \
  -d "select name, ntables, status from information_schema.ins_databases;" \
  172.21.16.11:6041/rest/sql

curl -L -H "Authorization: Basic cm9vdDp0YW9zZGF0YQ==" \
  -d " select * from test.sensor_data;" \
  172.21.16.11:6041/rest/sql

curl -L -H "Authorization: Basic dGVzdDoxMjM0NTY=" \
  -d " select * from test.sensor_data;" \
  172.21.16.11:6041/rest/sql

404 API_NOT_EXIST:Request Path Not Found
{"code": "API_NOT_EXIST", "message": "Request Path Not Found"}

alter user root pass '123456';
create user test pass '123456';
drop user test;

tdengine
http://172.21.16.11:6041/rest/sql
POST
Authorization 的值为 Basic + TDengine 的 {username}:{password} 经过 Base64 编码之后的字符串, 例如 root:taosdata 编码后实际填入的值为：Basic cm9vdDp0YW9zZGF0YQ==
Authorization Basic cm9vdDp0YW9zZGF0YQ==

mqtt://49.232.6.131:1883
npm install mqtt mockjs --save --registry=https://registry.npm.taobao.org
cnpm install mqtt mockjs --save
https://github.com/coco-iot/MQTT.js
node mock.js

sudo docker run -d --name=grafana -p 3000:3000 grafana/grafana:9.3.6
sudo docker stop grafana
sudo docker start grafana
sudo docker rm grafana
http://49.232.6.131:3000/
admin
admin
admin2023

TDengine plugin
https://grafana.com/grafana/plugins/tdengine-datasource/

docker exec -it grafana bash
bash -c "$(curl -fsSL \
  https://raw.githubusercontent.com/taosdata/grafanaplugin/master/install.sh)" -- \
  -a http://172.21.16.11:6041 \
  -u root \
  -p taosdata

http://172.21.16.11:6041
root
taosdata

sudo docker exec -it tdengine bash
taos
select avg(temperature) as temperature from test.sensor_data;
select count(1) from test.sensor_data;

select avg(temperature), avg(humidity), avg(volume), avg(PM10), avg(pm25), avg(SO2), avg(NO2), avg(CO)  from test.sensor_data where coll_time >= $from and coll_time < $to interval($interval)

select max(temperature), max(humidity), max(volume), max(PM10), max(pm25), max(SO2), max(NO2), max(CO), min(temperature), min(humidity), min(volume), min(PM10), min(pm25), min(SO2), min(NO2), min(CO)  from test.sensor_data where coll_time >= $from and coll_time < $to interval($interval)
```

```
The arguments are:

streamBuilder is a function that returns a subclass of the Stream class that supports the connect event. Typically a net.Socket.
options is the client connection options (see: the connect packet). Defaults:
wsOptions: is the WebSocket connection options. Default is {}. It's specific for WebSockets. For possible options have a look at: https://github.com/websockets/ws/blob/master/doc/ws.md.
keepalive: 60 seconds, set to 0 to disable
reschedulePings: reschedule ping messages after sending packets (default true)
clientId: 'mqttjs_' + Math.random().toString(16).substr(2, 8)
protocolId: 'MQTT'
protocolVersion: 4
clean: true, set to false to receive QoS 1 and 2 messages while offline
reconnectPeriod: 1000 milliseconds, interval between two reconnections. Disable auto reconnect by setting to 0.
connectTimeout: 30 * 1000 milliseconds, time to wait before a CONNACK is received
username: the username required by your broker, if any
password: the password required by your broker, if any
incomingStore: a Store for the incoming packets
outgoingStore: a Store for the outgoing packets
queueQoSZero: if connection is broken, queue outgoing QoS zero messages (default true)
customHandleAcks: MQTT 5 feature of custom handling puback and pubrec packets. Its callback:
  customHandleAcks: function(topic, message, packet, done) {/*some logic wit colling done(error, reasonCode)*/}
autoUseTopicAlias: enabling automatic Topic Alias using functionality
autoAssignTopicAlias: enabling automatic Topic Alias assign functionality
properties: properties MQTT 5.0. object that supports the following properties:
sessionExpiryInterval: representing the Session Expiry Interval in seconds number,
receiveMaximum: representing the Receive Maximum value number,
maximumPacketSize: representing the Maximum Packet Size the Client is willing to accept number,
topicAliasMaximum: representing the Topic Alias Maximum value indicates the highest value that the Client will accept as a Topic Alias sent by the Server number,
requestResponseInformation: The Client uses this value to request the Server to return Response Information in the CONNACK boolean,
requestProblemInformation: The Client uses this value to indicate whether the Reason String or User Properties are sent in the case of failures boolean,
userProperties: The User Property is allowed to appear multiple times to represent multiple name, value pairs object,
authenticationMethod: the name of the authentication method used for extended authentication string,
authenticationData: Binary Data containing authentication data binary
authPacket: settings for auth packet object
will: a message that will sent by the broker automatically when the client disconnect badly. The format is:
topic: the topic to publish
payload: the message to publish
qos: the QoS
retain: the retain flag
properties: properties of will by MQTT 5.0:
willDelayInterval: representing the Will Delay Interval in seconds number,
payloadFormatIndicator: Will Message is UTF-8 Encoded Character Data or not boolean,
messageExpiryInterval: value is the lifetime of the Will Message in seconds and is sent as the Publication Expiry Interval when the Server publishes the Will Message number,
contentType: describing the content of the Will Message string,
responseTopic: String which is used as the Topic Name for a response message string,
correlationData: The Correlation Data is used by the sender of the Request Message to identify which request the Response Message is for when it is received binary,
userProperties: The User Property is allowed to appear multiple times to represent multiple name, value pairs object
transformWsUrl : optional (url, options, client) => url function For ws/wss protocols only. Can be used to implement signing urls which upon reconnect can have become expired.
resubscribe : if connection is broken and reconnects, subscribed topics are automatically subscribed again (default true)
messageIdProvider: custom messageId provider. when new UniqueMessageIdProvider() is set, then non conflict messageId is provided.
```