# backup databases
  Support all databases and the database images can be find in docker.
  - [X] Support for custom backup commands.
  - [X] Obsolete files will be deleted automatically.
  - [X] You can copy the backup files to another server.
  - [x] Send email when backup failed.
  - [x] Automatic backup in every night.
  - [x] The maximum number of days to save backup files can be set (at least 3 days).

# 数据库备份工具
  原理：在原生的docker镜像基础上，加入一备份工具，增强备份功能。
  - [X] 支持的数据库需有docker镜像。如postgres, mysql
  - [X] 支持自定义命令
  - [X] 可以把备份后的文件存入另一台服务器
  - [X] 备份失败邮件通知
  - [X] 每日凌晨自动备份
  - [X] 可设置备份文件最大保存天数(最少3天)

## build docker images (You may not need to build docker images, if you use postgres)
```
# build docker and run
docker build . -t jeessy/backup-db
docker run -d jeessy/backup-db
```

## client
```
docker run -d \
--name backup-db_name \
--restart=always \
-v /opt/backup-files:/app/backup-files \
-e backup_server_ip=192.168.1.76 \
-e backup_server_port=9977 \
-e backup_project_name=db_name \
-e backup_command="pg_dump -a \"host=192.168.1.11 port=5433 user=postgres password=password dbname=db_name\" > #{DATE}.sql" \
-e max_save_days=30 \
-e notice_email=277172506@qq.com \
-e smtp_host=smtp.office365.com \
-e smtp_port=587 \
-e smtp_username=backup-db-docker@outlook.com \
-e smtp_password=kLhHbTC6Ak5B2hw \
jeessy/backup-db-postgres
```

## server
```
docker run -d \
--name backup-server \
--restart=always
-p 9977:9977 \
-v /opt/backup-files:/app/backup-files \
-e backup_server_port=9977 \
-e max_save_days=30 \
-e notice_email=277172506@qq.com \
-e smtp_host=smtp.office365.com \
-e smtp_port=587 \
-e smtp_username=backup-db-docker@outlook.com \
-e smtp_password=kLhHbTC6Ak5B2hw \
jeessy/backup-db-postgres
```