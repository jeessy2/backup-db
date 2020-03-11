# 数据库备份工具
  原理：在原生的docker镜像基础上，加入一备份工具，增强备份功能。
  提供postgres, mysql5镜像，可直接使用，如有需要请提issues。
  - [X] 可以自行构建docker镜像，支持不同的数据库及不同的版本，如mysql8, oracle, sqlserver2017+等等
  - [X] 支持自定义命令
  - [X] 可以把备份后的文件存入另一台服务器
  - [X] 备份失败邮件通知
  - [X] 每日凌晨自动备份
  - [X] 可设置备份文件最大保存天数(最少3天)

# backup databases
  Support all databases and the database images can be find in docker.
  - [X] Support for custom backup commands.
  - [X] Obsolete files will be deleted automatically.
  - [X] You can copy the backup files to another server.
  - [x] Send email when backup failed.
  - [x] Automatic backup in every night.
  - [x] The maximum number of days to save backup files can be set (at least 3 days).

## docker 环境变量说明
```
backup_server_ip 不填默认为二次备份的服务器
backup_server_port 二次备份服务器的端口
backup_project_name 项目名称，一般就是数据库名称。
backup_command 备份命令，必须包含#{DATE}
max_save_days 备份文件最大保存天数
notice_email 异常通知的邮箱
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
jeessy/backup-db:postgres-0.0.3
```

## client (postgress)
```
docker run -d \
--name backup-db-name \
--restart=always \
-v /opt/backup-files:/app/backup-files \
-e backup_server_ip=192.168.1.76 \
-e backup_server_port=9977 \
-e backup_project_name=db-name \
-e backup_command="pg_dump -a \"host=192.168.1.11 port=5433 user=postgres password=password dbname=db-name\" > #{DATE}.sql" \
-e max_save_days=30 \
-e notice_email=277172506@qq.com \
-e smtp_host=smtp.office365.com \
-e smtp_port=587 \
-e smtp_username=backup-db-docker@outlook.com \
-e smtp_password=kLhHbTC6Ak5B2hw \
jeessy/backup-db:postgres-0.0.3
```

## client (mysql5)
```
docker run -d \
--name backup-db-name \
--restart=always \
-v /opt/backup-files:/app/backup-files \
-e backup_server_ip=192.168.1.76 \
-e backup_server_port=9977 \
-e backup_project_name=db-name \
-e backup_command="mysqldump -h192.168.1.9 -uroot -p123456 db-name > #{DATE}.sql" \
-e max_save_days=30 \
-e notice_email=277172506@qq.com \
-e smtp_host=smtp.office365.com \
-e smtp_port=587 \
-e smtp_username=backup-db-docker@outlook.com \
-e smtp_password=kLhHbTC6Ak5B2hw \
jeessy/backup-db:mysql5-0.0.3
```

## build docker images (You may not need to build docker images, if you use postgres or mysql5)
```
# first git clone
# change Dockerfile
# build docker images
docker build . -f Dockerfile_mysql -t jeessy/backup-db:mysql5-0.0.3
docker build . -f Dockerfile -t jeessy/backup-db:postgres-0.0.3
```
