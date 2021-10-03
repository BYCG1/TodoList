package cache

import (
	"fmt"
	"strconv"
)

const (
	//RankKey 每日排名
	RankKey = "rank"
	////ElectricalRank 家电排名
	TaskRank = "TaskRank"
	////AccessoryRank 配件排名
	//AccessoryRank = "acceRank"
)

//ProductViewKey 视频点击数的key
func ProductViewKey(id uint) string {
	fmt.Printf("view:product:%s", strconv.Itoa(int(id)))
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
