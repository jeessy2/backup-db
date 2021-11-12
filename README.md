<a href="https://github.com/jeessy2/backup-db/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/release/jeessy2/backup-db.svg?logo=github&style=flat-square"></a>
# 数据库备份工具 [English](README-EN.md)
  原理：Docker容器中安装postgres-client和mysql-client，并加入本备份工具，增强备份功能。
  - [x] 支持自定义命令
  - [x] 网页中配置，简单又方便
  - [x] 支持多个项目备份，最多16个
  - [x] 支持备份后的文件另存到对象存储(在也怕硬盘坏了)
  - [x] 每日凌晨自动备份
  - [x] 可设置备份文件最大保存天数
  - [x] 可设置登陆用户名密码，默认为空
  - [x] webhook通知

## docker中使用
- 运行docker容器
  ```
  docker run -d \
    --name backup-db \
    --restart=always \
    -p 9977:9977 \
    -v /opt/backup-db-files:/app/backup-db-files \
    jeessy/backup-db
  ```
- 登录 http://your_docker_ip:9977 并配置
  ![avatar](https://raw.githubusercontent.com/jeessy2/backup-db/master/backup-db-web.png)


## 备份脚本参考
 - postgres

    |  说明   | 备份脚本  |
    |  ----  | ----  |
    | 备份单个  | PGPASSWORD="password" pg_dump --host 192.168.1.11 --port 5432 --dbname db-name --user postgres --clean --create --file #{DATE}.sql |
    | 备份全部  | PGPASSWORD="password" pg_dumpall --host 192.168.1.11 --port 5432 --user postgres --clean --file #{DATE}.sql |
    | 还原  | psql -U postgres -f 2021-11-12_10_29.sql |

 -  mysql/mariadb

    |  说明   | 备份脚本  |
    |  ----  | ----  |
    | 备份单个  | mysqldump -h192.168.1.11 -uroot -p123456 db-name > #{DATE}.sql |
    | 备份全部  | mysqldump -h192.168.1.11 -uroot -p123456 --all-databases > #{DATE}.sql |
    | 还原  | mysql -uroot -p123456 db-name <2021-11-12_10_29.sql |

## webhook
- 支持webhook, 备份更新成功或不成功时, 会回调填写的URL
- 支持的变量

  |  变量名   | 说明  |
  |  ----  | ----  |
  | #{projectName}  | 项目名称 |
  | #{fileName}  | 备份后的文件名称 |
  | #{fileSize}  | 文件大小 (MB) |
  | #{result}  | 备份结果（成功/失败） |

- RequestBody为空GET请求，不为空POST请求
- Server酱: `https://sc.ftqq.com/[SCKEY].send?text=#{projectName}项目备份#{result},文件名:#{fileName},文件大小:#{fileSize}`
- Bark: `https://api.day.app/[YOUR_KEY]/#{projectName}项目备份#{result},文件名:#{fileName},文件大小:#{fileSize}`
- 钉钉:
  - 钉钉电脑端 -> 群设置 -> 智能群助手 -> 添加机器人 -> 自定义
  - 只勾选 `自定义关键词`, 输入的关键字必须包含在RequestBody的content中, 如：`项目备份`
  - URL中输入钉钉给你的 `Webhook地址`
  - RequestBody中输入 `{"msgtype": "text","text": {"content": "#{projectName}项目备份#{result},文件名:#{fileName},文件大小:#{fileSize}"}}`

## 说明
  - v1版本开始发生重要变化，不兼容0.0.x
  - v1后开始使用web方式来配置
  - 如要加入https，可通过nginx代理
  - v2版本后，一个镜像同时支持postgres/mysql
