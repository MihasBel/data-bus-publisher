# Data Bus Publisher

Data Bus Publisher is a component of the [Data Bus](https://github.com/MihasBel/data-bus) project, which is responsible
for subscribing to messages in Kafka and providing them to consumers through a gRPC interface.

## Functionality

Data Bus Publisher provides the following features:

1. Subscribing to messages in Apache Kafka.
2. Providing received messages to consumers via a gRPC interface.
3. Messages are transferred in a serialized format for flexible data transmission.
4. High performance and scalability due to asynchronous architecture.
5. Support for gRPC and HTTP/2 for efficient, high-performance data streaming.

## Getting Started

### Prerequisites

- Go 1.2 or higher
- Docker and Docker Compose

### Installation