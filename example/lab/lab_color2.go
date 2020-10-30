package lab

import (
	"context"
	"fmt"
	"github.com/go-abtest/sdk"
)

// ComplexLabColorLayer1 字体与背景颜色 AB test, Layer1 逻辑
func ComplexLabColorLayer1(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "ComplexColor"
	layerID := "Layer1"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "A":
		fmt.Println("设置字体 黑色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		// 进入 Layer2-1
		ComplexLabColorLayer2_1(ctx, hashkey)
	case "B":
		fmt.Println("设置字体 红色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		// 进入 Layer2-2
		ComplexLabColorLayer2_2(ctx, hashkey)
	case "C":
		fmt.Println("设置字体 白色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
		// 进入 Layer2-3
		ComplexLabColorLayer2_3(ctx, hashkey)

	default:
		fmt.Println("设置字体默认颜色 黑色")
		fmt.Println("设置背景默认颜色 白色")
	}
}

// ComplexLabColorLayer2_1 字体与背景 AB test, Layer2-1 逻辑
func ComplexLabColorLayer2_1(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "ComplexColor"
	layerID := "Layer2-1"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "D":
		fmt.Println("设置背景 白色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	default:
		fmt.Println("设置背景 白色")
	}

	// 数据上报
	sdk.PushLabOutPut(labOutput)
}

// ComplexLabColorLayer2_2 字体与背景 AB test, Layer2-2 逻辑
func ComplexLabColorLayer2_2(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "ComplexColor"
	layerID := "Layer2-2"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "D":
		fmt.Println("设置背景 白色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	case "E":
		fmt.Println("设置背景 黑色")
		// 收集实验数据
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	default:
		fmt.Println("设置背景 白色")
	}

	// 数据上报
	sdk.PushLabOutPut(labOutput)
}

// ComplexLabColorLayer2_3 字体与背景 AB test, Layer2-3 逻辑
func ComplexLabColorLayer2_3(ctx context.Context, hashkey string) {
	labOutput := ctx.Value(sdk.CTXKey("output")).(*sdk.LabOutput)
	project := "ComplexColor"
	layerID := "Layer2-3"
	targetZone := sdk.GetABTZone(project, hashkey, layerID)
	switch targetZone.Value {
	case "E":
		fmt.Println("设置背景 黑色")
		// 收集实验数E
		labOutput.LabPath = append(labOutput.LabPath, targetZone.Value)
		labOutput.TagPath = append(labOutput.TagPath, targetZone.Tag)
	default:
		fmt.Println("设置背景 白色")
	}

	// 数据上报
	sdk.PushLabOutPut(labOutput)
}
