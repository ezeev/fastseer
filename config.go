package fastseer

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ServerPort int `yaml:"serverPort"`
	//AppDomain                   string            `yaml:"appDomain"`
	ShopifyApiKey               string            `yaml:"shopifyApiKey"`
	ShopifyApiSecret            string            `yaml:"shopifyApiSecret"`
	ScriptTagDomain             string            `yaml:"scriptTagDomain"`
	DbOptions                   map[string]string `yaml:"dbOptions"`
	DbImpl                      string            `yaml:"dbImpl"`
	DefaultIndexAddress         string            `yaml:"defaultIndexAddress"`
	IndexingWorkerServices      []string          `yaml:"indexingWorkerServices"`
	IndexingWorkerEndpoint      string            `yaml:"indexingWorkerEndpoint"`
	IndexingWorkerStatsEndpoint string            `yaml:"indexingWorkerStatsEndpoint"`
	SearchImpl                  string            `yaml:"searchImpl"`
	JwtSecret                   string            `yaml:"jwtSecret"`
}

func LoadConfigFromFile(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err) // kill app is config can't load
	}
	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err) // can't read config
	}
	return &conf
}
