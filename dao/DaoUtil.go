package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)
func OpenDb()*sql.DB{
	db,err:=sql.Open("mysql","root:327514@/smy?charset=utf8&parseTime=true&loc=Local")
	DbError(err)
	return db
}
func DbError(err error){
	if err!=nil{
		logger,_:=zap.NewProduction()
		defer logger.Sync()
		logger.Error("db Error",zap.String("msg",err.Error()))
	}
}
