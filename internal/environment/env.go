package environment

import (
	"fmt"
	"os"
	"strconv"
)

type Env struct {
	ServerPort    int
	MysqlDatabase string
	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlPort     int
}

func LoadEnv() (*Env, error) {
	env := Env{}

	env.ServerPort = loadEnvAsIntWithDefault("PORT", 8080)

	if v, err := loadEnvAsString("MYSQL_DATABASE"); err != nil {
		return nil, err
	} else {
		env.MysqlDatabase = v
	}

	if v, err := loadEnvAsString("MYSQL_USER"); err != nil {
		return nil, err
	} else {
		env.MysqlUser = v
	}

	if v, err := loadEnvAsString("MYSQL_PASSWORD"); err != nil {
		return nil, err
	} else {
		env.MysqlPassword = v
	}

	if v, err := loadEnvAsString("MYSQL_HOST"); err != nil {
		return nil, err
	} else {
		env.MysqlHost = v
	}

	if v, err := loadEnvAsInt("MYSQL_PORT"); err != nil {
		return nil, err
	} else {
		env.MysqlPort = v
	}

	return &env, nil
}

func loadEnvAsString(key string) (string, error) {
	if v := os.Getenv(key); v == "" {
		return "", fmt.Errorf("env[%s] is empty", key)
	} else {
		return v, nil
	}
}

func loadEnvAsInt(key string) (int, error) {
	if v, err := loadEnvAsString(key); err != nil {
		return 0, err
	} else {
		return strconv.Atoi(v)
	}
}

func loadEnvAsIntWithDefault(key string, defaultValue int) int {
	if v, err := loadEnvAsInt(key); err != nil {
		return defaultValue
	} else {
		return v
	}
}
