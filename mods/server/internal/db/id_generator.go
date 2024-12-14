package db

import (
	"github.com/sony/sonyflake"
	"time"
)

func snowFlakeGenerator() (*sonyflake.Sonyflake, error) {
	return sonyflake.New(sonyflake.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			return 1, nil
		},
	})
}
