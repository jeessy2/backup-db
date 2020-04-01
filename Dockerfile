# build stage(Only for server)
FROM golang AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go get -d -v . \
    && go install -v . \
    && go build -v .

# final stage, build server
FROM centos
WORKDIR /app
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
COPY --from=builder /app/backup-db /app/backup-db
ENTRYPOINT /app/backup-db
LABEL Name=backup-db-server Version=0.0.7
