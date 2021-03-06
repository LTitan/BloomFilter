package router

import (
	"fmt"
	"net/http"
	"strings"

	_ "github.com/LTitan/BloomFilter/docs"
	rh "github.com/LTitan/BloomFilter/internal/router/handler"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/LTitan/BloomFilter/pkg/signal"
	"github.com/LTitan/BloomFilter/pkg/sql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var fe rh.FE
var slave rh.Slave

// InitRouter .
func InitRouter(port string) {
	go signal.ExitBeautiful(func() {
		err := sql.DefaultDB.Close()
		logs.Logger.Warnf("db will close, error %v", err)
	})
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(Cors())
	r := router.Group("/api/v1")
	{
		r.GET("/host/info", fe.QueryCPUMemory)
		r.GET("/host/pagination", fe.QueryApplyPagination)
		r.POST("/host/applyinfo/update", fe.UpdateApplyRecord)
		r.DELETE("/host/applyinfo", fe.DeleteApplyRecord)
		r.GET("/host", fe.GetAliveHosts)
		r.GET("/host/single", fe.GetSingleAddressInfo)
		r.GET("/host/register", fe.RegisterDistribution)
		r.GET("/host/register/memory", fe.GetRegisterMemoryInfo)

		r.POST("/file/upload", upload)
		r.GET("/file/upload", slave.ReadUploadFile)

		r.POST("/user", fe.CreateUser)
		r.POST("/user/authorization", fe.QueryHasUser)
		r.POST("/bloomfilter/apply", slave.ApplyMemory)
		r.GET("/bloomfilter/query", slave.QueryValue)
		r.POST("/bloomfilter/query", slave.QueryMany)
		r.POST("/bloomfilter/add", slave.AddValues)
		r.DELETE("/bloomfilter/del/:uuid", slave.DeletKey)
		r.PUT("/bloomfilter/:address", slave.BackupSlave)
	}
	router.Run(port)
}

// Cors .
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(500, err.Error())
		return
	}
	files := form.File["file"]
	for _, file := range files {
		var path = file.Filename
		err = ctx.SaveUploadedFile(file, "/tmp/"+path)
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
	}
	ctx.String(200, "yes")
}
