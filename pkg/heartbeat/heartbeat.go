package heartbeat

import (
	"context"
	"fmt"
	"time"
)

func Heartbeat(ctx context.Context) {
	tick := time.Tick(5 * time.Second)

	for {
		select {
		case <-tick:
		case <-ctx.Done():
			return
		}
		fmt.Println("heartbeat running ...")
	}
}

// Need to work
func Heartbeat2() {
	fmt.Println("heartbeat...Not Actived yet")
	//tick := time.Tick(20 * time.Second)

	// for {
	// 	// <-tick
	// 	fmt.Println("heartbeat2...")
	// }
}

// Need to work
func Cron() {

	// for {
	// 	time.Sleep(time.Minute)
	// 	fmt.Println("cron...")
	// }
}
