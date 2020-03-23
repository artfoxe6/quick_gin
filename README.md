# quick_gin

##介绍
    quick_gin是基于gin框架的脚手架，
    根据自己经验整理的一套代码组织架构
    开箱即用，快速开始业务编写，
    封装了常用的辅助方法，
    目的就是为了提高开发效率

##组织架构
    Route->Service->Cache->Model
    路由层：路由规则和中间件
    Service：服务层 验证请求 逻辑处理 返回数据
    Cache： 缓存层 简单的查询缓存
    Model：数据层 数据库CURD
##目录说明
    
    /config/
        配置初始化
    
    /middleware/
        路由中间件
    
    /model/
        Model数据层
        
    /route/
        路由规则
        
    /service/
        Service逻辑层
        
    /util/
        存放工具类
        
    /config.ini
        真个项目的配置文件，采用ini配置方式
        
    /main.go
        入口文件，
        负责整个服务的启动，
        已开启平滑重启
        
##相关目录下还有更详细的惯例和说明

    