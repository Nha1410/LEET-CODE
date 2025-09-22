package main

import (
	"fmt"
	"sync"
	"time"
)

const playerCount = 10

var (
	currentHolder int        = 0 // Người đang giữ bóng
	mutex         sync.Mutex     // Bảo vệ truy cập currentHolder
	wg            sync.WaitGroup
)

func main1() {
	for i := range playerCount {
		wg.Add(1)
		go player(i)
	}

	wg.Wait()
}

func player(id int) {
	defer wg.Done()

	for {
		mutex.Lock()
		isHolder := (id == currentHolder)
		mutex.Unlock()

		if isHolder {
			// Nếu là người giữ bóng
			fmt.Printf("%s - Player %d: Pass Ball\n", time.Now().Format("15:04:05"), id)

			time.Sleep(10 * time.Second)

			// Chuyền bóng cho người tiếp theo
			mutex.Lock()
			currentHolder = (currentHolder + 1) % playerCount
			mutex.Unlock()
		} else {
			// Nếu không phải người giữ bóng
			fmt.Printf("%s - Player %d: pass for me\n", time.Now().Format("15:04:05"), id)

			time.Sleep(1 * time.Second)
		}
	}
}
