package clickhouse

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Click struct {
	connection driver.Conn
}

func New(database, username, password string) *Click {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%v", "localhost", "9000")},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		Debug: true,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(conn)

	err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS logs (
				timestamp DateTime,
				level String,
				message String
			) ENGINE = MergeTree()
			ORDER BY timestamp
	`)

	if err != nil {
		fmt.Println("MY ERROR", err)
		panic(err)
	}

	return &Click{
		connection: conn,
	}
}

func (c *Click) Write(p []byte) (n int, err error) {
	query := `INSERT INTO logs (timestamp, level, message) VALUES (?, ?, ?)`

	err = c.connection.Exec(context.Background(), query, time.Now(), zerolog.LevelDebugValue, string(p))

	return len(p), err
}

func (c *Click) Close() {
	_ = c.connection.Close()
}
