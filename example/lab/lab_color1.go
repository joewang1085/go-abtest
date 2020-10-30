package lab

import (
	"context"
	"fmt"
	"github.com/go-abtest/sdk"
)

// ColorLayer1 字体与背景颜色 AB test, Layer1 逻辑
func ColorLayer1(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "Color"
	layerID := "Layer1"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "A":
		fmt.Println("设置字体 黑色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "B":
		fmt.Println("设置字体 红色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "C":
		fmt.Println("设置字体 白色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)

	default:
		fmt.Println("设置字体默认颜色 黑色")
	}

	// 进入下一层
	ColorLayer2(ctx, hashkey)
}

// ColorLayer2 字体与背景 AB test, Layer2 逻辑
func ColorLayer2(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "Color"
	layerID := "Layer2"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "D":
		fmt.Println("设置背景 蓝色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "E":
		fmt.Println("设置背景 绿色")
		// 收集实验数E
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	default:
		fmt.Println("设置背景 蓝色")
	}

	// 数据上报
	sdk.PushLabOutPut(labOutput)
}
