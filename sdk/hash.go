/*Package sdk ...
*
	 hash 方法
*
*/
package sdk

import (
	"github.com/spaolacci/murmur3"
)

// hash 计算hash值: projectID 用于指定实验项目，globalID, layerID and date 用于计算hash，total用于取模
// 要求：1. 每层都随机 2. 每层都均匀 3. 可以按日期分流 4. 其他
func hash(projectID, hashkey, layerID string, total uint32) uint32 {

	// set seed
	var seed uint32 = total

	// get total weight
	if total == 0 {
		total = DefualtTotalWeight
	}

	// [1,100]
	return murmur3.Sum32WithSeed([]byte(hashkey+layerID), seed)%total + 1
}
