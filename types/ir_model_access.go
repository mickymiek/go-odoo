package types

import (
	"time"
)

type IrModelAccess struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	GroupId     Many2One  `xmlrpc:"group_id"`
	Id          int64     `xmlrpc:"id"`
	ModelId     Many2One  `xmlrpc:"model_id"`
	Name        string    `xmlrpc:"name"`
	PermCreate  bool      `xmlrpc:"perm_create"`
	PermRead    bool      `xmlrpc:"perm_read"`
	PermUnlink  bool      `xmlrpc:"perm_unlink"`
	PermWrite   bool      `xmlrpc:"perm_write"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrModelAccessNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	GroupId     interface{} `xmlrpc:"group_id"`
	Id          interface{} `xmlrpc:"id"`
	ModelId     interface{} `xmlrpc:"model_id"`
	Name        interface{} `xmlrpc:"name"`
	PermCreate  bool        `xmlrpc:"perm_create"`
	PermRead    bool        `xmlrpc:"perm_read"`
	PermUnlink  bool        `xmlrpc:"perm_unlink"`
	PermWrite   bool        `xmlrpc:"perm_write"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrModelAccessModel string = "ir.model.access"

type IrModelAccesss []IrModelAccess

type IrModelAccesssNil []IrModelAccessNil

func (s *IrModelAccess) NilableType_() interface{} {
	return &IrModelAccessNil{}
}

func (ns *IrModelAccessNil) Type_() interface{} {
	s := &IrModelAccess{}
	return load(ns, s)
}

func (s *IrModelAccesss) NilableType_() interface{} {
	return &IrModelAccesssNil{}
}

func (ns *IrModelAccesssNil) Type_() interface{} {
	s := &IrModelAccesss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModelAccess))
	}
	return s
}
