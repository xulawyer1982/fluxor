package configgen

const configTemplate = `mixed-port:
tproxy-port:
allow-lan: true
ipv6: true
bind-address: '*'
mode: rule
log-level: silent
unified-delay: true
external-controller: ''
external-controller-unix: '/var/apps/Fluxor/target/core.sock'
external-ui:
external-ui-url:
secret: ''

routing-mark: 255
find-process-mode: strict
client-fingerprint: chrome

profile:
  store-selected: true
  store-fake-ip: true

sniffer:
  enable: true
  sniff:
    HTTP:
      ports: [80, 8080-8880]
      override-destination: true
    TLS:
      ports: [443, 8443]
    QUIC:
      ports: [443, 8443]
  skip-domain:
    - "+.push.apple.com"

tun:
  enable: false
  stack: mixed
  # 显式不给 TUN 分配 IPv6 地址。宿主 net.ipv6.conf.default.disable_ipv6=1 时，
  # 新建 tun 接口同样禁用 IPv6，sing-tun 下发该地址会被内核拒绝（EACCES），
  # 整个 TUN 启动失败并报 "configure tun interface: permission denied"。
  # 留空则 sing-tun 跳过 IPv6 地址与 IPv6 DNS 劫持规则（0.4.28+ 对空值有守卫）。
  inet6-address: []
  dns-hijack:
    - "any:53"
    - "tcp://any:53"
  auto-route: true
  auto-redirect: true
  auto-detect-interface: true

geodata-mode: false
geo-auto-update: true
geo-update-interval: 24
`
