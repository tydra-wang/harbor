package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/goharbor/harbor/src/common/dao"
	"github.com/goharbor/harbor/src/common/models"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("POSTGRESQL_MAX_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("POSTGRESQL_MAX_OPEN_CONNS"))
	err := dao.UpgradeSchema(&models.Database{
		PostGreSQL: &models.PostGreSQL{
			Host:         os.Getenv("POSTGRESQL_HOST"),
			Port:         port,
			Username:     os.Getenv("POSTGRESQL_USERNAME"),
			Password:     os.Getenv("POSTGRESQL_PASSWORD"),
			Database:     os.Getenv("POSTGRESQL_DATABASE"),
			SSLMode:      os.Getenv("POSTGRESQL_SSLMODE"),
			MaxIdleConns: maxIdleConns,
			MaxOpenConns: maxOpenConns,
		},
	})
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		fmt.Fprint(os.Stderr, "ok")
	}
}
