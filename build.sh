#!/usr/bin/env sh

GOROOT=/opt/go
OUTPUT=/opt/bin

#arch=$(uname -m)


if [ "$#" -gt 0 ]; then
    OUTPUT=$1
fi

cd debug
sudo ${GOROOT}/bin/go build -o ${OUTPUT}/buildTestCustomErrors .
