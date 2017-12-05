#!/bin/sh -eux

targets=`find . -type f \( -name '*.go' -and -not -iwholename '*vendor*' -and -not -iwholename '*misc*' \)`
packages=`go list ./... | grep -v misc`

# Apply tools
export PATH=$(pwd)/build-cmd:$PATH
which goimports golint
goimports -w $targets
go tool vet $targets
golint $packages

go test $packages $@
