Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
Timeout: 2000

Auth:
  AccessSecret: # set your access-secret
  AccessExpire: 259200

CROSConf:
  Address: '*'

I18nConf:
  Dir: # set your dir

Log:
  ServiceName: {{.serviceName}}ApiLogger
  Mode: file
  Path: /home/data/logs/{{.serviceName}}/api
  Level: info
  Compress: false
  Encoding: json
  KeepDays: 7
  StackCoolDownMillis: 100

NacosConf:
  ServerAddr: # set your addr
  ServerPort: 8848
  DataId: "config"
  NamespaceId: "public"
  Group: "DEFAULT_GROUP"
  ApiServiceName: "core.api"
  RpcServiceName: "core.rpc"
  ClusterName: "DEFAULT"
  Username: "nacos"
  Password: "nacos"
  LogDir: "/home/data/logs/nacos"
  CacheDir: "/home/data/cache/nacos"
  LogLevel: "debug"


DatabaseConf:
  Type: mysql
  Host: # set your addr
  Port: 3306
  DBName: # set your database name
  Username: # set your username
  Password: # set your password
  MaxOpenConn: 100
  SSLMode: disable
  CacheTime: 5

RedisConf:
  Host: # set your addr
  Pass: # set your password
  Db: 0

CasbinConf:
  ModelText: |
    [request_definition]
    r = sub, obj, act
    [policy_definition]
    p = sub, obj, act
    [role_definition]
    g = _, _
    [policy_effect]
    e = some(where (p.eft == allow))
    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act

Telemetry:
  Name: core-api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger

Prometheus:
  Host: 0.0.0.0
  Port: 4008
  Path: /metrics