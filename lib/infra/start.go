package infra

import (
	"context"
	"fmt"

	"blogs/conf"
	"blogs/lib/infra/mysql"
	"blogs/lib/log"
)

func Start(ctx context.Context) error {
	var err error

	mysqlConf := &mysql.DBConf{}
	if err = conf.Parser(mysqlConf, conf.MysqlConfPath); err != nil {
		return err
	}
	mysqlConf.Logger = log.NewGormLogger()

	if MysqlClient, err = mysql.InitMysql(ctx, mysqlConf); err != nil {
		return err
	}
	ret := map[string]interface{}{}
	err = MysqlClient.Raw("show databases").Scan(&ret).Error
	fmt.Println(ret, err)
	return nil
}
