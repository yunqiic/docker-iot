```shell
cd docker
sudo docker-compose up -d
chmod -R 777 emqx1_data/
chmod -R 777 emqx2_data/
sudo docker-compose up -d
sudo docker-compose ps
sudo docker-compose logs -f
sudo docker exec -it emqx1 sh -c "emqx_ctl cluster status"

sudo docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:latest

sudo docker run -d --name emqx -p 1883:1883 -p 8024:8083 -p 8025:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:latest

sudo docker stop emqx
sudo docker rm emqx

sudo docker exec -it emqx bash
sudo docker cp emqx:/opt/emqx/etc etc

./bin/emqx_ctl users

users list                                     
users add <Username> <Password>                
users update <Username> <NewPassword>          
users del <Username>

ws:8024
wss:8025

49.232.6.131
1883
abc
abc
Topic:t/1
{
  "bar": "0.2"
}

http://49.232.6.131:18083/
admin
public

admin
public0128

test
test

Remark
administrator
viewer

https://www.emqx.com/zh/blog/store-mqtt-messages-to-time-series-database-iotdb
# rule
SELECT
  payload.msg as msg
FROM
  "t/#"
WHERE
  msg = 'hello'

SELECT
    clientid,
    now_timestamp('millisecond') as now_ts_ms,
    payload.bar as bar
FROM
    "t/#"

MQTT Bridge
ClientId
root
Username
root
Password
root

Forward Topic 
foo

Payload Template
{
 "device": "root.sg.${clientid}",
 "timestamp": ${now_ts_ms},
 "measurements": [
   "bar"
 ],
 "values": [
   ${bar}
 ]
}
```

```
https://www.emqx.io/docs/zh/v5/deploy/install-docker.html
sudo docker pull emqx:5.0.14
sudo docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083  emqx:5.0.14
```

```
https://www.emqx.com/zh/blog/store-mqtt-messages-to-time-series-database-iotdb
https://juejin.cn/post/6992031751790919688

https://www.cnblogs.com/emqx/p/11400663.html
$ ./bin/emqx_ctl users

users list
users add <Username> <Password>
users update <Username> <NewPassword>
users del <Username>

# Request
GET api/v3/auth_username

# Response
{
   "code": 0,
   "data": ["username1"]
}


# Request
POST api/v3/auth_username
{
   "username": "some_name",
   "password": "password"
}

# Response
{
   "code": 0
}

# Request
PUT api/v3/auth_username/$NAME
{
   "password": "password"
}

# Response
{
   "code", 0
}


# Request
GET api/v3/auth_username/$NAME

# Response
{
   "code": 0,
   "data": {
       "username": "some_username",
       "password": "hashed_password"
   }
}

# Request
DELETE api/v3/auth_username/$NAME

# Response
{
   "code": 0
}
```