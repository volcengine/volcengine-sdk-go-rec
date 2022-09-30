#!/bin/bash
set -e
export https_proxy=http://10.20.47.147:3128 \
       http_proxy=http://10.20.47.147:3128 \
       no_proxy="*.byted.org"

GitlabFolder=$(pwd)
# 列出与master的差异
Commits=$(git rev-list origin/master..HEAD)

# 拉取git仓库
cd ..
ls -alh
mkdir .tmp_git
cd .tmp_git
git clone https://github.com/volcengine/volcengine-sdk-go-rec.git
cd volcengine-sdk-go-rec

# Operates in a new branch
NEW_BRANCH="sync-$(date '+%Y%m%d_%H%M')"
git checkout -b ${NEW_BRANCH}
FileToDelete=$(ls | grep -v '.git')
rm -rf ${FileToDelete}
ls -alh
GithubFolder=$(pwd)
awk -v cmd="cp -ri "${GitlabFolder}"/* "${GithubFolder}"/" 'BEGIN {print "n" |cmd;}'

# Commits the changes
if [[ -n $(git status -s) ]]; then
  git remote set-url origin "https://ganlin.coder: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@github.com/volcengine/volcengine-sdk-go-rec.git"
  git remote -v
#  git config --global user.name "ganlin"
#  git config --global user.email ganlin.coder@bytedance.com
  git add -A .
  git commit --amend --reset-author -m "[SDK] Sync to github." -m "$Commits"
  git push origin ${NEW_BRANCH}
  body=$(awk '{printf "%s\\n", $0}' <<< ${commits}) # Escape new lines
  curl \
    -X POST \
    -H "Accept: application/vnd.github.v3+json" \
    -u "ganlin.coder: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" \
    https://api.github.com/repos/volcengine/volcengine-sdk-go-rec/pulls \
    -d "{\"head\": \"$NEW_BRANCH\", \"base\": \"sync-20220930_0832\", \"maintainer_can_modify\": true, \"title\": \"[SDK] Sync to github.\", \"body\": \"$body\"}"
fi

#git checkout "$SOURCE_BRANCH"
#mkdir .tmp_git
#cd .tmp_git
#git clone https://github.com/bytedance/fedlearner.git
#cd fedlearner
## Operates in a new branch
#git checkout -b "$NEW_BRANCH"
#rm -rf web_console_v2
#cp -R "$ROOT_FOLDER/web_console_v2" web_console_v2/
#
#if [[ -n $(git status -s) ]]; then
#  git remote set-url origin "https://$USERNAME:$ACCESS_TOKEN@github.com/bytedance/fedlearner.git"
#  # Commits the changes
#  git add -A .
#  git commit -m "[WebConsole] Sync to github." -m "$commits"
#  git push origin "$NEW_BRANCH"
#  # Creates pull request
#  body=$(awk '{printf "%s\\n", $0}' <<< "$commits") # Escape new lines
#  curl \
#    -X POST \
#    -H "Accept: application/vnd.github.v3+json" \
#    -u "$USERNAME:$ACCESS_TOKEN" \
#    https://api.github.com/repos/bytedance/fedlearner/pulls \
#    -d "{\"head\": \"$NEW_BRANCH\", \"base\": \"master\", \"maintainer_can_modify\": true, \"title\": \"[WebConsole] Sync to github.\", \"body\": \"$body\"}"
#fi