#!/bin/bash

BASEPATH=$(realpath $(dirname $0)/..)

mkdir -p ${BASEPATH}/zerostick_web/certs
cd ${BASEPATH}/zerostick_web/certs
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
