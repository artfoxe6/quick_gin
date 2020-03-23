##模型说明

    本项目未使用gorm库
    https://github.com/jinzhu/gorm

##惯例和约定

    比如我要新建一个User模型，在model中新建 UserModel文件夹【文件夹名大写驼峰】，
    然后在UserModel下新建 UserModel文件【文件名大写驼峰】，
    可能有一些字段我们并不想返回给前端，所以进行资源化很有必要，通常在还需要
    新建一个Source文件


##迁移文件

    migrate目录下的migrate.go是数据库迁移代码