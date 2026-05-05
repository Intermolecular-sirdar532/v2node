# 配置说明

## 配置文件位置

```
/etc/v2node/config.json
```

## 完整配置示例

```json
{
    "Log": {
        "Level": "info",
        "Output": "/var/log/v2node.log",
        "Access": "/var/log/v2node_access.log"
    },
    "Nodes": [
        {
            "ApiHost": "https://your-panel.com/",
            "NodeID": 1,
            "ApiKey": "your-secret-key",
            "Timeout": 15,
            "RetryCount": 3
        }
    ],
    "PprofPort": 6060
}
```

## 参数详解

### Log 配置

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| Level | string | info | 日志级别：debug/info/warning/error |
| Output | string | "" | 日志输出文件路径 |
| Access | string | "none" | 访问日志路径 |

### Nodes 配置

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| ApiHost | string | ✅ | 面板 API 地址 |
| NodeID | int | ✅ | 节点 ID |
| ApiKey | string | ✅ | 通讯密钥 |
| Timeout | int | ❌ | 超时时间(秒) |
| RetryCount | int | ❌ | 重试次数 |