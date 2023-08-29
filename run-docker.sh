docker-compose up -d postgres redis etcd
sleep 5s
docker-compose up -d simple-douyin
sleep 5s
docker-compose up -d