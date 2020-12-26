#!/bin/bash

realpath() {
    [[ $1 = /* ]] && echo "$1" || echo "$PWD/${1#./}"
}

BASEPATH=$(realpath $(dirname $0)/..)

mkdir -p ${BASEPATH}/zerostick_web/certs
cd ${BASEPATH}/zerostick_web/certs
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
