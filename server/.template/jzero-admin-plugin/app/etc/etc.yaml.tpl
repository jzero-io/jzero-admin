rest:
    name: {{ .APP }}-api
    host: 0.0.0.0
    port: 8001

log:
    serviceName: {{ .APP }}
    encoding: plain
    level: info
    mode: console

sqlx:
    driverName: "mysql"
    dataSource: "root:123456@tcp(127.0.0.1:3306)/jzero-admin?charset=utf8mb4&parseTime=True&loc=Local"

redis:
    host: "127.0.0.1:6379"
    type: "node"
    pass: "123456"