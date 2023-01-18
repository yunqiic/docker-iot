```shell
Using Docker：The path to the dockerfile is https://github.com/apache/iotdb/tree/master/docker/src/main
mvn clean package -pl distribution -am -DskipTests

maven:3.8.6-openjdk-11
yiluxiangbei/maven:3.8.6-openjdk-11

find . -name "README.md"|grep docs|xargs -I {} grep 'docker run' {}
```