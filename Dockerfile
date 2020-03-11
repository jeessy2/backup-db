# build stage
FROM golang AS builder
WORKDIR /app
COPY . .
RUN go get -d -v . \
    && go install -v . \
    && go build -v .

# final stage
# you can replace "postgres" to other images, emample: "mysql"
FROM postgres
WORKDIR /app
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
COPY --from=builder /app/backup-db /app/backup-db
ENTRYPOINT /app/backup-db
LABEL Name=backup-db Version=0.0.2
