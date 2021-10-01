package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {


	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("good")
	})

	c.Start()
	t1 := time.NewTimer(time.Second*10)
	//for{
	//	select {
	//	case <- t1.C:
	//		t1.Reset(time.Second*10)
	//	}
	//}


	go func() {
		fmt.Println("当前时间为:", time.Now())

		t := <-t1.C

		fmt.Println("当前时间为:", t)
	}()

	for{
		time.Sleep(1*time.Second)
	}


	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()


	// todo gjw 热更新
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
