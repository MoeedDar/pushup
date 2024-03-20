package single

import (
	"context"
	"pushup/globals"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func init() {
	config, _ := pg.ParseURL(globals.Env["POSTGRES_CONN_URL"])
	config.TLSConfig.InsecureSkipVerify = false
	config.TLSConfig.ServerName = globals.Env["POSTGRES_SERVER_NAME"]
	DB = pg.Connect(config)
	if err := DB.Ping(context.Background()); err != nil {
		panic(err)
	}
}
