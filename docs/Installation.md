# 安装指南

## 系统要求

### 支持的操作系统

| 系统 | 最低版本 |
|------|----------|
| CentOS | 7 |
| Ubuntu | 16 |
| Debian | 8 |
| Alpine | 3.10 |

### 硬件要求

- CPU: 1核以上
- 内存: 512MB以上
- 带宽: 100Mbps以上

## 一键安装

```bash
wget -N https://raw.githubusercontent.com/Mcloud136/v2node/master/script/install.sh && bash install.sh
```

## 手动安装

### 步骤 1：下载二进制

```bash
mkdir -p /usr/local/v2node
cd /usr/local/v2node
curl -sL "https://github.com/Mcloud136/v2node/releases/latest/download/v2node-linux-amd64.zip" | unzip -
chmod +x v2node
```

### 步骤 2：配置文件

```bash
mkdir -p /etc/v2node
cat > /etc/v2node/config.json << EOF
{
    "Log": {
        "Level": "info"
    },
    "Nodes": [
        {
            "ApiHost": "https://your-panel.com/",
            "NodeID": 1,
            "ApiKey": "your-secret-key"
        }
    ]
}
EOF
```

### 步骤 3：创建服务

```bash
cat > /etc/systemd/system/v2node.service << EOF
[Unit]
Description=v2node Service
After=network.target

[Service]
ExecStart=/usr/local/v2node/v2node server
Restart=always

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable v2node
systemctl start v2node
```