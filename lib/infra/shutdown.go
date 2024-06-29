package infra

import (
	"context"

	"blogs/lib/infra/mysql"
)

func Shutdown(ctx context.Context) error {
	_ = mysql.Shutdown(MysqlClient)

	return nil
}
