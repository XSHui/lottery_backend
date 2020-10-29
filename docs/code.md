# 代码结构说明
```
.
├── Dockerfile      // docker build server and image
├── README.md
├── build.sh        
├── docs
│   ├── api.md      // api doc
│   ├── build.md    // build and deploy doc
│   ├── code.md     // code doc
│   ├── design.md   // design doc
│   ├── images
│   └── sql.md      // init peize table sql
├── go.mod
├── go.sum
├── lottery           
├── lottery_backend.tar
├── src
│   ├── access  
│   │   ├── api                    // api and lottery
│   │   │   ├── article.go         // article related api
│   │   │   ├── inputtransform.go  
│   │   │   ├── lottery.go         // lottery related api
│   │   │   ├── prize.go           // peize related api
│   │   │   ├── record.go          // record related api
│   │   │   ├── router.go 
│   │   │   ├── server.go
│   │   │   └── user.go            // user related api
│   │   └── model                  // api model
│   │       ├── article.go
│   │       ├── lottery.go 
│   │       ├── record.go
│   │       └── user.go
│   ├── config                // config related
│   │   ├── config.go
│   │   ├── config.toml.example_1
│   │   ├── config_test.go
│   │   └── cst.go           // const
│   ├── main.go
│   ├── redis
│   │   └── redis.go         // redis lock
│   ├── utils
│   │   ├── session.go
│   │   └── utils.go
│   ├── xlog
│   │   ├── xlog.go
│   │   └── xlog_test.go
│   ├── xml                 // local console
│   │   ├── axios.min.js
│   │   └── lottery.html
│   └── xorm
│       ├── article.go
│       ├── model    // xorm  model
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
├── start.sh  // docker container start.sh
├── test 
│   └── test_lottery.go  // a simple test for backend api
└── test_lottery
```
