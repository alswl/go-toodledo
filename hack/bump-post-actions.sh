#!/usr/bin/env bash

# cd root of the repo
pushd "$(dirname "$0")/.." > /dev/null

set -e
cat package.json | jq -M --raw-output ".version=\"$(cat VERSION | head -n 1 | sed 's/^v//g')\"" | sponge package.json
popd > /dev/null
