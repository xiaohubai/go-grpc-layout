if [ -d "/usr/local/volumes" ];then
	echo "文件夹: volumes已存在"
else
	cp -r volumes /usr/local
	chmod 777 /usr/local/volumes
	chmod 777 /usr/local/volumes/consul
	chmod 777 /usr/local/volumes/elasticsearch/data
	chmod 777 /usr/local/volumes/elasticsearch/plugins/ik
	chmod 777 /usr/local/volumes/grafana/data
	chmod 777 /usr/local/volumes/kafka
	chmod 777 /usr/local/volumes/mysql
	chmod 777 /usr/local/volumes/prometheus/data
	chmod 777 /usr/local/volumes/redis/data
	echo "success"
fi