package services

import "crm-glonass/api/dto"

type Client struct {
	ID    string
	Name  string
	Value string
}

func (cl *Client) CreateClient(data dto.ClientCreate) {

}
