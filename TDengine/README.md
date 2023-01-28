```shell
https://docs.taosdata.com/reference/docker/
sudo docker run -d --name tdengine -p 6041:6041 tdengine/tdengine
curl -u root:taosdata -d "show databases" localhost:6041/rest/sql
sudo docker exec -it tdengine taos

sudo docker run -d --name tdengine --network host tdengine/tdengine
sudo docker run -d \
   --name tdengine \
   -e TAOS_FQDN=tdengine \
   -p 6030:6030 \
   -p 6041-6049:6041-6049 \
   -p 6041-6049:6041-6049/udp \
   tdengine/tdengine

echo 127.0.0.1 tdengine |sudo tee -a /etc/hosts
taos -h tdengine -P 6030

sudo docker network create td-net
sudo docker run -d --name tdengine --network td-net \
   -e TAOS_FQDN=tdengine \
   tdengine/tdengine

sudo docker run --rm -it --network td-net -e TAOS_FIRST_EP=tdengine --entrypoint=taos tdengine/tdengine
# or
#sudo docker run --rm -it --network td-net --entrypoint=taos tdengine/tdengine -h tdengine
```