package controller

import "github.com/gin-gonic/gin"

type Routers struct {
}

func (r *Routers) CreateRouters(engine *gin.RouterGroup) {
	engine.POST("register", r.Register)
	engine.POST("login", r.Login)
}

func (r *Routers) CreateBookRouters(router *gin.RouterGroup) {
	//初始页面显示部分，获取图书信息---图书大纲，书的名称，已完成章节数
	router.GET("books/:id", r.GetBooksInfo)
	//进入到一本书的编辑页面，获取章节数目，章节标题，书籍大纲，
	router.GET("books/:id")
	//选定章节 获取具体的章节内容和章节大纲
	router.GET("books/:id/chapter/:id")
	//获取某本书中的剧情设定
	router.GET("books/:id/plotsetting")
	//获取剧情设定中的具体某一个
	router.GET("books/:id/plotsetting/:id")
	//获取某本书里的人物设定
	router.GET("books/:id/rolesetting")
	//获取人物设定中的一个
	router.GET("books/:id/rolesetting/:id")
	//获取灵感
	router.GET("books/:id/inspirations")
	//获取灵感中的某一个
	router.GET("books/:id/inspirations/:id")

	//创建一本新书，只需要填写书名和当时出现的大致想法
	router.POST("books")
	//创建一个章节 要章节数和章节名称
	router.POST("books/chapter")
	//创建一个灵感
	router.POST("books/inspiration")
	//创建一个剧情设定
	router.POST("books/plotsetting")
	//创建一个人物设定
	router.POST("books/rolesetting")

	//修改章节的内容，标题，内容，大纲
	router.PUT("books/:id/chapter/:id")
	//修改书籍设定内容
	router.PUT("books/:id")
	//修改书籍的灵感
	router.PUT("books/:id/inspiration/:id")
	//修改书籍的人物设定
	router.PUT("books/:id/rolesetting/:id")
	//修改书籍的剧情设定
	router.PUT("books/:id/plotsetting/:id")

	//删除具体章节
	router.DELETE("books/:id/chapter/:id")
	//删除掉一本书
	router.DELETE("books/:id")
	router.DELETE("books/:id/inspiration/:id")
	router.DELETE("books/:id/rolesetting/:id")
	router.DELETE("books/:id/plotsetting/:id")

}
