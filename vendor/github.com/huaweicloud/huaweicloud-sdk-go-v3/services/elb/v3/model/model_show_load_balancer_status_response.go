package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLoadBalancerStatusResponse struct {
	Statuses *LoadBalancerStatusResult `json:"statuses,omitempty"`
	// 请求ID。  注：自动生成 。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLoadBalancerStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerStatusResponse", string(data)}, " ")
}