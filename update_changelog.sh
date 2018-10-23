#!/usr/bin/env bash

set -e

usage() {
  echo "$0 <tag> <repo>" >&2;
}

if [ "$1" = "-h" -o "$1" = "--help" ]; then
  usage
  exit 1;
fi

if [ -z "$2" ]
then
  REPO=$(git ls-remote --get-url origin | \
  sed -u 's/git@//g; s/https:\/\///g; s/github.com\///g; s/\.git//g')
else
  REPO=$2
fi

NEW_TAG=$1
CURRENT_DATE=$(date +"%Y-%m-%d")

LAST_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))
LAST_DATE=$(git log -1 --format=%ai $LAST_TAG)

CHANGES=$(curl -s "https://api.github.com/repos/${REPO}/pulls?state=closed" | \
jq --arg l "$LAST_DATE" -r '.[] | select((.merged_at != null) and (.closed_at > $l)) | "- [Pull #\(.number)](\(.html_url)): \(.title)"')

sed -i "4i ## [$NEW_TAG] - $CURRENT_DATE\n### Added\n${CHANGES//$'\n'/\\$'\n'}\n" CHANGELOG.md
