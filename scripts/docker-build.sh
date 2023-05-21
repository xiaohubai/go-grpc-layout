# 若制作镜像,保证项目分支所有更改已提交到远程服务器,并打上新的TAG(git版本号和镜像版本号一致)

if [ -n "$(git status -s)" ];then
    echo "[\e[31m中断\e[0m] 本地已更改未提交"
    exit
fi

if [ -n "$(git cherry -v)" ];then
    echo "[\e[31m中断\e[0m] 本地新提交未push"
    exit
fi

Dir=$(pwd)
CONFIG="$Dir/configs/conf/conf.yaml"
GIT_BRACH=$(git symbolic-ref --short -q HEAD)
GIT_COMMIT=$(git rev-parse --short HEAD)
GIT_TAG=$(git describe --tags --abbrev=0)
New_TAG=$(echo ${GIT_TAG#*v} | awk -F "." '{ printf("%s.%s.%s",$1,$2,$3+1)}')
New_GIT_TAG="v"$New_TAG
echo $GIT_TAG $New_TAG $New_GIT_TAG
git tag -a $New_GIT_TAG -m $New_GIT_TAG
git push origin --tags

APP_NAME=$(cat $CONFIG | grep "appName" | awk -F ":" '{print}' | awk '{gsub(/^\s+|\s+$/," ");print $2}')

Image="xiaohubai/$APP_NAME:$New_TAG-$GIT_BRACH-$GIT_COMMIT"
docker build . -t $Image
docker push $Image