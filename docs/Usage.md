# 使用教程

## 服务管理

### 启动服务

```bash
systemctl start v2node
# 或
v2node start
```

### 停止服务

```bash
systemctl stop v2node
# 或
v2node stop
```

### 重启服务

```bash
systemctl restart v2node
# 或
v2node restart
```

### 查看状态

```bash
systemctl status v2node
# 或
v2node status
```

### 查看日志

```bash
journalctl -u v2node -f
# 或
v2node log
```

## 配置热更新

```bash
# 修改配置文件后
v2node reload
```

## 命令行参数

```bash
# 使用默认配置
v2node server

# 指定配置文件
v2node server -c /path/to/config.json

# 显示帮助
v2node server --help
```