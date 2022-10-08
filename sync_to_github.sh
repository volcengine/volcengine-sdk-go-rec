#!/bin/bash
#TODO 环境变量的传递
set -e
export https_proxy=http://10.20.47.147:3128 \
       http_proxy=http://10.20.47.147:3128 \
       no_proxy="*.byted.org"

GitlabFolder=$(pwd)
# 列出与master的差异.
Commits=$(git rev-list --pretty=format:'[%s]' origin/master..HEAD)

# 拉取github仓库
Repo="volcengine-sdk-go-rec"
cd ..
mkdir .tmp_git
cd .tmp_git
git clone https://github.com/volcengine/${Repo}.git
cd ${Repo}

# 创建新的分支
NEW_BRANCH="sync-$(date '+%Y%m%d_%H%M')"
git checkout -b ${NEW_BRANCH}

# 复制gitlab的修改到github的目录下
rm -rf $(ls | grep -v '.git')
GithubFolder=$(pwd)
awk -v cmd="cp -ri "${GitlabFolder}"/* "${GithubFolder}"/" 'BEGIN {print "n" |cmd;}'
rm -rf .codebase
rm -rf sync_to_github.sh

# 提交修改
if [[ -n $(git status -s) ]]; then
  # 如果有修改才提交
  git remote set-url origin https://ganlin.coder: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@github.com/volcengine/${Repo}.git
  git remote -v
  git add -A .
  git commit --amend --reset-author -m "[SDK] Sync to github." -m "$Commits"
  git push origin ${NEW_BRANCH}
  body=$(awk '{printf "%s\\n", $0}' <<< ${commits}) # Escape new lines
  # 用github api创建pr
  curl \
    -X POST \
    -H "Accept: application/vnd.github.v3+json" \
    -u "ganlin.coder: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" \
    https://api.github.com/repos/volcengine/${Repo}/pulls \
    -d "{\"head\": \"$NEW_BRANCH\", \"base\": \"master\", \"maintainer_can_modify\": true, \"title\": \"[SDK] Sync to github.\", \"body\": \"$body\"}"
fi