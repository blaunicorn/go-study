# go-web

实战演练

# gin

gin 的基本用法
restful-api 增删改查

# gorm-postgres

go 使用 gorm 连接 postgres 的基本用法，并存储用户数据

# redis

go 客户端连接 redis 的用法，并缓存用户记录

# mongodb

go 客户端连接 mongodb，并存储访问记录

# Docker 搭建 Postgresql 数据库

# 使用 Docker 搭建 Postgresql 数据库，及设置远程访问

# 一、拉取 Postgresq 镜像

```
docker pull postgres
```

# 二、构建镜像容器

```
docker run -it --name postgres --restart always -e POSTGRES_PASSWORD='abc123' -e ALLOW_IP_RANGE=0.0.0.0/0 -v /home/postgres/data:/var/lib/postgresql -p 8030:5432 -d postgres
```

-itd 是三个参数：-i 和-t。前者表示打开并保持 stdout，后者表示分配一个终端（pseudo-tty）。此时如果使用 exit 退出，则容器的状态处于 Exit，而不是后台运行。如果想让容器一直运行，而不是停止，可以使用快捷键 ctrl+p ctrl+q 退出，此时容器的状态为 Up。-d 后台运行
–name : 自定义容器名称
POSTGRES_PASSWORD：数据库密码
-e ALLOW_IP_RANGE=0.0.0.0/0，这个表示允许所有 ip 访问，如果不加，则非本机 ip 访问不了
-v :进行映射,本地目录：容器内路径
-p：映射端口,宿主机端口：容器端口
最后是 镜像名称:端口号

# 三、进入 postgres 容器

```
docker exec -it postgres bash
// 或  docker attach container_name/container_id
```

浏览、启动、停止、重启容器命令：

```
[root@xxx ~]# docker ps -a
[root@xxx ~]# docker start container_name/container_id
[root@xxx ~]# docker stop container_name/container_id
[root@xxx ~]# docker restart container_name/container_id
```

后台启动一个容器后，如果想进入到这个容器，可以使用 attach 命令：

```
[root@xxx ~]# docker attach container_name/container_id
```

删除容器的命令：

```
[root@xxx ~]# docker rm container_name/container_id
```

# 四、切换当前用户，再登录数据库

将当前 root 切换成 postgres

```
su postgres
```

输入用户名，密码再命令执行完后，再根据提示输入

```
psql -U postgres -W
```

输入密码 abc123，登录成功
注：创建语句结束要记得加；

```
create user gitchat with password '123456';

 create database test;

 grant all privileges on database test to gitchat;

 \l  // 浏览数据库
```

# 五、设置远程访问许可（很重要）

共需要修改两个配置文件：pg_hba.conf、postgresql.conf，步骤分别如下：

1、修改 postgres 配置文件
首先，确认已进入容器，再按下列步骤操作：

1）、修改 pg_hba.conf 文件

#### A、用命令将 postgres 中的 pg_hba.conf 文件，复制到目录/home 中

```
docker cp postgres:/var/lib/postgresql/data/pg_hba.conf /home
```

#### B、用 Xftp 连接服务器，找到 home 下的 pg_hba.conf，并用记事本打开及修改用户的访问权限（#开头的行是注释内容）：

```
# TYPE DATABASE USER CIDR-ADDRESS METHOD

# "local" is for Unix domain socket connections only

local all all trust

# IPv4 local connections:

host all all 127.0.0.1/32 trust
_host all all 0.0.0.1/0 md5_

# IPv6 local connections:

host all all ::1/128 trust
```

#### C、用命令将修改后的文件，替换掉原来的配置文件

docker cp /home/pg_hba.conf postgres:/var/lib/postgresql/data
1
2）、修改 postgresql.conf 文件
如上述修改 pg_hba.conf 一样，先复件到/home 文件夹中，然后打开编辑，最后替换掉原配置文件。
定位到 #listen_addresses = ’localhost’，再将行开头都#去掉，并将行内容修改为 localhost 改成：\*

```
listen*addresses = ’*’
```

默认只接受来自本机 localhost 的连接请求，\_ 允许数据库服务器监听来自任何主机的连接请求。

2、设置防火墙
首先，执行 exit 命令，退出 postgres 用户，再执行一次 exit，退出容器。然后才可以进行防火墙设置。

1）、检查 firewalld 启动状态

```
systemctl status firewalld
```

若未启动，则执行启动命令：

```
systemctl start firewalld
```

2）、检查 firewall-cmd 运行状态

```
firewall-cmd --state
```

3）、防火墙正常，则分别执行以下命令，进行配置

```
firewall-cmd --zone=public --add-port=5432/tcp --permanent
```

```
firewall-cmd --reload
```

重启防火墙

```
systemctl stop firewalld.service
```

```
systemctl start firewalld.service
```

# redis 配置

一、cenOS 7 安装 redis
Redis 是一个开源的内存数据结构存储库，具有用于持久性的可选磁盘写入。它可以用作键值数据库，也可以用作缓存和消息代理。
Redis 具有内置事务，复制和对各种数据结构（如字符串，哈希，列表，集等）的支持。Redis 可以通过 Redis Sentinel 高度可用，
并支持使用 Redis Cluster 自动分区。

在本节中，您将添加 EPEL 存储库，然后使用它来安装 Redis。

1.1 添加 EPEL 存储库，并更新 YUM 以确认您的更改：
sudo yum install epel-release sudo yum update
1
1.2 安装 Redis：
sudo yum install redis
1
1.3 启动 Redis：
sudo systemctl start redis
1
1.4 可选：要在启动时自动启动 Redis：
sudo systemctl enable redis
1
1.5 验证 Redis 是否正在运行 redis-cli：
redis-cli ping
1
如果 Redis 正在运行，它将返回：PONG

二、Redis——配置允许远程连接
防火墙添加 6379 端口
确保防火墙开启了对 6379 端口的放行。
修改 redis.conf
在 redis.conf 中修改以下参数

# 允许任何主机连接、访问

bind 127.0.0.1 改为 bind 0.0.0.0

# 关闭保护模式

protected-mode yes 改为 protected-mode no

# 允许启动后在后台运行，即关闭命令行窗口后仍能运行

daemonize no 改为 daemonize yes

修改完成后，重启 redis

systemctl restart redis

```
127.0.0.1 => 0.0.0.0
port 6379
requirepass 123456
```

# Mongo

Mongo 和 Redis 一样，是内存数据库，对高并发的支持比 Redis 可用性更高，因为 Redis 是单线程的。

Mongo 的业务场景面向一些量级很大（未来可能会变得很大）的数据，下面我们在前面的 restful-redis-pg 的基础上，增加一个访问日志记录，这些日志被记入 Mongo。

Mongo 除了快以外，还有一个好处，如果 DB 和 Collection 不存在时，会自动创建。
mongodb://admin:123qwe@81.70.104.105:27017/testdb
