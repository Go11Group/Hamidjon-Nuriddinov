#!/bin/zsh
CURRENT_DIR=$1

# Log yo'llarni ko'rsatish
echo "Running script with current directory: ${CURRENT_DIR}"

# protos ichidagi barcha fayllarni o'chirish, lekin direktoriyani saqlab qolish
if [ -d "${CURRENT_DIR}/protos" ]; then
  find ${CURRENT_DIR}/protos -type f -delete
else
  mkdir -p ${CURRENT_DIR}/protos
fi

# transport.proto va weather.proto fayllarni kompilyatsiya qilish
for proto_file in transport weather; do
  echo "Processing file: ${proto_file}.proto"
  protoc -I=${CURRENT_DIR}/proto -I=/usr/local/go --go_out=${CURRENT_DIR}/protos/${proto_file} \
  --go-grpc_out=${CURRENT_DIR}/protos/${proto_file} ${CURRENT_DIR}/proto/${proto_file}.proto
  
  # Agar xatolik yuz bersa logni ko'rsatish
  if [ $? -ne 0 ]; then
    echo "Error processing ${proto_file}.proto"
  fi
done

# Yaratilgan fayllarni ko'rsatish
echo "Generated files in ${CURRENT_DIR}/protos:"
find ${CURRENT_DIR}/protos -type f -name "*.pb.go"
