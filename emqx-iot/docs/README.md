```
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
```