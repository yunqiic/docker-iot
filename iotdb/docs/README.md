```shell
Using Docker：The path to the dockerfile is https://github.com/apache/iotdb/tree/master/docker/src/main
mvn clean package -pl distribution -am -DskipTests

maven:3.8.6-openjdk-11
yiluxiangbei/maven:3.8.6-openjdk-11

cd git/iotdb
sudo docker run -it --rm -v "$HOME/.m2":/root/.m2 -v "$PWD:/iotdb" maven:3.8.6-openjdk-11 bash
cd iotdb
mvn clean package -pl distribution -am -DskipTests

wget https://downloads.apache.org/iotdb/0.12.6/apache-iotdb-0.12.6-server-bin.zip
wget https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/0.12.6/apache-iotdb-0.12.6-server-bin.zip

https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-cli-bin.zip
https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-confignode-bin.zip
https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-datanode-bin.zip
https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-grafana-plugin-bin.zip
https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-library-udf-bin.zip
https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-source-release.zip

wget https://mirrors.tuna.tsinghua.edu.cn/apache/iotdb/1.0.0/apache-iotdb-1.0.0-all-bin.zip

cd docker
sudo docker build -t yiluxiangbei/iotdb:v1.0.0 .
sudo docker push yiluxiangbei/iotdb:v1.0.0

cd git/docker-iot/iotdb/docker

docker run --name iotdb -p 6667:6667 -v /home/boring/iotdb-0.12.6/data:/iotdb/data -v /home/boring/iotdb-0.12.6/logs:/iotdb/logs -d iotdb:0.12.6 /iotdb/bin/start-server.sh

sudo docker run --name iotdb -it --rm -p 6667:6667 -v "$(pwd)"/data:/iotdb/data -v "$(pwd)"/logs:/iotdb/logs yiluxiangbei/iotdb:v1.0.0 bash

cd git/docker-iot/iotdb/docker
sudo docker run --name iotdb -itd -p 2883:1883 -p 6667:6667 -v "$(pwd)"/data:/iotdb/data -v "$(pwd)"/conf:/iotdb/conf -v "$(pwd)"/logs:/iotdb/logs yiluxiangbei/iotdb:v1.0.0
sudo docker ps -a|grep iotdb
sudo docker exec -it iotdb bash
apt-get install telnet
sudo docker logs -f iotdb
telnet localhost 2883
telnet 172.21.16.11 2883
telnet localhost 6667
telnet 172.21.16.11 6667

tail -f logs/*.log
172.21.16.11
2883

sudo docker exec -it iotdb bash
/iotdb/sbin/start-cli.sh -h 127.0.0.1 -p 6667 -u root -pw root
CREATE DATABASE root.ln
SHOW DATABASES
CREATE TIMESERIES root.ln.wf01.wt01.status WITH DATATYPE=BOOLEAN, ENCODING=PLAIN
CREATE TIMESERIES root.ln.wf01.wt01.temperature WITH DATATYPE=FLOAT, ENCODING=RLE
SHOW TIMESERIES
SHOW TIMESERIES root.ln.wf01.wt01.status
INSERT INTO root.ln.wf01.wt01(timestamp,status) values(100,true);
INSERT INTO root.ln.wf01.wt01(timestamp,status,temperature) values(200,false,20.71)
SELECT status FROM root.ln.wf01.wt01
SELECT * FROM root.ln.wf01.wt01
SET time_zone=+08:00
SELECT * FROM root.ln.wf01.wt01

/iotdb/sbin/start-cli.sh -h 127.0.0.1 -p 6667 -u root -pw root
SHOW DATABASES
SHOW TIMESERIES
SHOW TIMESERIES root.sg.abc
SELECT * FROM root.sg.abc

SHOW TIMESERIES root.sg.mqttx_ed36b238.bar
SELECT * FROM root.sg.mqttx_ed36b238
SELECT bar FROM root.sg.mqttx_ed36b238

docker stop iotdb
docker rm iotdb

docker logs -f iotdb
[dumb-init] /iotdb/bin/start-standalone.sh: No such file or directory

find . -name "README.md"|grep docs|xargs -I {} grep 'docker run' {}

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.21.16.11  netmask 255.255.240.0  broadcast 172.21.31.255
        inet6 fe80::5054:ff:fec9:1f8f  prefixlen 64  scopeid 0x20<link>
        ether 52:54:00:c9:1f:8f  txqueuelen 1000  (Ethernet)
        RX packets 1413152213  bytes 666593816239 (620.8 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 1483768630  bytes 417078226940 (388.4 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

```
docker iotdb
```

```shell
sudo docker run -ti --volume="$(pwd)":/app --rm paddlepaddle/paddle:2.2.0 bash
sudo docker run -ti --volume="$(pwd)":/app --rm python:3.6 bash
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
sudo docker run -it --name fire-smoke-detect-paddle -p 8866:8866 -d yiluxiangbei/fire-smoke-detect-paddle:v1.1
sudo docker run -it --name all-paddle-hub -p 8866:8866 -d yiluxiangbei/all-paddle-hub:v1.1
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
hyperledger/fabric-peer docker run
docker run -p 8000:8000 casbin/casdoor-all-in-one:$LATEST-RELEASE
sudo docker run -p 8091:8000 casbin/casdoor-all-in-one
sudo docker run -ti --volume="$(pwd)":/app --rm openjdk:8-jre bash
sudo docker run -ti -d --name=xxl-job-admin -p 8089:8080 --volume="$(pwd)":/xxl-job yiluxiangbei/xxl-job-admin:v1.0
sudo docker run -it --rm --name my-maven-project -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven maven:3.3-jdk-8 mvn clean install
sudo docker run -it --rm -v "$PWD":/usr/src/mymaven -v "$HOME/.m2":/root/.m2 -v "$PWD/target:/usr/src/mymaven/target" -w /usr/src/mymaven maven mvn clean package
sudo docker run -it --rm -v "$HOME/.m2":/root/.m2 -v "$PWD:/pageplug" maven:3.8.6-openjdk-11 bash
sudo docker run -ti --volume="$(pwd)":/app --rm php:7.4.30-fpm bash
sudo docker run -ti --volume="$(pwd)":/app --rm php:7.3.7-fpm bash
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/php:7.3.7-fpm bash
sudo docker run -i -t -d -p 80:80 --restart=always \
sudo docker run -i -t -d -p 80:80 -p 443:443 \
sudo docker run -i -t -d -p 443:443 --restart=always \
sudo docker run -i -t -d -p 1443:443 --restart=always onlyoffice/documentserver
sudo docker run -i -t -d -p 443:443 \
sudo docker run -i -t -d -p 1443:443 \
sudo docker run -ti --volume="$(pwd)":/DeepSolar --rm python:3.9 bash
sudo docker run -ti --volume="$(pwd)":/DeepSolar --rm python:3.8 bash
sudo docker run -it --name="anaconda" -p 8888:8888 continuumio/anaconda3 /bin/bash
sudo docker run -it --name="anaconda" --volume="$(pwd)":/DeepSolar --rm -p 8888:8888 continuumio/anaconda3 /bin/bash
sudo docker run -it --name="anaconda" --volume="$(pwd)":/DeepSolar --rm -p 8888:8888 continuumio/anaconda3:2019.07 /bin/bash
docker run -it --name="anaconda_2" -p 8888:8888 -v `pwd`:/root continuumio/anaconda3 /bin/bash
docker run -it -d --name [container-name] -p 8088:80 [image-name]
docker run -it -d --name container-name -p p1:p1 -p p2:p2 new-image-name
docker run -it --name="anaconda" -p 8888:8888 continuumio/anaconda3 /bin/bash
sudo docker run -p 8000:8000 casbin/casdoor-all-in-one:$LATEST-RELEASE
sudo docker run -p 8091:8000 casbin/casdoor-all-in-one
sudo docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker -v $(pwd):/app yiluxiangbei/waar bash
sudo docker run -it --rm python:3.7 bash
sudo docker run -it yiluxiangbei/canary-platform:20220424 bash
sudo docker run -it -d --name canary-platform yiluxiangbei/canary-platform:20220424
sudo docker run -it -d -p 8072:8080 --name canary-platform yiluxiangbei/canary-platform:20220424
sudo docker run -it --rm python:3.7 bash
sudo docker run -it yiluxiangbei/ldap-admin:20220424 bash
sudo docker run -it -d --name ldap-admin yiluxiangbei/ldap-admin:20220424
sudo docker run -it -d -p 8072:8080 --name ldap-admin yiluxiangbei/ldap-admin:20220424
sudo docker run -ti --privileged --volume="$(pwd)":/maxkey-web-frontend -v "$(pwd)"/root:/root --rm node:14 bash
sudo docker run -it --rm openjdk:8-jre bash
sudo docker run -p 389:389 -p 636:636 --name my-openldap-container --detach osixia/openldap:1.5.0
sudo docker run -p 389:389 -p 636:636 \
sudo docker run -p 8074:443 \
sudo docker run -ti --volume="$(pwd)":/app --rm paddlepaddle/paddle:2.2.0 bash
sudo docker run -ti --volume="$(pwd)":/app --rm python:3.6 bash
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
sudo docker run -it --name fire-smoke-detect-paddle -p 8096:8866 -d yiluxiangbei/fire-smoke-detect-paddle:v1.1
sudo docker run -it --name all-paddle-hub -p 8096:8866 -d yiluxiangbei/all-paddle-hub:v1.1
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:paddlex1-v1.0 bash
sudo docker run -it --rm yiluxiangbei/all-paddle-hub:paddlex1-v1.0 bash
sudo docker run -it --name all-paddle-hub1 -p 8097:8866 --rm yiluxiangbei/all-paddle-hub:paddlex1-v1.1
sudo docker run -it --name all-paddle-hub1 -p 8097:8866 -d yiluxiangbei/all-paddle-hub:paddlex1-v1.1
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
sudo docker run -it --rm yiluxiangbei/all-paddle-hub:v1.0 bash
sudo docker run -it --name all-paddle-hub -p 8096:8866 --rm yiluxiangbei/all-paddle-hub:v1.1
sudo docker run -it --name all-paddle-hub -p 8096:8866 -d yiluxiangbei/all-paddle-hub:v1.1
sudo docker run -ti --volume="$(pwd)":/app --rm yiluxiangbei/paddlehub:v1.0 bash
sudo docker run -it --name cartoonize -p 8701:8080 -d yiluxiangbei/cartoonize
sudo docker run -it --volume="$(pwd)":/app --rm yiluxiangbei/cartoonize:1.0 bash
sudo docker run -it --volume="$(pwd)":/app --rm yiluxiangbei/cartoonize-pip:1.0 bash
sudo docker run -it --volume="$(pwd)":/app --rm python:3.7-slim bash
sudo docker run -ti --volume="$(pwd)":/DeepSolarV1 --rm python:2.7 bash
# sudo docker run -ti --volume="$(pwd)":/DeepSolarV1 --rm python:3.7 bash
sudo docker run -it --volume="$(pwd)":/app1 --rm yiluxiangbei/magnus-pytorchserver:2.0.4 bash
docker run yiluxiangbei/magnus-kratos:v2 bash
sudo docker run -ti --privileged --volume="$(pwd)":/webssh -v $(pwd)/root:/root --rm node:16 bash
# sudo docker run -ti --privileged --volume="$(pwd)":/webssh -v $(pwd)/root:/root --rm node:14 bash
sudo docker run -itd --name op1 --network host myide/opencmdb:v1
sudo docker run -itd --name op1 myide/opencmdb:v1
sudo docker run -it --rm --name op1 myide/opencmdb:v1
sudo docker run --name tipdm-db -d --privileged=true -p 5431:5432 yiluxiangbei/centos7-tipdm-db:20201112
sudo docker run --name tipdm -d --privileged=true yiluxiangbei/centos7-tipdm:20201112 /usr/sbin/init
sudo docker run -d --name tipdm --privileged=true centos:7.8.2003 /usr/sbin/init
sudo docker run -d \
sudo docker run -d \
sudo docker run --rm -it --privileged -v $(pwd):/app python:3.6 /bin/bash
sudo docker run --rm -it -p 8011:8000 -v $(pwd):/app --privileged yiluxiangbei/heart-rate-detection:v1 bash
sudo docker run --rm -it -p 8011:8000 --privileged yiluxiangbei/heart-rate-detection:v1 bash
sudo docker run --rm -it -p 8011:8000 yiluxiangbei/heart-rate-detection:v2 --entrypoint /bin/bash
sudo docker run --rm -it -p 8011:8000 yiluxiangbei/heart-rate-detection:v2
sudo docker run --name heart-rate -itd -p 8011:8000 yiluxiangbei/heart-rate-detection:v2
sudo docker run --name heart-rate -itd -v $(pwd):/app -p 8011:8000 yiluxiangbei/heart-rate-detection:v2
docker run --rm -it registry.cn-hangzhou.aliyuncs.com/rotanova/rotanava-boot-module-gateway:1.0.0 sh
```

```
//显示版本
show version
 
//新建存储组
SET STORAGE GROUP TO root.test
 
//显示已经创建的存储组
SHOW STORAGE GROUP
 
//删除存储组
delete storage group root.test
 
//创建时间序列
create timeseries root.test.wf01.wt01.s0(elec_meter) with datatype=FLOAT, encoding=RLE, compression=SNAPPY,MAX_POINT_NUMBER=3 tags(tag1=t1, tag2=t2) attributes(attr1=a1, attr2=a2)
 
//删除时间序列
delete timeseries root.test.wf01.wt01.s0
delete timeseries root.test.wf01.wt01.*
 
//更改时间序列
alter timeseries root.test.wf01.wt01.s0 DROP tag1, tag2
alter timeseries root.test.wf01.wt01.s0 SET tag1=t3, attr1=a3
alter timeseries root.test.wf01.wt01.s0 ADD TAGS tag3=t3, tag4=t4
 
//显示存储组下时间序列
show timeseries root.test
 
//模糊查询设备下的序列
show timeseries root.test.*.wt01
 
//加条件的查询时间序列
show timeseries root.test where tag1=t1
 
//统计时间序列
count timeseries root.test.*.wt01
 
//根据层级分组，统计出每组下的时间序列个数
count timeseries root.test group by level=3
 
//统计存储组对应的层级下有几个节点
count nodes root.test level=3
 
//统计存储组下几个设备
show devices root.test
 
//显示某一层下的子路径
show child paths root.test.wf01

# CRUD
//插入记录
insert into root.test.wf01.wt02(timestamp,elec_meter) values (1,23.896)
 
//更新记录
insert同时间戳覆盖即可
 
//删除记录
delete from root.test.wf01.wt02.s0,root.test.wf01.wt02.s1 where time < now()
 
 
//查询记录
select temperature from root.test.wf01.wt01 where time >0 and time <5
select count(s0) from root.test.wf01.wt02 WHERE root.test.wf01.wt02.s0 < 25
select max_value(s0) from root.test.wf01.wt02 where time < now() - 5h
select min_time(s0) from root.test.wf01.wt02
 
//排序 0.11
select * from root.test.wf01.wt01 where time > 1 order by time desc limit 10;
 
//分组应用
//统计这段时间区间内（前包，后不包）每隔5分钟统计s0个数
select count(s0) from root.test.wf01.wt02 group by ([1,259222222),5m)
//统计这段时间内每隔5小时，前2个小时有数据个数
select count(s0) from root.test.wf01.wt02 group by ([86400000,259222222),2h,5h)
 
 
 
//分页
//时间戳升序，首条下标为0，取3条
select * from root.test.wf01.wt02 limit 3 offset 0
 
 
//筛选最后一条记录
select last temperature from root.test.wf01.wt02
```