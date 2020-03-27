#### 路由说明

    在route下面新建路由目录，比如api，在api目录下新建路由，路由按功能拆分成不同文件，
    比如文章板块article.go，用户板块user.go
#### ApiDoc
    如果需要生成接口文档，请在路由中包含apidoc描述，案例：article.go中，
    安装apidoc： https://apidocjs.com/
    
    在项目目录运行 apidoc -i route/ ,默认接口文档生成在doc目录，
    
#### 惯例和约定

    路由目录和文件名小写驼峰

    路由文件必须包含一个加载函数，命名规则 LoadXXXRoute，例如：
    
    func LoadArticleRoute(r *gin.Engine) {
    	g := r.Group("/article")
    
    	g.POST("/add", func(c *gin.Context) {
    		ArticleService.Add(request.New(c))
    	})
    }

    
#### 加载路由
    
    在route目录下的route.go中调用所有的路由加载函数