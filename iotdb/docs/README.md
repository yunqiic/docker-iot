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

sudo docker run --name iotdb -it -p 6667:6667 --privileged=true -v "$(pwd)"/data:/iotdb/data -v "$(pwd)"/logs:/iotdb/logs -d yiluxiangbei/iotdb:v1.0.0 /iotdb/sbin/start-standalone.sh
sudo docker ps -a|grep iotdb
sudo docker exec -it iotdb bash

/iotdb/sbin/start-cli.sh -h 127.0.0.1 -p 6667 -u root -pw root

docker stop iotdb
docker rm iotdb

docker logs -f iotdb
[dumb-init] /iotdb/bin/start-standalone.sh: No such file or directory

find . -name "README.md"|grep docs|xargs -I {} grep 'docker run' {}
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