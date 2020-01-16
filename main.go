/**
* @Author :Administrator
* @Description:
* @File:main
* @Version:1.0.0
* @Date:
 */

package main

import (
	"files-system/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	initRouters()
}

func main()  {
}

/*
初始化配置文件
 */
func initConfig()  {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".")
}

/*
初始化路由
 */
func initRouters() {
router := gin.Default()
	routers.InitWeb(router)
}

/*
初始化文件日志
 */
func  initLogFile()  {
	
}

/*
初始化levelDB
 */
func initLevelDB()  {
	
}
/*
初始化redis
 */
func initRedis()  {
	
}
/*
定时任务
 */

func  initTask()  {
	
}