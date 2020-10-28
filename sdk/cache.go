/*Package sdk ...
*
	 实验配置本地缓存
*
*/
package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
	// "github.com/go-abtest/sdk/model"
)

var (

	// cacheTestABTZonesCache for demo test, ABT config cache
	// cacheTestABTZonesCache map[string][]*Zone = make(map[string][]*Zone)
	cacheTestABTZonesCache sync.Map

	// CTX , CatcheCancer catch 同步任务上下文，单例
	cacheCTX context.Context

	// Cancer  ...
	cacheCancer context.CancelFunc

	// Mu 互斥锁，单例
	cacheMu sync.Mutex

	// Running 标记是否已经运行 cache 任务
	cacheRunning bool
)

// runCacheSyncDBTask ..
func runCacheSyncDBTask(ctx context.Context, projects []string, duration time.Duration) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n>>>>Thread runCacheSyncDBTask Info<<<<Info", "runCacheSyncDBTask thread was closed by ctx!")
			return
		case <-time.Tick(duration):
			fmt.Println("\n>>>>Thread runCacheSyncDBTask Info<<<<Info", "runCacheSyncDBTask doSyncDB begin")
			doSyncDB(projects)
		}
	}
}

// doSyncDB ..
func doSyncDB(projects []string) {
	for _, project := range projects {
		content, err := getRemoteConfig(project, time.Second*60)
		if err != nil {
			log.Fatal("doSyncDB call ioutil.ReadFile failed, error:", err)
		}
		zones := make([]*Zone, 0)
		err = json.Unmarshal(content, &zones)
		if err != nil {
			log.Fatal("doSyncDB call json.Unmarshal failed, error:", err)
		}
		cacheZones, ok := cacheTestABTZonesCache.Load(project)
		if ok {
			cacheTestABTZonesCache.Store(project, append(cacheZones.([]*Zone), zones...))
		} else {
			cacheTestABTZonesCache.Store(project, zones)
		}
	}
	fmt.Println("Once doSyncDB syncing task done!")
}
