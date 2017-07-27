package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAccountTemplateService struct {
	client *Client
}

func NewAccountAccountTemplateService(c *Client) *AccountAccountTemplateService {
	return &AccountAccountTemplateService{client: c}
}

func (svc *AccountAccountTemplateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAccountTemplateModel, name)
}

func (svc *AccountAccountTemplateService) GetByIds(ids []int) (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getByIds(types.AccountAccountTemplateModel, ids, a)
}

func (svc *AccountAccountTemplateService) GetByName(name string) (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getByName(types.AccountAccountTemplateModel, name, a)
}

func (svc *AccountAccountTemplateService) GetAll() (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getAll(types.AccountAccountTemplateModel, a)
}

func (svc *AccountAccountTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAccountTemplateModel, fields, relations)
}

func (svc *AccountAccountTemplateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAccountTemplateModel, ids, fields, relations)
}

func (svc *AccountAccountTemplateService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAccountTemplateModel, ids)
}