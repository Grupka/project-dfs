#!/bin/bash

go build -o StorageServer storage_server/storage_server.go
go build -o NamingServer naming_server/main/main.go

mkdir -p run/naming_server
mkdir -p run/storage_server_1
mkdir -p run/storage_server_2

export NAMING_SERVER_ADDRESS="localhost"
export NAMING_SERVER_PORT="5678"
export STORAGE_SERVER_1_PORT="1967"
export STORAGE_SERVER_2_PORT="1968"

bash -c "cd run/naming_server \
  && ADDRESS=0.0.0.0:${NAMING_SERVER_PORT} ../../NamingServer" &
bash -c "cd run/naming_server \
  && ADDRESS=0.0.0.0:${NAMING_SERVER_PORT} \
     NAMING_SERVER_ADDRESS=${NAMING_SERVER_ADDRESS}:${NAMING_SERVER_PORT}
     ../../NamingServer" &
bash -c "cd run/naming_server \
  && ADDRESS=0.0.0.0:${NAMING_SERVER_PORT} \
  NAMING_SERVER_ADDRESS=${NAMING_SERVER_ADDRESS}:${NAMING_SERVER_PORT}
  ../../NamingServer" &

echo "Servers started. Type anything to stop. Don't forget to press Return in the end."
read -r var1
pkill NamingServer
pkill StorageServer
echo "All servers killed."

