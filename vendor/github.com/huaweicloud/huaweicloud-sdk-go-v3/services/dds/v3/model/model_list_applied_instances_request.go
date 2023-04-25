package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppliedInstancesRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAppliedInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAppliedInstancesRequest", string(data)}, " ")
}
