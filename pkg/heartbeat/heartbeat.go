package heartbeat

import (
	"context"
	"time"

	"jayantapaul-18/uptime/util/logger"
)

func Heartbeat(ctx context.Context) {
	tick := time.Tick(5 * time.Second)
	logger := logger.New() // Create New logger
	for {
		select {
		case <-tick:
		case <-ctx.Done():
			return
		}
		logger.Info().Msgf("heartbeat running ...")
	}
}

// Need to work
func Heartbeat2() {
	// logger.Info().Msgf("heartbeat...Not Actived yet")
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
