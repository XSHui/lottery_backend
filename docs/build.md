# 编译/发布说明

## 一、编译
采用docker编译
```
sh -x build.sh $version
```

## 二、发布

1. Mysql 使用UCloud-Mysql服务

2. Redis 使用UCloud-Redis服务

3. lottery-backend 采用docker方式部署在UCloud快杰云主机上

4. lottery-console [暂时未实现]
