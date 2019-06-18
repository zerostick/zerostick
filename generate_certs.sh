#!/bin/bash

mkdir -p zerostick_web/certs
cd zerostick_web/certs
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
