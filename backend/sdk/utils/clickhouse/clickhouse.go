package clickhouse

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Click struct {
	connection driver.Conn
}

func New(database, username, password string) *Click {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%v", "", "")},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
	})

	if err != nil {
		panic(err)
	}

	err = conn.Exec(context.Background(), "", "")

	if err != nil {
		panic(err)
	}

	return &Click{
		connection: conn,
	}
}

func (c *Click) Write(p []byte) (n int, err error) {
	err = c.connection.Exec(context.Background(), string(p))

	return len(p), err
}

func (c *Click) Close() {
	_ = c.connection.Close()
}
