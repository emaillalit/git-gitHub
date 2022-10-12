#!/bin/bash

bucket_name="{{ env }}-weblogs.{{ aws_region }}.{{ aws_account_id }}"
log_prefix="container-logs"

year=$(date -u +%Y)
month=$(date -u +%m)
yesterday=$(date  -d "1 day ago" +"%d")

services=("app" "notify" "plm" "bliss")


for service in "${services[@]}"
  do
    echo "Start to pull logs for service $service"
    mkdir -p ~/logs/"$service/$yesterday"

    pushd ~/logs/"$service"
    rm -rf *

    aws s3 sync s3://${bucket_name}/"${log_prefix}"/"${service}"/server/"${year}"/"${month}"/${yesterday} ${yesterday}

    files=$(find . -maxdepth 5 -type f)
    for file in ${files}
      do
        cat $file | jq -r .log > "${file}.log"
        rm -rf $file
      done

    popd
  done

aws s3 sync s3://legs.{{ aws_region }}.{{ aws_account_id }} ~/legs/deploy --exclude "legs"
