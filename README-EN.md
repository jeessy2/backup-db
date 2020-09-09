## backup databases
  Support all databases and the database images can be find in docker.
  - [X] Support for custom backup commands.
  - [X] Obsolete files will be deleted automatically.
  - [X] You can copy the backup files to another server.
  - [x] Send email when backup failed.
  - [x] The server checks the backup files that are not uploaded at 10 o'clock every day. If the backup files are not checked, an email notice will be sent
  - [x] Automatic backup in every night.
  - [x] The maximum number of days to save backup files can be set.

## server
```
docker run -d \
  --name backup-db-server \
  --restart=always \
  -p 9977:9977 \
  -v /opt/backup-files:/app/backup-files \
  jeessy/backup-db:v1.0.0-server
```

## client (postgress)
```
docker run -d \
  --name backup-db-postgres \
  --restart=always \
  -p 9977:9977 \
  -v /opt/backup-files:/app/backup-files \
  jeessy/backup-db:v1.0.0-postgres
```

## client (mysql5)
```
docker run -d \
  --name backup-db-mysql5 \
  --restart=always \
  -p 9977:9977 \
  -v /opt/backup-files:/app/backup-files \
  jeessy/backup-db:v1.0.0-mysql5
```

## client (mysql8)
```
docker run -d \
  --name backup-db-mysql8 \
  --restart=always \
  -p 9977:9977 \
  -v /opt/backup-files:/app/backup-files \
  jeessy/backup-db:v1.0.0-mysql8
```

## Release
```
git tag v0.0.x -m "xxx" 
git push --tags
```
