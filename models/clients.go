package models

type Client struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty"`
	CreatedAt string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type AgentData struct {
	AgentId       string      `json:"agentId"`
	ParentId      string      `json:"parentId"`
	ParentName    string      `json:"parentName"`
	Name          string      `json:"name"`
	FullName      string      `json:"fullName"`
	AgentInfoType int         `json:"agentInfoType"`
	IsForeign     interface{} `json:"isForeign"`
	District      string      `json:"district"`
	Region        string      `json:"region"`
	City          string      `json:"city"`
	Inn           string      `json:"inn"`
	Kpp           string      `json:"kpp"`
	Address       string      `json:"address"`
	AddressFact   string      `json:"addressFact"`
	Email         string      `json:"email"`
	Director      string      `json:"director"`
	BankName      string      `json:"bankName"`
	BankBIK       string      `json:"bankBIK"`
	BankRS        string      `json:"bankRS"`
	BankKS        string      `json:"bankKS"`
}
