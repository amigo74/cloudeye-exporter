package model

import (
	"encoding/json"

	"strings"
)

// 创建证书请求参数
type CreateCertificateOption struct {
	// SSL证书的管理状态；暂不支持。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// HTTPS协议使用的证书内容。 取值范围：PEM编码格式。

	Certificate string `json:"certificate"`
	// SSL证书的描述。

	Description *string `json:"description,omitempty"`
	// 服务器证书所签域名。该字段仅type为server时有效。默认值：\"\" 总长度为0-1024，由若干普通域名或泛域名组成，域名之间以\",\"分割，不超过30个域名。 普通域名：由若干字符串组成，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com； 泛域名：在普通域名的基础上仅允许首字母为\"*\"。例：*.test.com

	Domain *string `json:"domain,omitempty"`
	// SSL证书的名称。

	Name *string `json:"name,omitempty"`
	// HTTPS协议使用的私钥。仅type为server时有效。type为server时必选。 取值范围：PEM编码格式。

	PrivateKey *string `json:"private_key,omitempty"`
	// SSL证书所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。 默认值：server

	Type *string `json:"type,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateCertificateOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCertificateOption struct{}"
	}

	return strings.Join([]string{"CreateCertificateOption", string(data)}, " ")
}