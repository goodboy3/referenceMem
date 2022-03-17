package sprMgr

import (
	"fmt"

	"github.com/coreservice-io/RedisSpr"
	"github.com/goodboy3/referenceMem/basic"
)

var instanceMap = map[string]*RedisSpr.SprJobMgr{}

func GetInstance() *RedisSpr.SprJobMgr {
	return instanceMap["default"]
}

func GetInstance_(name string) *RedisSpr.SprJobMgr {
	return instanceMap[name]
}

func Init(redisConfig *RedisSpr.RedisConfig) error {
	return Init_("default", redisConfig)
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, redisConfig *RedisSpr.RedisConfig) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("spr instance <%s> has already initialized", name)
	}

	if redisConfig.Addr == "" {
		redisConfig.Addr = "127.0.0.1"
	}
	if redisConfig.Port == 0 {
		redisConfig.Port = 6379
	}
	//////// ini spr job //////////////////////

	spr, err := RedisSpr.New(RedisSpr.RedisConfig{
		Addr:     redisConfig.Addr,
		Port:     redisConfig.Port,
		Password: redisConfig.Password,
		UserName: redisConfig.UserName,
		Prefix:   redisConfig.Prefix,
		UseTLS:   redisConfig.UseTLS,
	})

	if err != nil {
		return err
	}

	spr.SetULogger(basic.Logger)

	instanceMap[name] = spr

	return nil
}
