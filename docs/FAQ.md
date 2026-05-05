# 常见问题

## 基础问题

### Q: 这个项目是什么？

A: v2node 是一个基于修改版 xray-core 的 V2board 节点服务端。

### Q: 需要搭配什么面板使用？

A: 需要搭配 [修改版 V2board](https://github.com/wyx2685/v2board) 使用。

### Q: 支持哪些协议？

A: 支持 VLESS、VMESS、Trojan、Shadowsocks、Hysteria2、TUIC、AnyTLS。

## 配置问题

### Q: ApiHost 需要什么格式？

A: 需要以 `/` 结尾，例如：`https://your-panel.com/`

### Q: 如何配置多节点？

A: 在 Nodes 数组中添加多个节点配置即可。

## 性能问题

### Q: 如何优化性能？

A: 
1. 降低日志级别为 warning
2. 确保服务器资源充足
3. 使用最新版本

### Q: 支持多少并发连接？

A: 取决于服务器配置，建议使用 2核以上 CPU 和 1GB 以上内存。

## 安全问题

### Q: 如何保护 ApiKey？

A: 
1. 不要泄露 ApiKey
2. 使用强密码
3. 定期更换

### Q: 是否支持 HTTPS？

A: HTTPS 配置在面板中完成，节点会自动获取证书。