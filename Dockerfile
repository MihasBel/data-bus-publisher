FROM golang:1.20.4 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
  build-essential \
  git \
  wget \
  bash \
  g++ \
  make \
  libssl-dev \
  zlib1g-dev \
  liblz4-dev \
  libzstd-dev \
  pkg-config

RUN wget https://github.com/edenhill/librdkafka/archive/refs/tags/v1.9.0.tar.gz && \
    tar xzf v1.9.0.tar.gz && \
    cd librdkafka-1.9.0 && \
    ./configure --prefix /usr && \
    make && \
    make install

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -tags dynamic -installsuffix cgo -o main ./cmd/main.go

FROM golang:1.20.4

COPY --from=builder /usr/lib/librdkafka.so* /usr/lib/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .


CMD ["./main"]
