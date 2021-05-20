#!/bin/env bash
set -e
set +x
set -o pipefail

function userType()
{
    if [ "$EUID" = '0' ]; then
        echo "root"
    else
        net session > /dev/null 2>&1
        if [ $? -eq 0 ]; then
            echo "admin"
        else
            echo "user";
        fi
    fi
}

if [ "$(userType)" = "user" ]; then
    echo "Please run as administrator or root"
    exit 1
fi

PB_REL='https://github.com/protocolbuffers/protobuf/releases'
PB_VERSION='3.17.0'
# win32 | win64 | linux-x86_64 | linux-x86_32 | osx-x86_64 | linux-aarch_64 | linux-ppcle_64 | linux-s390_64
OS_ARCH=win64
INSTALL_PATH=/usr/local/bin

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ARCHIVE_FILE_NAME="protoc-$PB_VERSION-$OS_ARCH.zip"
curl -LO "$PB_REL/download/v$PB_VERSION/$ARCHIVE_FILE_NAME"

TEMP_DIR=`mktemp -d -p "$CURRENT_DIR"`

function cleanup {      
  rm -rf "$TEMP_DIR"
  echo "Deleted temp archive directory $TEMP_DIR"
  rm "$ARCHIVE_FILE_NAME"
  echo "Deleted temp download file $ARCHIVE_FILE_NAME"
}
trap cleanup EXIT
mkdir -p "$INSTALL_PATH"
unzip "$ARCHIVE_FILE_NAME" -d "$TEMP_DIR"
for f in "$TEMP_DIR/bin/*"; do
  FILE_NAME=$(basename $f)
  mv "$TEMP_DIR/bin/$FILE_NAME" "$INSTALL_PATH/$FILE_NAME"
done

GREEN='\033[0;32m'
NC='\033[0m'
printf "${GREEN}Installed version:\n$(protoc --version)\n${NC}"

echo "Installing Go plugins"

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
