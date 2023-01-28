```shell
sudo docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:latest

sudo docker run -d --name emqx -p 1883:1883 -p 8024:8083 -p 8025:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:latest

http://49.232.6.131:18083/
admin
public

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
```