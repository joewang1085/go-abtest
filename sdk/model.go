/*Package sdk ...
*
	实验配置: 一组 Zone
*
*/
package sdk

import (
	"time"
)

const (

	// DefualtTotalWeight ..
	DefualtTotalWeight = uint32(100)
)

var (
	// Port abtest server 端口
	Port string = "9527"

	// Address abtest server 地址
	Address string = "127.0.0.1"

	// DefaultTimeout http request timeout
	DefaultTimeout time.Duration = time.Second * 60
)

// Layer 层，只属于某一个域，同一个域的不同层正交, ParentZonesIDs为空则表示第一层
type Layer struct {
	LayerID     string  // 全局唯一：格式：layerxxx (组合主键)
	ParentZones []*Zone // 父域,用来规则正确性校验   // TODO: 层的父域互斥，不能正交，增加校验
	TotalWeight uint32  // 总权重，该层的所有域的权重和等于总权重
}

// Zone 域，同一层的域互斥
type Zone struct {
	ProjectID  string   // projectID,全局唯一， 格式: xxx (组合主键)
	Layer               // 层信息
	ZoneID     string   // 全局唯一，ID，格式：zonexxx (组合主键)
	Weight              // 比重,同一层的所有域总和等于100，可以自定义
	Value      string   // 用户自定义值，用于匹配实验
	Tag        string   // 用户自定义值，用于实验标签或者分类 // 可要可不要
	UserGroups []string // 流量过滤标签
}

// Weight zone的流量范围
type Weight struct {
	Min uint32
	Max uint32
}

// CTXKey 自定义类型
type CTXKey string

// LabOutput 实验输出
type LabOutput struct {
	ProjectID string
	UserID    string
	UserGroup string // 流量标签或分类
	Time      time.Time
	Data      LabData
	LabPath   []string // TODO:设计数据机构，方便快速查询和匹配，
	TagPath   []string // TODO
}

// LabData 实验数据
type LabData map[string]interface{}
