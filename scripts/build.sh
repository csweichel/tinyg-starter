#!/bin/bash

set -eoux pipefail

# Build the project for the target
# print usage if no arguments are given
if [ $# -eq 0 ]; then
    echo "Usage: $0 <target_name>"
    exit 1
fi

# cd to the root
cd "$(dirname "$0")/.."

# get the target name
target_name=$1

# build the target
tinygo build -target=pico -opt=1 -stack-size=8kb -size=short -monitor -o "$target_name.uf2" "cmd/$target_name/main.go"
