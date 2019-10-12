# smyapp
 #### introduction
 this is a simple api sever project without thirdpart web  and orm framework which is base on go language,
 i do not use much character of go such as channel 、goroutine e.g.  but they do used in native go libs. 
 #### Project structure
 ##### 1.controller
 > structs with methods like func (r http.ResponseWriter,w *http.Request), which are registed  in fun init()
 ##### 2.service
 > deal with business and return operateMessage
 ##### 3.dao
 > connect with db,and exec sqls 
e.g

 #### frameworks or tools used
 1. github.com/bwmarrin/snowflake   distribute id generator
 2. github.com/satori/go.uuid        uuid for go 
 3. github.com/robfig/cron          time task for go(not used yet)
 4. go.uber.org/zap                 effective log frame (i think i do not use it well)
 5. github.com/go-sql-driver/mysql  driver for mysql implements from go database/sql 
 
