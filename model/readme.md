## 模型说明

    本项目未使用gorm库 https://github.com/jinzhu/gorm，

## 惯例和约定

    比如我要新建一个User模型，在model中新建 UserModel文件夹【文件夹名大写驼峰】，
    
    然后在UserModel下新建 UserModel文件【文件名大写驼峰】，
    
    可能有一些字段我们并不想返回给前端，或者需要对结果集进行在包装修改，
    
    所以可能还需要一个Source文件来处理这些事情


## 迁移文件

    migrate目录下的migrate.go是数据库迁移代码