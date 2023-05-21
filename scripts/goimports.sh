Dir=$(pwd)
mod="$Dir/go.mod"

APP_NAME=$(cat $mod | grep "module" | awk '{gsub(/^\s+|\s+$/," ");print $2}')

goimports-reviser -rm-unused -format -local $APP_NAME ./...