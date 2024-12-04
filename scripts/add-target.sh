#!/bin/bash

set -eoux pipefail

# Add a new target to the project copied from the template

# print usage if no arguments are given
if [ $# -eq 0 ]; then
    echo "Usage: $0 <target_name>"
    exit 1
fi

# cd to the root
cd "$(dirname "$0")/.."

# get the target name
target_name=$1

# copy the template target to the new target
cp -r cmd/template "cmd/$target_name"
