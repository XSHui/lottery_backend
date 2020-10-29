# 线上征文抽奖活动设计文档

## 一、设计目标

### 1.1 活动背景

为了庆祝业务周年，需要举行一个线上征文活动


### 1.2 活动内容

* 用户需要在页面用户填写手机号码报名（需要验证码校验，一个手机号只能报名一次），并上传一段500字以内的文本。
* 成功的用户可以在页面参与抽奖，每天可以抽一次
* 需要提取所有报名用户的手机号码和对应的征文内容
* 需要导出所有获奖记录


### 1.3 奖池内容

```
    物品    数量    规则
    手机    5部     每天限制产出1部，概率1%
    电话卡  100张   每个用户活动期间内最多获得2张，概率5%
    贴纸    不限量  概率94%
```

### 1.4 开发信息

* 发送验证码无需实现，定义空函数即可
* 开发语言不限但建议使用go/php
* 存储可用redis或mysql
* 前端使用javascript进行请求并根据返回结果进行渲染
* 代码建议提交到git.code.tencent.com或其他git仓库并提供可读权限


## 二、整体框架 [TODO]


## 三、前端[TODO]


## 四、后端

### 2.1 后端业务 [TODO]


### 2.2 DB设计(mysql)

**瓶颈/难点**:

- 后端DB压力
 
  - 请求量大

  - 热点问题，抽奖活动会集中在奖品表的部分数据上，db设计不当会产生热点问题进而成为系统瓶颈


**解决方案**:

1. API层限流 / 防刷 [TODO]

2. 业务侧加锁(redis分布式锁)

3. 热点问题，在真实业务场景可以考虑一些策略将数据进行打散


> table: user
```
  3 type User struct {
  4     Id       string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
  5     PhoneNumber uint64 `json:"phone_number" xorm:"not null unique default 0 INT(11)"`
  6     // TODO: State
  7     CreateTime  int    `json:"create_time" xorm:"INT(10)"`
  8     ModifyTime  int    `json:"modify_time" xorm:"INT(10)"`
  9     DeleteTime  int    `json:"delete_time" xorm:"INT(10)"`
 10 }
```

> table: activity.go 
```
  3 // TODO: multi lottery
  4 type Activity struct {
  5     // Id
  6     // Name
  7     // StartTime
  8     // EndTime
  9 }
```

> table: article
```
  3 type Article struct {
  4     Id     string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
  5     UserId string `json:"user_id" xorm:"not null default '' VARCHAR(128)"`
  6     // TODO: ActivityId
  7     Text   string `json:"article" xorm:"not null default '' VARCHAR(1024)"`
  8 }
```

> table: permmition
```
 3 type Permission struct {
  4     UserId       string `json:"user_id" xorm:"not null unique default '' VARCHAR(128)"`
  5     // TODO: ActivityId
  6     Permitted    int `json:"permitted" xorm:"default 0 TINYINT(1)"`
  7 }
```

> table: prize
```
  3 type Prize struct {
  4     Id string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
  5     // TODO: ActivityId
  6     Award     uint64  `json:"award" xorm:"INT(10)"`
  7     Name      string  `json:"prize name" xorm:"VARCHAR(256)"`
  8     Odds      float32 `json:"odds" xorm:"FLOAT"`
  9     Total     int     `json:"total" xorm:"INT(10)"`
 10     Left      int     `json:"left" xorm:"INT(10)"`
 11     Unlimited int     `json:"unlimited" xorm:"default 0 TINYINT(1)"`
 12     // TODO: State
 13     CreateTime int `json:"create_time" xorm:"INT(10)"`
 14     ModifyTime int `json:"modify_time" xorm:"INT(10)"`
 15     DeleteTime int `json:"delete_time" xorm:"INT(10)"`
 16 }
```

> table: record
```
  3 type Record struct {
  4     Id     string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
  5     UserId string `json:"user_id" xorm:"not null unique(permission) default '' VARCHAR(128)"`
  6     // TODO: ActivityId
  7     PrizeId string `json:"user_id" xorm:"not null unique(permission) default '' VARCHAR(128)"`
  8     // TODO: State
  9     CreateTime int `json:"create_time" xorm:"not null unique(permission) default 0 INT(10)"`
 10     ModifyTime int `json:"modify_time" xorm:"INT(10)"`
 11     DeleteTime  int    `json:"delete_time" xorm:"INT(10)"`
 12 }
```
