#!/bin/bash

set -e

ROOT_FOLDER=$(git rev-parse --show-toplevel)
cd "$ROOT_FOLDER"

NEW_BRANCH="sync-$(date '+%Y%m%d_%H%M')"


# Proxies
export https_proxy=http://10.20.47.147:3128 \
       http_proxy=http://10.20.47.147:3128 \
       no_proxy="*.byted.org"

commits=$(git rev-list --format=%B --ancestry-path "$TARGET_COMMIT".."$SOURCE_COMMIT")

git checkout "$SOURCE_BRANCH"
mkdir .tmp_git
cd .tmp_git
git clone https://github.com/volcengine/volcengine-sdk-go-rec.git
cd volcengine-sdk-go-rec
# Operates in a new branch
git checkout -b "$NEW_BRANCH"

if [[ -n $(git status -s) ]]; then
  git remote set-url origin "https://$USERNAME:$ACCESS_TOKEN@github.com/volcengine/volcengine-sdk-go-rec.git"
  # Commits the changes
  git add -A .
  git commit -m "[WebConsole] Sync to github." -m "$commits"
  git push origin "$NEW_BRANCH"
  # Creates pull request
  body=$(awk '{printf "%s\\n", $0}' <<< "$commits") # Escape new lines
  curl \
    -X POST \
    -H "Accept: application/vnd.github.v3+json" \
    -u "$USERNAME:$ACCESS_TOKEN" \
    https://api.github.com/repos/volcengine/volcengine-sdk-go-rec/pulls \
    -d "{\"head\": \"$NEW_BRANCH\", \"base\": \"master\", \"maintainer_can_modify\": true, \"title\": \"[WebConsole] Sync to github.\", \"body\": \"$body\"}"
fi
