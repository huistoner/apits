package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main ()  {
	fmt.Println("starting")
	i := 0
	spec :="* */1 * * * *"
	c := cron.New(cron.WithSeconds())
	c.AddFunc(spec, func() {
		i++
		fmt.Printf("定时器 %d 次执行。\n",i)
	})
	c.Start()

	select {}
}
