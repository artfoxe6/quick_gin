# quick_gin
基于Gin整理的一个轻量级接口模板,适合简单的后端接口项目

#### 基于Gin作了什么?
规划了一套代码组织架构 <br />
对gin.Context扩展 <br />
已配置好MySQL,Redis,MongoDB,JWT开箱即用 <br />
使用Gorm <br />
使用跨域中间件 <br />
平滑重启 <br />
完整的代码实例和说明文档 <br />

#### 运行条件：
+ Golang >= 1.12
+ Go Module
+ MySQL/MariaDB

#### 安装运行
git clone https://github.com/artfoxe6/quick_gin.git <br />
cd quick_gin <br />
\#初始化项目 <br />
./init.sh my_blog && cd ../my_blog <br />
\#修改config.ini <br />
go run main.go <br />

##### 项目结构说明
> model:模型定义,目录和文件命名方式请参考案例 <br />
> route:路由放在api下面 <br />
> service:业务层层

由于本项目定位小型的接口系统,model层除了定义模型关系还要负责数据的存取
service层只负责业务逻辑处理


    

    