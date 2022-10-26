#!/bin/bash
set -e
export https_proxy=http://${PROXY_HOST}:${PROXY_PORT} \
       http_proxy=http://${PROXY_HOST}:${PROXY_PORT} \

Repo=$(echo $CI_REPO_NAME|awk '{split($1,arr,"/"); print arr[2]}')
git branch -v
git remote add origin-git  https://${GIT_NAME}:${GIT_TOKEN}@github.com/volcengine/${Repo}.git
git remote -v
git tag -l
git checkout -b master
git push origin-git master
git push origin-git --tags

echo "Sync Success!"