package main

import (
	"fmt"
	"time"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(starttime string, machineID int64) (err error) {
	// 时间 机器id
	// 从那一年开始
	var st time.Time
	if st, err = time.Parse("2006-01-02", starttime);err !=nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return nil
}

func GenID() int64 {
	return node.Generate().Int64()
}

func main()  {
	if err := Init("2021-01-01", 1);err != nil {
		fmt.Printf("init snowflake fail: %v\n", err)
	}
	id := GenID()
	fmt.Println(id)
}