package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-abtest/sdk"
)

var (
	// Lab 名称
	Lab string

	// Users 用户数量
	Users int
)

func init() {
	flag.StringVar(&Lab, "lab", "Display", "Lab name")
	flag.IntVar(&Users, "users", 100, "参与测试的用户总数")
	flag.StringVar(&sdk.Address, "addr", "127.0.0.1", "AB Test Server Address")
	flag.StringVar(&sdk.Port, "port", "9527", "AB Test Server Port")
}

func main() {
	flag.Parse()

	// SetCacheSyncDBFrequency
	fmt.Println("设置缓存同步数据库周期")
	sdk.SetCacheSyncDBFrequency([]string{"search", "Order", "Display"}, time.Second*60)

	// 开启实验统计线程
	go func() {
		// 实验 数据统计
		// sdk.LabOutputCache.Range(func(k, v interface{}) bool {
		// 	fmt.Println("Debug for LabOutputCache: ", k, v)
		// 	return true
		// })
		fmt.Println("\n开启实验数据统计 线程...")
		for {
			select {
			case <-time.Tick(time.Second * 10):
				switch Lab {
				case "search":
					printLabData(sdk.GetLabOutput("search", "A", ""))
					printLabData(sdk.GetLabOutput("search", "B", ""))
					// printLabData(sdk.GetLabOutput("search", "defualt", ""))
				case "Order":
					printLabData(sdk.GetLabOutput("Order", "A1->A2-1", ""))
					printLabData(sdk.GetLabOutput("Order", "A1->A2-2", ""))
					printLabData(sdk.GetLabOutput("Order", "B2-1->B2-1->B3", ""))
					printLabData(sdk.GetLabOutput("Order", "B2-1->B2-2->B3", ""))
					printLabData(sdk.GetLabOutput("Order", "B2-1->B2-1->B3", ""))
					printLabData(sdk.GetLabOutput("Order", "B2-1->B2-2->B3", ""))
				case "Display":
					fmt.Println("按 labPath 查询...")
					printLabData(sdk.GetLabOutput("Display", "A", ""))
					printLabData(sdk.GetLabOutput("Display", "B", ""))
					// fmt.Println("按 tagPath 查询...")
					// printLabData(sdk.GetLabOutput("Display", "A", "A1"))
					// printLabData(sdk.GetLabOutput("Display", "A", "A2"))
				}
			}
		}

	}()

	fmt.Print(fmt.Sprintf("实验（%s） 进行中...", Lab))
	userIDPrefix := 34082219900000
	// printPoint(3)
	for userID := 1 + userIDPrefix; userID <= Users+userIDPrefix; userID++ {
		switch Lab {
		case "search":
			ProjectSearchService(strconv.Itoa(userID))
		case "Order":
			ProjectOrderService(strconv.Itoa(userID))
		case "Display":
			ProjectDisplayService(strconv.Itoa(userID))
		}
		if userID%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	c := make(chan bool, 0)
	<-c
}

// ProjectSearchService project search 业务
func ProjectSearchService(userID string) {
	// start lab
	ctx := context.Background()
	labOutput := sdk.LabOutput{
		ProjectID: "search",
		UserID:    userID,
		Time:      time.Now(),
		Data:      make(map[string]interface{}), // TODO: 优化数据结构
		LabPath:   make([]string, 0),            // TODO: 优化数据结构
	}

	searchLabLayer1(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
}

func searchLabLayer1(ctx context.Context) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(sdk.LabOutput)
	project := "search"
	layerID := "1"
	userID := labOutput.UserID
	targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	switch targetZone.Value {
	case "A":
		labOutput.Data["1-Score"] = 90
		labOutput.LabPath = append(labOutput.LabPath, "A")
		labOutput.TagPath = append(labOutput.TagPath, "A")
	case "B":
		labOutput.Data["1-Score"] = 95
		labOutput.LabPath = append(labOutput.LabPath, "B")
		labOutput.TagPath = append(labOutput.TagPath, "B")
	default:
		// ...省略业务
		labOutput.Data["1-Score"] = 95
		labOutput.LabPath = append(labOutput.LabPath, "defualt")
		labOutput.TagPath = append(labOutput.TagPath, "defualt")
	}

	// push lab output
	// fmt.Println("push lab output,debug", labOutput)
	sdk.PushLabOutPut(labOutput)
}

// ProjectDisplayService project Display 业务
func ProjectDisplayService(userID string) {
	// start lab
	ctx := context.Background()
	labOutput := sdk.LabOutput{
		ProjectID: "Display",
		UserID:    userID,
		UserGroup: getUserGroup(userID),
		Time:      time.Now(),
		Data:      make(map[string]interface{}), // TODO: 优化数据结构
		LabPath:   make([]string, 0),            // TODO: 优化数据结构
	}

	displayLabLayer1(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
}

func displayLabLayer1(ctx context.Context) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(sdk.LabOutput)
	project := "Display"
	layerID := "1"
	userID := labOutput.UserID

	// TODO: 选用什么作为分流的条件，应该在ABT server配置好，这样程序可以更灵活，
	targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	// fmt.Println("debug:", targetZone.Value, targetZone.Tag)
	switch true {
	case "A" == targetZone.Value && checkUserGroup(targetZone.UserGroups, labOutput.UserGroup):
		labOutput.Data["Score"] = 83
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "B" == targetZone.Value && checkUserGroup(targetZone.UserGroups, labOutput.UserGroup):
		labOutput.Data["Score"] = 90
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	default:
		// ...省略业务
	}

	// push lab output
	// fmt.Println("push lab output,debug", labOutput)
	sdk.PushLabOutPut(labOutput)
}

// ProjectOrderService project order 业务
func ProjectOrderService(userID string) {
	// start lab
	ctx := context.Background()
	labOutput := sdk.LabOutput{
		ProjectID: "Order ",
		UserID:    userID,
		Time:      time.Now(),
		Data:      make(map[string]interface{}),
		LabPath:   make([]string, 0),
	}

	// TODO: 水平控制
	orderLabLayer1(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	// orderLabLayerA2(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	// orderLabLayerB2(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	// orderLabLayerB3(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))

}

func orderLabLayer1(ctx context.Context) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(sdk.LabOutput)
	project := "Order"
	layerID := "1"
	userID := labOutput.UserID
	targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	// fmt.Println("debug,orderLabLayer1, project, userID, layerID, targetZone.Value ", project, userID, layerID, targetZone.Value)
	switch targetZone.Value {
	case "A1":
		labOutput.Data["1-Score"] = 90
		labOutput.LabPath = append(labOutput.LabPath, "A1")
		labOutput.TagPath = append(labOutput.TagPath, "A1")
		orderLabLayerA2(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	case "B1-1":
		labOutput.Data["1-Score"] = 80
		labOutput.LabPath = append(labOutput.LabPath, "B1-1")
		labOutput.TagPath = append(labOutput.TagPath, "B1-1")
		orderLabLayerB2(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	case "B1-2":
		labOutput.Data["1-Score"] = 95
		labOutput.LabPath = append(labOutput.LabPath, "B2-1")
		labOutput.TagPath = append(labOutput.TagPath, "B2-1")
		orderLabLayerB2(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	default:
		// ...省略业务
		labOutput.LabPath = append(labOutput.LabPath, "defualt")
		labOutput.TagPath = append(labOutput.TagPath, "defualt")
		orderLabLayerA2(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	}
}

func orderLabLayerA2(ctx context.Context) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(sdk.LabOutput)
	project := "Order"
	layerID := "A2"
	userID := labOutput.UserID
	targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	switch targetZone.Value {
	case "A2-1":
		labOutput.Data["A2-Score"] = 90
		labOutput.LabPath = append(labOutput.LabPath, "A2-1")
		labOutput.TagPath = append(labOutput.TagPath, "A2-1")
	case "A2-2":
		labOutput.Data["A2-Score"] = 95
		labOutput.LabPath = append(labOutput.LabPath, "A2-2")
		labOutput.TagPath = append(labOutput.TagPath, "A2-2")
	default:
		// ...省略业务
		labOutput.LabPath = append(labOutput.LabPath, "defualt")
		labOutput.TagPath = append(labOutput.TagPath, "defualt")
	}

	// push lab output
	// fmt.Println("push lab output,debug", labOutput)
	sdk.PushLabOutPut(labOutput)
}

func orderLabLayerB2(ctx context.Context) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(sdk.LabOutput)
	project := "Order"
	layerID := "B2"
	userID := labOutput.UserID
	targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	switch targetZone.Value {
	case "B2-1":
		labOutput.Data["B2-Score"] = 90
		labOutput.LabPath = append(labOutput.LabPath, "B2-1")
		labOutput.TagPath = append(labOutput.TagPath, "B2-1")
		orderLabLayerB3(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	case "B2-2":
		labOutput.Data["B2-Score"] = 95
		labOutput.LabPath = append(labOutput.LabPath, "B2-2")
		labOutput.TagPath = append(labOutput.TagPath, "B2-2")
		orderLabLayerB3(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	default:
		// ...省略业务
		labOutput.LabPath = append(labOutput.LabPath, "defualt")
		labOutput.TagPath = append(labOutput.TagPath, "defualt")
		orderLabLayerB3(context.WithValue(ctx, sdk.CTXKey("output"), labOutput))
	}
}

func orderLabLayerB3(ctx context.Context) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(sdk.LabOutput)
	project := "Order"
	layerID := "B3"
	userID := labOutput.UserID
	targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	switch targetZone.Value {
	case "B3":
		labOutput.Data["B3-Score"] = 90
		labOutput.LabPath = append(labOutput.LabPath, "B3")
		labOutput.TagPath = append(labOutput.TagPath, "B3")
	default:
		// ...省略业务
		labOutput.LabPath = append(labOutput.LabPath, "defualt")
		labOutput.TagPath = append(labOutput.TagPath, "defualt")
	}

	// push lab output
	// fmt.Println("push lab output,debug", labOutput)
	sdk.PushLabOutPut(labOutput)
}

func printLabData(datas []sdk.LabOutput) {
	fmt.Print("实验结果查询:  ")
	// printPoint(3)
	fmt.Println("结果包含总人数:", len(datas))
	fmt.Println("忽略详细打印...")
	// 忽略详细打印
	return

	for _, data := range datas {
		fmt.Println("Project :", data.ProjectID)
		fmt.Println("User :", data.UserID)
		labPath := ""
		for _, path := range data.LabPath {
			if labPath == "" {
				labPath = path
			} else {
				labPath = labPath + "->" + path
			}
		}
		fmt.Println("LabPath :", labPath)
		tagPath := ""
		for _, path := range data.TagPath {
			if tagPath == "" {
				tagPath = path
			} else {
				tagPath = tagPath + "->" + path
			}
		}
		fmt.Println("TagPath :", tagPath)
		for k, v := range data.Data {
			fmt.Println("Lab Data:", k, ":", v)
		}
		fmt.Println("UserGroup :", data.UserGroup)
	}

	fmt.Println("--------------------分割线-----------------------")
	time.Sleep(time.Second)
}

func getUserGroup(userID string) string {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(3) {
	case 0:
		return "VIP"
	case 1:
		return "Normal"
	case 2:
		return "Frequent"
	case 3:
		return "Seldom"
	default:
		return "Others"
	}
}

func checkUserGroup(groups []string, cur string) bool {
	for _, group := range groups {
		if cur == group {
			return true
		}
	}

	return false
}

func printPoint(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	time.Sleep(time.Second)
	fmt.Println("Done!")
	time.Sleep(time.Second)
}
