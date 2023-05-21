Dir=$(pwd)
CONFIG="$Dir/configs/conf/conf.yaml"

APP_NAME=$(cat $CONFIG | grep "appName" | awk -F ":" '{print}' | awk '{gsub(/^\s+|\s+$/," ");print $2}')

goimports -local $APP_NAME -w $Dir