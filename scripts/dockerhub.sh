if [ -n "$(git status -s)" ];then
    echo "[\e[31m中断\e[0m] 本地已更改未提交"
    exit
fi



git fetch --tags

CONFIG="../configs/configs.yaml"

GIT_BRACH=$(git symbolic-ref --short -q HEAD)
GIT_TAG=$(git describe --tags --always)
SUB_GIT_TAG=${GIT_TAG#*v}
GIT_COMMIT=$(git rev-parse --short HEAD)

APP_VERSION=$(cat $CONFIG | grep "version" | awk -F ":" '{print}' | awk '{gsub(/^\s+|\s+$/," ");print $2}')
APP_NAME=$(cat $CONFIG | grep "appName" | awk -F ":" '{print}' | awk '{gsub(/^\s+|\s+$/," ");print $2}')

sed -i 's/version: '$APP_VERSION'/version: '$SUB_GIT_TAG'/' $CONFIG

#docker build . -t "xiaohubai/""$APP_NAME"":"${SUB_GIT_TAG}"-"$GIT_BRACH"-"$GIT_COMMIT

echo "xiaohubai/$APP_NAME:$SUB_GIT_TAG-$GIT_BRACH-$GIT_COMMIT"


