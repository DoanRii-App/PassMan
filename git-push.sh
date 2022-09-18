#! /bin/bash
if [ -z $1 ]; then
        echo "Empty commit message"
        exit 0
else
        git add . && git commit -m "$1" && git push https://$2@$3
fi