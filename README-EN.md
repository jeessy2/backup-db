# backup-db
  A database backup tool with web interfaces.
  - [x] Support custom commands.
  - [x] Obsolete files will be deleted automatically.
  - [ ] Support the backup files copy to simple data storage(s3).
  - [x] Automatic backup in everyday night.
  - [x] Webhook support

## use in docker
  ```
    docker run -d \
    --name backup-db \
    --restart=always \
    -p 9977:9977 \
    -v /opt/backup-db-files:/app/backup-db-files \
    jeessy/backup-db
  ```

  ![avatar](https://raw.githubusercontent.com/jeessy2/backup-db/master/backup-db-web.png)

