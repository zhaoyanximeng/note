package main

import "github.com/gin-gonic/gin"

// 装饰器模式：允许向一个现有的对象添加新功能，又不改变其结构

//func CheckLogin(f http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if r.URL.Query().Get("token") == "" {
//			w.Write([]byte("token error"))
//		}else{
//			f(w,r)
//		}
//	}
//}
//
//func index(w http.ResponseWriter,r *http.Request) {
//	w.Write([]byte("index"))
//}
//
//func main() {
//	http.HandleFunc("/",CheckLogin(index))
//	http.ListenAndServe(":8080",nil)
//}

func CheckLogin(f gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Query().Get("token") == "" {
			c.JSON(400,gin.H{
				"msg":"token error",
			})
		} else {
			f(c)
		}
	}
}

func index(c *gin.Context) {
	c.JSON(200,gin.H{
		"msg":"index",
	})
}

func main() {
	r := gin.New()
	r.GET("/",CheckLogin(index))
	r.Run(":8080")
}