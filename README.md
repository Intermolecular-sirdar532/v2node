# v2node

A high-performance v2board backend based on modified xray-core.

## 特性

- 支持多种协议：VLESS、VMESS、Trojan、Shadowsocks、Hysteria2、TUIC、AnyTLS
- 高性能并发设计，支持高负载场景
- 完善的流量限制和设备管理
- 自动证书管理和更新
- 配置热更新，无需重启服务
- 优雅的错误处理和日志管理

## 改进内容

### 性能优化

- 使用 `atomic.Pointer` 实现无锁读取，提升并发性能
- `sync.Map` 替代普通 map，支持高效并发读写
- 复合 key 设计简化数据结构，减少内存开销
- `determineSpeedLimit` 逻辑优化，代码减少 60%

### 类型安全

- 使用类型断言替代反射，提高代码安全性和执行效率
- 添加 `int64` 类型支持，增强兼容性

### 配置管理

- 新增配置验证机制，启动前校验参数有效性
- 端口可用性检查，避免启动失败
- 配置热更新失败时保持原配置运行，确保服务连续性
- 支持增量配置更新

### 资源管理

- 统一资源生命周期管理
- pprof 服务支持优雅关闭
- 日志文件句柄统一管理，避免资源泄漏

### 文档完善

- 添加详细的用户手册 `DOCUMENTATION.md`
- 包含完整的配置说明和使用指南
- 常见问题解答

## 系统要求

- Linux 系统（CentOS 7+/Ubuntu 16+/Debian 8+/Alpine 3.10+）
- 推荐 2 核 CPU / 1GB 内存
- 支持 x86_64、ARM64 架构

## 安装

### 一键安装

```bash
wget -N https://raw.githubusercontent.com/Mcloud136/v2node/master/script/install.sh && bash install.sh
```

### 带参数安装

```bash
bash install.sh --api-host https://your-panel.com/ --node-id 1 --api-key your-secret-key
```

### 手动安装

```bash
mkdir -p /usr/local/v2node /etc/v2node
cd /usr/local/v2node
curl -sL https://github.com/Mcloud136/v2node/releases/latest/download/v2node-linux-amd64.zip | unzip -
chmod +x v2node
```

## 配置

配置文件路径：`/etc/v2node/config.json`

```json
{
    "Log": {
        "Level": "info",
        "Output": "",
        "Access": "none"
    },
    "Nodes": [
        {
            "ApiHost": "https://your-panel.com/",
            "NodeID": 1,
            "ApiKey": "your-secret-key",
            "Timeout": 15
        }
    ],
    "PprofPort": 0
}
```

## 服务管理

```bash
v2node start          # 启动服务
v2node stop           # 停止服务
v2node restart        # 重启服务
v2node status         # 查看状态
v2node log            # 查看日志
v2node generate       # 生成配置
v2node update         # 更新版本
v2node version        # 查看版本
```

## 协议支持

| 协议 | 状态 |
|------|------|
| VLESS | ✅ |
| VMESS | ✅ |
| Trojan | ✅ |
| Shadowsocks | ✅ |
| Hysteria2 | ✅ |
| TUIC | ✅ |
| AnyTLS | ✅ |

## 构建

```bash
GOEXPERIMENT=jsonv2 go build -v -o build_assets/v2node -trimpath -ldflags "-X 'github.com/Mcloud136/v2node/cmd.version=1.0.0' -s -w -buildid="
```

## 注意事项

- 本项目需要搭配 [修改版 V2board](https://github.com/wyx2685/v2board) 使用
- 确保服务器时间同步，否则可能导致连接问题
- 建议使用官方提供的一键安装脚本

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

## 相关项目

- [V2board](https://github.com/wyx2685/v2board)
- [Xray-core](https://github.com/XTLS/Xray-core)