package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NovaDisassociateSecurityGroupRequest struct {
	// 云服务器ID。

	ServerId string `json:"server_id"`

	Body *NovaDisassociateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NovaDisassociateSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupRequest", string(data)}, " ")
}