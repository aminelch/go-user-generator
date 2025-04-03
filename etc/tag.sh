#!/usr/bin/env bash

###
#  Shell script for automation creating tag
#  Copyright (C) 2022 Amine LOUHICHI, Inc. All rights reserved.
#
#  This program is a free software; you can redistribute it
#  and/or modify it under the terms of the GNU AFFERO GENERAL PUBLIC LICENSE
###

set -e

if ! command -v date >/dev/null 2>&1; then
  echo "It looks like Date is not installed on your system. Please install it first by running 'apt-get install curl'
  and try again to continue the installation."
  exit 1
fi

DATE="$(date +'%Y-%m-%d %H:%M:%S')"
VERSION=
TARGET_FILE=$(realpath "$PWD/VERSION")

[[ -z  $1 ]] && VERSION="1.0.0" || VERSION=$1

echo version: "$VERSION $DATE" > "$TARGET_FILE"


