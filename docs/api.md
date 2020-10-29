# HTTP API LIST
1. 用户注册
```
req:
curl -v -X POST -d '{"PhoneNumber": 12345, "VerifyCode": 1,  "Action": "LogIn"}' http://10.23.221.43:8888

res:
{"Action":"LogInResponse","RetCode":0,"Message":"","UserId":"de44d5af-f447-4d68-bc54-85b77cc2e158"}
```

2. 用户已存在
```
req:
curl -v -X POST -d '{"PhoneNumber": 12345, "Action": "UserExist"}' http://10.23.221.43:8888

res:
{"Action":"UserExistResponse","RetCode":0,"Message":"","Exist":true}
```

3. 提交文章
```
req:
curl -v -X POST -d '{"UserId": "de44d5af-f447-4d68-bc54-85b77cc2e158", "Text": "good goot study, day day up!!!!",  "Action": "SubmitArticle"}' http://10.23.221.43:8888

res:
{"Action":"SubmitArticleResponse","RetCode":0,"Message":""}
```

4. 文章列表
```
req:
curl -v -X POST -d '{"Offset": 0, "Limit": 10, "Action": "ListArticle"}' http://10.23.221.43:8888

res:
{"Action":"ListArticleResponse","RetCode":0,"Message":"","DataSet":[{"id":"5990a2cf-2343-402a-8710-e28247c90807","user_id":"de44d5af-f447-4d68-bc54-85b77cc2e158","article":"good goot study, day day up!!!!"}]}
```

5. 抽奖
```
curl -v -X POST -d '{"PhoneNumber": 10, "Action": "Lottery"}' http://10.23.221.43:8888

{"Action":"LotteryResponse","RetCode":0,"Message":"","Win":true,"PrizeName":"stickers","PrizeId":"0E9BBC05-5A16-400F-8436-966F1FD3BD7A"}

{"Action":"LotteryResponse","RetCode":0,"Message":"once a day","Win":false,"PrizeName":"","PrizeId":""}
```

6. 中奖记录
```
curl -v -X POST -d '{"Offset": 0, "Limit": 10, "Action": "ListRecord"}' http://10.23.221.43:8888

{"Action":"ListRecordResponse","RetCode":0,"Message":"","DataSet":[{"id":"9104c387-d1f9-479b-9af9-088666d3eb3c","user_id":"ab750cfe-e1ec-410d-bf93-52632c7a61d7","prize_id":"0E9BBC05-5A16-400F-8436-966F1FD3BD7A","create_time":1603950732,"modify_time":1603950732,"delete_time":0}]}
```

// TODO:
7. 注册活动:
8. 注册奖品:
9. 活动场次List
10. 奖品List:



