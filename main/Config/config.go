package Config

import (
	"NiuGame/main/common"
	"bufio"
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Name string
}

func GetConfig() *Config {
	return cfg
}

var cfg *Config = nil

func Init(name string) error {
	c := Config{
		Name: name,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		log.Println(err)
		return err
	}
	// 监控配置文件变化并热加载程序
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		//viper.AddConfigPath("./Chapter16/conf/") //goland debug的配置
		viper.AddConfigPath(common.Config_Path) //命令行的配置
		viper.SetConfigName(common.Config_FileName)
	}
	viper.SetConfigType(common.Config_FileType)  // 设置配置文件格式为YAML
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path) //读取文件
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader) //解析json
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
