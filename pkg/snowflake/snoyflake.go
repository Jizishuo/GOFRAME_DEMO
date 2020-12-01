package main

import (
	"github.com/sony/sonyflake"
)

var (
	sonyFlake *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需要传入当前机器的id
