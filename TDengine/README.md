```shell
https://docs.taosdata.com/reference/docker/
sudo docker run -d --name tdengine -p 6041:6041 tdengine/tdengine
curl -u root:taosdata -d "show databases" localhost:6041/rest/sql
sudo docker exec -it tdengine taos
```