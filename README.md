# Message-Loop
This is a small research project using Golang, Kafka, grpc and protobuf.

## Project Functionality

This loop is composed of two services that will pass one piece of data 3 times.

- Alpha -> Beta via Kafka (perform simple transformation)
- Beta -> Alpha via gRPC (perform simple transformation)
- Alpha -> Beta via protobuf (perform simple transformation and log)

## Project Goals

This project outlines simple methods to utilize the following technologies using Golang:

- Inbound Kafka Messages
- Outbound Kafka Messages
- Protobuf
- gRPC

## Deployment and Execution

