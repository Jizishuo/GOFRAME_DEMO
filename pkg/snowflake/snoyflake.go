package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需要传入当前机器的id
func Initsf(starttime string, machineID uint16) (err error) {
	sonyMachineID = machineID
	var st time.Time
	if st, err = time.Parse("2006-01-02", starttime);err!=nil {
		return
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return nil
}

// genid 生成id
func GenIDsn() (id uint16, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flaske not init")
		return
	}
	id ,err = sonyFlake.NextID()
	return
}

func main()  {
	if err := Initsf("2021-01-01", 1); err != nil {
		fmt.Printf("init faild err: %v\n", err)
		return
	}
	id, _ := GenIDsn()
	fmt.Println(id)
}