package lab

import (
	"context"
	"fmt"
	"github.com/go-abtest/sdk"
	"math"
	"math/rand"
	"time"
)

// HomeLayer1 新主页 AB test, Layer1 逻辑
func HomeLayer1(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "Home"
	layerID := "Layer1"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "A":
		fmt.Println("返回原主页")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		labOutput.Data["点击次数"] = count(90, 100)
	case "B":
		fmt.Println("返回新主页")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		labOutput.Data["点击次数"] = count(150, 200)
	case "C":
		fmt.Println("返回原主页")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		labOutput.Data["点击次数"] = count(90, 100)
	case "D":
		fmt.Println("返回原主页")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		labOutput.Data["点击次数"] = count(90, 100)

	default:
		fmt.Println("返回原主页")
	}

	// 数据上报
	sdk.PushLabOutPut(labOutput)
}

func count(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)%int(math.Abs(float64(max-min))) + min
}
