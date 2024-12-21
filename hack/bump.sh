#!/usr/bin/env bash
# This script is used to bump the version of the operator. It uses semtag to bump the version.
# github.com/alswl/makefile-go
#
# Author: alswl
# Version: v0.3.1
#
# Usage: hack/bump.sh [--stage <stage>] [--scope <scope>] [--dry-run <dry-run>] [--post-actions-script <post-actions.sh>]
# stage: final, alpha, beta, candidate
# scope: major, minor, patch, final
#
# 发版之前打 tag，会自动写入 VERSION 文件，并生成 tag
# 写入 VERSION 文件会先使用发布版本 push tag
# 推送 tag 之后会使用下一个版本 + `-dev` 作为新版本

# cd root of the repo
pushd "$(dirname "$0")/.." > /dev/null
set -e

bump_stage="final"
bump_scope="major"
bump_dry_run="true"
post_actions_script=""

function runPostActions {
  if [ -n "$post_actions_script" ]; then
    if [ ! -f "$post_actions_script" ]; then
      echo "post actions script not found: $post_actions_script"
    fi
    $post_actions_script
  fi
}

# shift params
while [ $# -gt 0 ]; do
  case "$1" in
    --stage)
      bump_stage=$2
      shift
      ;;
    --scope)
      bump_scope=$2
      shift
      ;;
    --dry-run)
      bump_dry_run=$2
      shift
      ;;
    --post-actions-script)
      post_actions_script=$2
      shift
      ;;
    *)
      echo "unknown option: $1"
      exit 1
      ;;
  esac
  shift
done

next=$(semtag "$bump_stage" -s "$bump_scope" -f -o)
echo "next version: $next"

# dry run
if [ "$bump_dry_run" = "true" ]; then
  echo "dryrun: true"
  exit 0
fi

# bump and tag
echo "dryrun: false"
# VERSION in file always has the -dev suffix
echo "${next}" > VERSION
runPostActions
git add .
git commit -m "chore: Bump version to $next"
# git tag did not contains dev suffix
semtag "$bump_stage" -s "$bump_scope"

# next version
echo "${next}-dev" > VERSION
runPostActions
git add .
git commit -m "chore: prepare next version $next-dev"

popd > /dev/null

