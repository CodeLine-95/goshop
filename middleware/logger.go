package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goshop/utils/logger"
	"time"
)

//gin自定义日志中间件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		//执行时间
		latencyTime := endTime.Sub(startTime).Seconds()
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求IP
		ClientIp := c.ClientIP()
		//用户标识
		UserAgent := c.Request.UserAgent()

		//日志格式
		logger.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    ClientIp,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"user_agent":   UserAgent,
		}).Info()

	}
}
