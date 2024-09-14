#!/bin/sh

echo ""
echo " Choose a command run in \"$1\":"
sed -n 's/^##//p' $2 | column -t -s ':' |  sed -e 's/^/  /'
