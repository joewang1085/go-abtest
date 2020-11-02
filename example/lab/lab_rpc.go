package lab

import (
	"context"
	"fmt"
	"github.com/go-abtest/sdk"
)

// ThemeLayer1 主题 AB test, Layer1 逻辑
func ThemeLayer1(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "Theme"
	layerID := "Layer1"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "A":
		fmt.Println("返回 Theme one")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		labOutput.Data["点击次数"] = count(90, 100)
	case "B":
		fmt.Println("返回 Theme two")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		labOutput.Data["点击次数"] = count(150, 200)

	default:
		fmt.Println("返回 Theme defualt")
	}

	rpc1Client := &rpc1{}
	rpc1Client.RemoteThemeLayer2(ctx, hashkey)

	rpc2Client := &rpc2{}
	rpc2Client.RemoteThemeLayer3(ctx, hashkey)
}

// RemoteThemeLayer2 字体与背景颜色 AB test, Layer1 逻辑
func (rpc *rpc1) RemoteThemeLayer2(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "Theme"
	layerID := "Layer2"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "C":
		fmt.Println("设置字体 黑色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "D":
		fmt.Println("设置字体 红色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "E":
		fmt.Println("设置字体 白色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)

	default:
		fmt.Println("设置字体 默认颜色")
	}
}

// RemoteThemeLayer3 字体与背景 AB test, Layer2 逻辑
func (rpc *rpc2) RemoteThemeLayer3(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "Theme"
	layerID := "Layer3"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "F":
		fmt.Println("设置背景 蓝色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "G":
		fmt.Println("设置背景 绿色")
		// 收集实验数E
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	default:
		fmt.Println("设置背景 默认颜色")
	}

	// 数据上报
	sdk.PushLabOutPut(labOutput)
}

type rpc1 struct{}
type rpc2 struct{}
