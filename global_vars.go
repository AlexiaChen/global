package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"gitlab.landui.cn/gomod/database"
	"go.uber.org/zap"
)

func init() {
}

var (
	Database    = database.Crud{} // 持久化数据源，对上层的名字，底层可能是MySQL、PgSQL、SQLite等等
	RedisClient *redis.Client     // 高速缓存
	WebPort     = ""              // web服务端口
	HttpClient  *resty.Client     // http服务
	Logger      *zap.Logger       // 日志收集器
	LoggerLevel int8              // 日志等级
	LogPath     string            // 日志文件存放路径
	ContextName string            // 上下文名称
	Validate    *validator.Validate
	Trans       ut.Translator
)
