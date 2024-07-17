package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	SERVERPORT  int
	DBPORT      int
	DBHOST      string
	DBUSER      string
	DBPASS      string
	DBNAME      string
	API_STORAGE string
	RDB_ADDR    string
	RDB_USER    string
	RDB_PASS    string
	ACCESID     string
	SECRETKEY   string
	ACCOUNTID  string
	BUCKET 	string
}

func InitConfig() *AppConfig {
	var res = new(AppConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *AppConfig {
	var res = new(AppConfig)

	godotenv.Load(".env")

	if val, found := os.LookupEnv("SERVERPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		res.SERVERPORT = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		res.DBPORT = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHOST = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUSER = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPASS = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBNAME = val
	}
	if val, found := os.LookupEnv("API_STORAGE"); found {
		res.API_STORAGE = val
	}
	if val, found := os.LookupEnv("RDB_ADDR"); found {
		res.RDB_ADDR = val
	}
	if val, found := os.LookupEnv("RDB_USER"); found {
		res.RDB_USER = val
	}
	if val, found := os.LookupEnv("RDB_PASS"); found {
		res.RDB_PASS = val
	}

	if val, found := os.LookupEnv("ACCESID"); found {
		res.ACCESID = val
	}
	if val, found := os.LookupEnv("SECRETKEY"); found {
		res.SECRETKEY= val
	}
	if val, found := os.LookupEnv("ACCOUNTID"); found {
		res.ACCOUNTID = val
	}
	if val, found := os.LookupEnv("BUCKET"); found {
		res.BUCKET = val
	}
	return res
}