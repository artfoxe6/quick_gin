![image](https://github.com/codeAB/store/blob/master/images/quick_gin_1.gif?raw=true)

## 安装使用
###### Linux, Mac OS下安装使用【推荐】

    wget -qO create_gin_app https://github.com/codeAB/quick_gin/releases/download/v0.1/create_gin_app
    chmod +x create_gin_app
    (可选) mv create_gin_app /usr/bin/
    create_gin_app blog

###### Window下安装使用
    git clone https://github.com/codeAB/quick_gin.git blog
    删除 .git目录
    更改项目中的 quick_gin关键字为你的项目名

## 运行本项目需确保：
    golang版本要求大于等于1.12
    为了依赖能够顺利下载，推荐设置环境变量 GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
    确保MySQL/MariaDB已安装并启动，正确配置config.ini中的database参数，
    如果需要运行案例，还需要创建 test 数据库，运行程序将自动创建user和article两张表
    确保Redis已安装并启动，正确配置config中redis参数
    
## 基于gin做的一些封装和整理    
    规划了一套代码组织架构
    对gin.Context进行了扩展，增加很多辅助方法
    集成跨域中间件
    集成请求缓存中间件，
    使用GORM
    封装Jwt
    已经配置好MySQL，Redis 开箱即用
    Sentry日志
    env全局配置
    自带平滑重启
    完整的代码实例和说明文档
    命令行创建项目 create_gin_app blog 【命令行创建项目只适用于Linux】
    

    
## 目录说明
    /config/        配置
    /middleware/    中间件
    /model/         模型
    /route/         路由规则
    /service/       业务逻辑
    /util/          工具类
    /config.ini     整个项目的配置文件，采用ini配置方式
    /main.go        负责整个服务的启动，
**请注意具体目录下更详细的惯例和说明readme**
    
## 代码组织架构
    Route->Cache->Service->Model
    
    Route：      路由规则和中间件
    Cache：      对需要的请求做缓存
    Service：    验证请求，逻辑处理
    Model：      数据库CURD，数据资源化

## 如何开始上手编写业务
    
    检查config.ini     修改配置参数
    新建route文件       路由映射到service的方法
    新建service文件     用到的数据由model层提供
    新建model文件       负责数据库CURD
    
    上述步骤可以参考自带的实例
    

    