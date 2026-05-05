# 故障排查

## 常见问题

### Q1: 启动失败

**错误信息**：`Exec format error`

**原因**：二进制文件与系统架构不匹配

**解决方案**：
```bash
uname -m  # 查看系统架构
# 下载对应架构的版本
```

### Q2: 无法连接面板

**错误信息**：`Connection refused`

**解决方案**：
1. 检查 ApiHost 是否正确
2. 检查网络连接
3. 检查防火墙设置

### Q3: 证书错误

**错误信息**：`SSL certificate error`

**解决方案**：
```bash
# 更新系统证书
apt-get update && apt-get install ca-certificates
```

## 日志分析

### 日志位置

- Systemd 日志：`journalctl -u v2node`
- 文件日志：`/var/log/v2node.log`（如果配置了）

### 常见日志级别

| 级别 | 说明 |
|------|------|
| DEBUG | 详细调试信息 |
| INFO | 一般信息 |
| WARNING | 警告信息 |
| ERROR | 错误信息 |

## 调试模式

```bash
# 修改配置文件
{
    "Log": {
        "Level": "debug"
    }
}

# 重启服务
v2node restart
```