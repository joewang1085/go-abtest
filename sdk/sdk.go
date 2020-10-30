/*Package sdk ...
*
	 对外提供 sdk
*
*/
package sdk

import (
	"context"
	"fmt"
	"strings"
	"time"
)

/*
*
	流量分流的方式：
	1. userID + layerID
	2. cookie(deviceID等) + layerID
	3. userID + Date + layerID
	4. cookie(deviceID等) + Date + layerID
*
*/

// GetABTZone 通过匹配，获取当前选择 Zone. projectID 用于指定实验项目，globalID, layerID and date 用于计算hash
func GetABTZone(projectID, hashkey, layerID string) *Zone {
	// 匹配库的Zone
	// fmt.Println(fmt.Sprintf("Debug for hash: project:%s, userID:%s, layerID:%s, hash:%v", projectID, userID, layerID, Hash(projectID, userID, layerID)))
	return matchZone(getZonesByProjectIDandLayerID(projectID, layerID), projectID, hashkey, layerID)
}

// PushLabOutPut 数据采点，上报ABT实验输出
func PushLabOutPut(data *LabOutput) {
	path := ""
	for _, p := range data.LabPath {
		if path == "" {
			path = p
		} else {
			path = path + "->" + p
		}
	}
	tag := ""
	for _, p := range data.TagPath {
		if tag == "" {
			tag = p
		} else {
			tag = tag + "->" + p
		}
	}

	outputPush(fmt.Sprintf("Project:%s,UserID:%s,Path:%s,TagPath:%s", data.ProjectID, data.UserID, path, tag), data)
}

// GetLabOutput 获取实验数据
func GetLabOutput(projectID, path, tag string) []*LabOutput {
	tg := tag
	if tg == "" {
		tg = "空"
	}
	if path == "" {
		return nil
	}
	path = "Path:" + path
	fmt.Println(fmt.Sprintf("查询条件：Project=%s,paht=%s,tag=%s", projectID, path, tg))
	data := make([]*LabOutput, 0)
	outputRange(func(k, v interface{}) bool {
		if strings.Contains(k.(string), projectID) &&
			strings.Contains(k.(string), tag) &&
			strings.Contains(k.(string), path) {
			data = append(data, v.(*LabOutput))
		}

		return true
	})

	return data
}

// SetCacheSyncDBFrequency 设置本地缓存同步DB的周期
func SetCacheSyncDBFrequency(projects []string, duration time.Duration) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	// 先关闭任务
	if cacheRunning {
		cacheCancer()
	}

	// 不能共用一个 ctx
	cacheCTX, cacheCancer = context.WithCancel(context.Background())

	cacheRunning = true
	doSyncDB(projects)
	go runCacheSyncDBTask(cacheCTX, projects, duration)
}

// CheckUserGroup 检查用户是否是目标流量
func CheckUserGroup(groups []string, cur string) bool {
	for _, group := range groups {
		if cur == group {
			return true
		}
	}

	return false
}

func matchZone(zones []*Zone, projectID, hashkey, layerID string) *Zone {
	// 根据userID , layerID 匹配对应的域
	// ...
	if len(zones) == 0 {
		return &Zone{}
	}
	hash := hash(projectID, hashkey, layerID, zones[0].TotalWeight)

	// check current zone
	for _, zone := range zones {
		if zone.Max >= hash && zone.Min <= hash {
			// check parent zone, make sure user comes from parent zones
			isFromParent := false
			for _, parent := range zone.ParentZones {
				if parent.ZoneID == matchZone(getZonesByProjectIDandLayerID(parent.ProjectID, parent.LayerID), parent.ProjectID, hashkey, parent.LayerID).ZoneID {
					isFromParent = true
					break
				}
			}

			if len(zone.ParentZones) == 0 || isFromParent {
				return zone
			}
		}
	}
	return &Zone{}
}

func getZonesByProjectIDandLayerID(projectID, layerID string) []*Zone {
	cacheZones, ok := cacheTestABTZonesCache.Load(projectID)
	if !ok {
		return make([]*Zone, 0)
	}

	// 从缓存中获取
	zones := make([]*Zone, 0)
	for _, zone := range cacheZones.([]*Zone) {
		if layerID == zone.LayerID {
			zones = append(zones, zone)
		}
	}

	return zones
}
