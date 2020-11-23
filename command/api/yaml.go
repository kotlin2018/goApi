package api

var yaml = `# casbin configuration
casbin:
  model-path: './resource/rbac_model.conf'

# jwt configuration
jwt:
  signing-key: 'qmPlus'

# mysql connect configuration
mysql:
  username: root
  password: 'root'
  path: '127.0.0.1:3306'
  db-name: 'qmPlus'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  max-idle-conns: 10
  max-open-conns: 10
  log-mode: false

#sqlite 配置
sqlite:
  path: db.db
  log-mode: true
  config: 'loc=Asia/Shanghai'

# oss configuration

# 切换本地与七牛云上传，分配头像和文件路径
localupload:
  local: false
  avatar-path: uploads/avatar
  file-path: uploads/file

# 请自行七牛申请对应的 公钥 私钥 bucket 和 域名地址
qiniu:
  access-key: '25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ'
  secret-key: 'pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY'
  bucket: 'qm-plus-img'
  img-path: 'http://qmplusimg.henrongyi.top'

# redis configuration
redis:
  addr: '127.0.0.1:6379'
  password: ''
  db: 0

# system configuration
system:
  use-multipoint: false
  env: 'public'  # Change to "develop" to skip authentication for development mode
  addr: 8888
  db-type: "mysql"  # support mysql/sqlite

# captcha configuration
captcha:
  key-long: 6
  img-width: 240
  img-height: 80

# logger configuration
log:
  prefix: '[GIN-VUE-ADMIN]'
  log-file: true
  stdout: 'DEBUG'
  file: 'DEBUG'
`
