#!/usr/bin/env bash
# This script is used to bump the version of the sub module.
#
# github.com/alswl/makefile-go
#
# Author: alswl
# Version: 0.2.2
#
# Usage: hack/bump-sub-module.sh <sub> <next> <dryrun>
# stage: final, alpha, beta, candidate
# scope: major, minor, patch, final
#
# 发版之前打 tag，会自动写入 VERSION 文件，并生成 tag
# 写入 VERSION 文件会强制带上 `-dev` 结尾，但是 push tag 则没有 -dev
# 因此，master 会触发包构建产生的镜像都有 -dev 结尾，而 tag push 产生的是明确的没有 -dev 的版本（尽管他们的 hash 一样）
# 即以 tag 作为明确版本发布信号

# cd root of the repo
pushd "$(dirname "$0")/.." > /dev/null

set -e

sub=$1
next=$2
bump_dry_run=$3

if [ -z "$sub" ]; then
  echo "sub mod is required"
  exit 1
fi
if [ ! -d "$sub" ]; then
  echo "sub mod $sub does not exist"
  exit 1
fi
if [ -z "$next" ]; then
  echo "next is required"
  exit 1
fi
if [ -z "$bump_dry_run" ]; then
  echo "bump dryrun is required"
  exit 1
fi

echo "next version: $next"

# dry run
if [ "$bump_dry_run" = "true" ]; then
  echo "dryrun: true"
  exit 0
fi

# 1. bump VERSION
echo "dryrun: false"
echo "${next}" > "${sub}"/VERSION
git add "${sub}"/VERSION
git commit -m "chore: Bump version to $next"

# 2. git tag
git tag "${sub}/${next}"

# 3. modified VERSION in stage
# VERSION file always has the -dev suffix
echo "${next}-dev" > "${sub}"/VERSION
git add "${sub}"/VERSION

# 4. prompt
echo "# Your should update your local version file to dev"
echo "# Run: "
echo "git push origin ${sub}/${next}"
echo "vim ${sub}/VERSION"
echo 'git commit -m "chore: new dev version"'
echo "git push"

popd > /dev/null
