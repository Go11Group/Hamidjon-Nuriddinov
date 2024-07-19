#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Foydalanish: $0 <hozirgi_katalog>"
  exit 1
fi

CURRENT_DIR=$1

if [ ! -d "$CURRENT_DIR" ]; then
  echo "Xatolik: $CURRENT_DIR katalogi mavjud emas."
  exit 1
fi

rm -rf ${CURRENT_DIR}/genproto
mkdir -p ${CURRENT_DIR}/genproto

PROTO_FILES=$(find ${CURRENT_DIR}/protos -name "*.proto")
for PROTO_FILE in $PROTO_FILES; do
  protoc -I=$(dirname ${PROTO_FILE}) -I=${CURRENT_DIR}/protos -I /usr/local/go \
    --go_out=paths=source_relative:${CURRENT_DIR}/genproto \
    --go-grpc_out=paths=source_relative:${CURRENT_DIR}/genproto \
    ${PROTO_FILE}

  if [ $? -ne 0 ]; then
    echo "Xatolik: ${PROTO_FILE} uchun kod generatsiya qilishda muvaffaqiyatsiz bo'ldi"
    exit 1
  fi
done

echo "Protobuf generatsiyasi muvaffaqiyatli yakunlandi."
