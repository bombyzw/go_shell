package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

var (
	Cfg Config
)

type Config struct {
	Shells     []string          `toml:"shells"`
	SyncShells []string          `toml:"sync_shells"`
	ShellMap   map[string]ShInfo `toml:"shell_map"`
}

type ShInfo struct {
	Name string
	Cmd  string
}

func InitConfig(configpath string) error {
	_, err := toml.DecodeFile(configpath, &Cfg)
	if err != nil {
		return err
	}

	log.Println(Cfg)

	return nil
}
