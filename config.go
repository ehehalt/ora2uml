package ora2uml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	User           ConfigUser     `json:"User"`
	Database       ConfigDatabase `json:"Database"`
	Tables         []ConfigTable  `json:"Tables"`
	ColumnsIgnored []string       `json:"ColumnsIgnored"`
}

type ConfigUser struct {
	UserId   string `json:"UserId"`
	Password string `json:"Password"`
}

type ConfigDatabase struct {
	Host        string `json:"Host"`
	Port        int    `json:"Port"`
	ServiceName string `json:"ServiceName"`
}

type ConfigTable struct {
	Owner string   `json:"Owner"`
	Name  string   `json:"Name"`
	Tags  []string `json:"Tags"`
	Color int      `json:"Color"`
}

func Read(filename string) (Config, error) {
	var config Config

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}

	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return config, err
	}

	return config, nil
}

func (config Config) ConnectionString() string {
	return fmt.Sprintf("%s/%s@%s", config.User.UserId, config.User.Password, config.Database.ServiceName)
}
