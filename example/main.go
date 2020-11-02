package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-abtest/example/lab"
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
	flag.StringVar(&Lab, "lab", "ComplexColor", "Lab name")
	flag.IntVar(&Users, "users", 100, "参与测试的用户总数")
	flag.StringVar(&sdk.Address, "addr", "127.0.0.1", "AB Test Server Address")
	flag.StringVar(&sdk.Port, "port", "9527", "AB Test Server Port")
}

func main() {
	flag.Parse()

	// SetCacheSyncDBFrequency
	fmt.Println("设置缓存同步数据库周期")
	sdk.SetCacheSyncDBFrequency([]string{"Home", "Color", "ComplexColor", "Theme"}, time.Second*60)

	// 开启实验统计线程
	go func() {
		// 实验 数据统计
		fmt.Println("\n开启实验数据统计 线程...")
		for {
			select {
			case <-time.Tick(time.Second * 10):
				switch Lab {
				case "Home":
					// sdk.DebugOutput()
					printLabData(sdk.GetLabOutput("Home", "A", ""))
					printLabData(sdk.GetLabOutput("Home", "B", ""))
					printLabData(sdk.GetLabOutput("Home", "C", ""))
					printLabData(sdk.GetLabOutput("Home", "D", ""))
				case "Color":
					printLabData(sdk.GetLabOutput("Color", "A->D", ""))
					printLabData(sdk.GetLabOutput("Color", "A->E", ""))
					printLabData(sdk.GetLabOutput("Color", "B->D", ""))
					printLabData(sdk.GetLabOutput("Color", "B->E", ""))
					printLabData(sdk.GetLabOutput("Color", "C->D", ""))
					printLabData(sdk.GetLabOutput("Color", "C->E", ""))
				case "ComplexColor":
					sdk.DebugOutput()
					printLabData(sdk.GetLabOutput("ComplexColor", "A->D", ""))
					printLabData(sdk.GetLabOutput("ComplexColor", "B->D", ""))
					printLabData(sdk.GetLabOutput("ComplexColor", "B->E", ""))
					printLabData(sdk.GetLabOutput("ComplexColor", "C->E", ""))
				case "Theme":
					// sdk.DebugOutput()
					printLabData(sdk.GetLabOutput("Theme", "A->C->F", ""))
					printLabData(sdk.GetLabOutput("Theme", "A->C->G", ""))
					printLabData(sdk.GetLabOutput("Theme", "B->C->F", ""))
					printLabData(sdk.GetLabOutput("Theme", "B->C->G", ""))
					printLabData(sdk.GetLabOutput("Theme", "A->D->F", ""))
					printLabData(sdk.GetLabOutput("Theme", "A->D->G", ""))
					printLabData(sdk.GetLabOutput("Theme", "B->D->F", ""))
					printLabData(sdk.GetLabOutput("Theme", "B->D->G", ""))
					printLabData(sdk.GetLabOutput("Theme", "A->E->F", ""))
					printLabData(sdk.GetLabOutput("Theme", "A->E->G", ""))
					printLabData(sdk.GetLabOutput("Theme", "B->E->F", ""))
					printLabData(sdk.GetLabOutput("Theme", "B->E->G", ""))
				}
			}
		}

	}()

	fmt.Print(fmt.Sprintf("实验（%s） 进行中...", Lab))
	userIDPrefix := 34082219900000

	// start lab
	for userID := 1 + userIDPrefix; userID <= Users+userIDPrefix; userID++ {
		// 定义输出0
		labOutput := &sdk.LabOutput{
			ProjectID: Lab,
			UserID:    strconv.Itoa(userID),
			Time:      time.Now(),
			Data:      make(map[string]interface{}), // TODO: 优化数据结构
			LabPath:   make([]string, 0),            // TODO: 优化数据结构
		}
		ctx := context.WithValue(context.Background(), sdk.CTXKey("output"), labOutput)
		switch Lab {
		case "Home":
			lab.HomeLayer1(ctx, fmt.Sprintf("userID:%d,date:%d", userID, time.Now().UnixNano()/3600/24))
		case "Color":
			lab.ColorLayer1(ctx, fmt.Sprintf("userID:%d,date:%d", userID, time.Now().UnixNano()/3600/24))
		case "ComplexColor":
			lab.ComplexLabColorLayer1(ctx, fmt.Sprintf("userID:%d,date:%d", userID, time.Now().UnixNano()/3600/24))
		case "Theme":
			lab.ThemeLayer1(ctx, fmt.Sprintf("userID:%d,date:%d", userID, time.Now().UnixNano()/3600/24))
		}
		if userID%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	c := make(chan bool, 0)
	<-c
}

func printLabData(datas []*sdk.LabOutput) {
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

func printPoint(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	time.Sleep(time.Second)
	fmt.Println("Done!")
	time.Sleep(time.Second)
}
