```shell
Using Docker：The path to the dockerfile is https://github.com/apache/iotdb/tree/master/docker/src/main
mvn clean package -pl distribution -am -DskipTests

find . -name "README.md"|grep docs|xargs -I {} grep 'docker run' {}
```