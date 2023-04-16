# 若制作镜像,保证项目分支所有更改已提交到远程服务器,并打上新的TAG(git版本号和镜像版本号一致)

if [ -n "$(git status -s)" ];then
    echo "[\e[31m中断\e[0m] 本地已更改未提交"
    exit
fi

if [ -n "$(git cherry -v)" ];then
    echo "[\e[31m中断\e[0m] 本地新提交未push"
    exit
fi

git fetch --tags
Dir=$(pwd)
CONFIG="$Dir/configs/configs.yaml"
GIT_BRACH=$(git symbolic-ref --short -q HEAD)
GIT_TAG=$(git describe --tags --abbrev=0)
SUB_GIT_TAG=$(echo ${GIT_TAG#*v} | awk -F "." '{ printf("%s.%s.%s",$1,$2,$3+1)}')
GIT_COMMIT=$(git rev-parse --short HEAD)

APP_NAME=$(cat $CONFIG | grep "appName" | awk -F ":" '{print}' | awk '{gsub(/^\s+|\s+$/," ");print $2}')
#APP_VERSION=$(cat $CONFIG | grep "version" | awk -F ":" '{print}' | awk '{gsub(/^\s+|\s+$/," ");print $2}')
#sed -i 's/version: '$APP_VERSION'/version: '$SUB_GIT_TAG'/' $CONFIG

IMAGE="xiaohubai/$APP_NAME:$SUB_GIT_TAG-$GIT_BRACH-$GIT_COMMIT"
docker build . -t $IMAGE
docker push $IMAGE