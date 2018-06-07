#!/bin/sh

set -o nounset
set -o errexit

# helper script to update build date and build number
# typically used in a pre-commit hook.

CURRENT_BUILD_NUM=$(grep -o "internal_BUILD_NUMBER\\s*=\\s*[0-9]\\+" version.go | grep -o "[0-9]\\+")
NEW_NUM=$((CURRENT_BUILD_NUM+1))
echo "Bumping build number, old=${CURRENT_BUILD_NUM} new=$NEW_NUM"

sed -i.bak "s/internal_BUILD_NUMBER[[:space:]]*=[[:space:]]*[0-9]*/internal_BUILD_NUMBER\\ =\\ ${NEW_NUM}/g" version.go


NEW_TS=$(date +%s)
echo "Bumping build timestamp new=$NEW_TS"

sed -i.bak "s/internal_BUILD_TIMESTAMP[[:space:]]*=[[:space:]]*[0-9][0-9]*/internal_BUILD_TIMESTAMP\\ =\\ ${NEW_TS}/g" version.go

rm -f version.go.bak
