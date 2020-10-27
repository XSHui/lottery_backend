# 代码结构说明
```
.
   ├── Dockerfile           // docker file, build code and docker docker image
   ├── README.md      
   ├── build.sh             // build shell
   ├── docs          
   │   ├── api.md           // api doc
   │   ├── build.md         // how to build and deploy
   │   ├── code.md          // structure of this prj
   │   ├── design.md        // design of this prj
   │   └── images           // image [todo]
   ├── go.mod
   ├── go.sum
   ├── lottery              // bin
   ├── lottery_backend.tar  // docker image
   ├── src
   │   ├── config
   │   │   ├── config.go                 // server config
   │   │   ├── config.toml.example_1     // config toml [todo]
   │   │   ├── config_test.go            // UT [todo]
   │   │   └── cst.go                    // const
   │   ├── main.go                     
   │   ├── model            // api model
   │   │   ├── article.go  
   │   │   ├── lottery.go
   │   │   ├── record.go
   │   │   └── user.go
   │   ├── redis            
   │   │   └── redis.go     // distributed lock implement by redis
   │   ├── server           // api server, UT [todo]
   │   │   ├── article.go
   │   │   ├── inputtransform.go
   │   │   ├── lottery.go
   │   │   ├── prize.go
   │   │   ├── record.go
   │   │   ├── router.go
   │   │   ├── server.go
   │   │   ├── session.go       // session related
   │   │   ├── user.go      
   │   │   └── utils.go         // some utils
   │   ├── xlog
   │   │   ├── xlog.go          // log
   │   │   └── xlog_test.go     // UT [todo]
   │   └── xorm                 // xorm
   │       ├── article.go 
   │       ├── model            // xorm model
   │       │   ├── activity.go
   │       │   ├── article.go
   │       │   ├── permmition.go
   │       │   ├── prize.go
   │       │   ├── record.go
   │       │   └── user.go
   │       ├── permission.go
   │       ├── prize.go
   │       ├── record.go
   │       ├── user.go
   │       └── xorm.go
   └── start.sh           // start shell of docker container
   ```