package db

import (
	"SK-builder/internal/conf"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
)

// Data .
type Data struct {
	mysql *sql.DB
}

// // NewData .
// func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
// 	cleanup := func() {
// 		log.NewHelper(logger).Info("closing the data resources")
// 	}
// 	return &Data{}, cleanup, nil
// }

func NewMysql(data *conf.Data, logger log.Logger) (*Data, func(), error) {
	// db, _ := sql.Open("mysql", data.Mysql.Dsn)
	// //设置数据库最大连接数
	// db.SetConnMaxLifetime(100)
	// //设置上数据库最大闲置连接数
	// db.SetMaxIdleConns(10)
	// //验证连接
	// if err := db.Ping(); err != nil {
	// 	fmt.Println("open database fail")
	// 	log.NewHelper(logger).Info("open database fail")
	// 	return nil, nil, err
	// }

	return &Data{
			//mysql: db,
		}, func() {
			log.NewHelper(logger).Info("closing the data resources")
			//db.Close()
		}, nil
}
