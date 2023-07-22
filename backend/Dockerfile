FROM golang:1.20-alpine AS builder

LABEL maintainer="Vic Shóstak <vic@shostak.dev> (https://shostak.dev/)"

# Move to working directory (/build).
WORKDIR /build
ENV GOPROXY=https://goproxy.cn,direct

# Copy and download dependency using go mod.
# COPY go.mod go.sum ./
# RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -ldflags="-s -w" -o oa .

# FROM scratch
FROM hub.fastonetech.com/bases/debian:bullseye

RUN sed -i -e 's/http:\/\/[^\/]*/http:\/\/mirrors.ustc.edu.cn/g' /etc/apt/sources.list \
    && apt update && apt install -y iputils-ping netcat iproute2 net-tools procps bsdmainutils tree lsof sudo wget gnupg htop \
    && rm -rf /var/lib/apt/lists/*

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/oa", "/build/.env", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/oa"]