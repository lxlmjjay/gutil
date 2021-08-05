package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

const TokenError = "token 错误"
const ParaError = "参数错误"
const OpSuccess = "操作成功"
const TokenNone = "请求未携带token，无权限访问"
const OpFail = "操作失败-"

func RspAny(key string, data interface{}) string {
	m := make(map[string]interface{})
	m[key] = data
	return Json(m)
}

func RspByte(key string, data interface{}) []byte {
	m := make(map[string]interface{})
	m[key] = data
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	return marshal
}


func RspShopError(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status": "error",
		"msg":    msg,
	})
	return
}

func RspShopSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    msg,
	})
	return
}

func RspTokenError(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		//"status": "fail",
		"status": "token_fail",
		"msg":    msg,
	})
	return
}

func RspError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"status": "fail",
		"msg":    msg,
	})
	return
}

func RspSuccessSimple(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    "操作成功",
	})
	return
}

func RspSuccess(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    msg,
	})
	return
}

func RspSuccessWithData(c *gin.Context, data interface{}, msg string) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    msg,
		"data":   data,
	})
	return
}

func RspBool(c *gin.Context, msg bool) {
	c.JSON(200, gin.H{
		"status": "success",
		"data":   msg,
	})
	return
}

//token过期
func RspTokenExpire(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status": "tokenExpire",
		"token":  msg,
	})
	return
}

func RspDataPage(c *gin.Context, data interface{}, total int64) {
	c.JSON(200, gin.H{
		"status": "success",
		"total":  total,
		"data":   data,
	})
	return
}

func RspToken(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"token":  data,
	})
	return
}

func RspData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
	return
}

func RspDataWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    msg,
		"data":   data,
	})
	return
}
