# ！切记：部署容器的外部数据卷，不要删除

# 前期准备工作(Ubuntu 22.04 LTS)
``` go
//root设置密码
sudo passwd root
//新增用户
adduser xiaohu
//删除用户
deluser ubuntu
//删除用户组
groupdel ubuntu
//验证xiaohu用户是否有sudo权限
sudo -l -U xiaohu
//用户设置sudo无密码执行
chmod -v u+w /etc/sudoers
vim /etc/sudoers
xiaohu  ALL=(ALL:ALL) NOPASSWD: ALL
chmod =r /etc/sudoers
//
```
## 安装docker docker-compose
``` go
sudo apt update
sudo apt upgrade
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
sudo apt autoremove -y && sudo apt autoclean -y
sudo mkdir -p /etc/docker
sudo vim /etc/docker/daemon.json

{
  "registry-mirrors": ["https://mirror.ccs.tencentyun.com"]
}

sudo systemctl daemon-reload
sudo systemctl restart docker
sudo systemctl enable docker

```
## 挂载并编排docker-compose.yml
``` go
//打包本地deploy到远端服务器
tar -zcvf deploy.tar.gz deploy
//解压deploy.tar.gz
tar -zxvf deploy.tar.gz

cd deploy
sudo chmod 777 volumes/redis/data
sudo chmod 777 volumes/prometheus/data
sudo chmod 777 volumes/mysql
sudo chmod 777 volumes/kafka
sudo chmod 777 volumes/grafana/data
sudo chmod 777 volumes/elasticsearch/data
sudo chmod 777 volumes/consul

sudo chmod 777 /usr/local
mv volumes /usr/local/
sudo docker compose -f ./docker-compose.yml down
sudo docker compose -f ./docker-compose.yml up -d --force-recreate


```
