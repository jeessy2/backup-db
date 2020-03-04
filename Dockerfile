
# build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.3/main" > /etc/apk/repositories \
    && echo "https://mirrors.ustc.edu.cn/alpine/v3.3/community" >> /etc/apk/repositories \
    # && apk update \
    # && go get -d -v . \
    # && go install -v . \
    && go build -v .

# final stage
FROM postgres
WORKDIR /app
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
COPY --from=builder /app /app
ENTRYPOINT /app/backup-db
LABEL Name=backup-db Version=0.0.1
