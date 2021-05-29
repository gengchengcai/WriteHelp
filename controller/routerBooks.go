package controller

import "github.com/gin-gonic/gin"

func (r *Routers) CreateBookRouters(router *gin.RouterGroup) {
	//初始页面显示部分，获取图书信息---图书大纲，书的名称，已完成章节数
	router.GET("v1/books")
	//进入到一本书的编辑页面，获取章节数目，章节标题，书籍大纲，
	router.GET("v1/books/:id")
	//选定章节 获取具体的章节内容和章节大纲
	router.GET("v1/books/:id/chapter/:id")
	//获取某本书中的任务设定
	router.GET("v1/books/:id/plotsetting")
	//获取任务设定中的具体某一个
	router.GET("v1/books/:id/plotsetting/:id")
	//获取某本书里的任务设定
	router.GET("v1/books/:id/rolesetting")
	//获取任务设定中的一个
	router.GET("v1/books/:id/rolesetting/:id")
	//

	//创建一本新书，只需要填写书名和当时出现的大致想法
	router.POST("v1/books")
	//创建一个章节 要章节数和章节名称
	router.POST("v1/books")

	//修改章节的内容，标题，内容，大纲
	router.PUT("v1/books/:id/chapter/:id")
	//修改书籍设定内容
	router.PUT("v1/book/:id")

	//删除具体章节
	router.DELETE("v1/books/:id/chapter/:id")
	//删除掉一本书
	router.DELETE("v1/books/:id")

}
