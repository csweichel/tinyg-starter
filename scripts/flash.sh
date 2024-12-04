#!/bin/bash

set -eoux pipefail

# print usage if not enough arguments are given
if [ $# -lt 2 ]; then
    echo "Usage: $0 <environmentID> <target_name>"
    exit 1
fi

# get target name
envid=$1
target_name=$1

gitpod env ssh-config
gitpod env ssh "$envid" ./scripts/build.sh "$target_name"

scp "$envid".gitpod.environment:$target_name.uf2 /Volumes/RPI-RP2
