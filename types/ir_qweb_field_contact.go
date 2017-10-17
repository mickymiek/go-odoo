package types

import (
	"time"
)

type IrQwebFieldContact struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldContactNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldContactModel string = "ir.qweb.field.contact"

type IrQwebFieldContacts []IrQwebFieldContact

type IrQwebFieldContactsNil []IrQwebFieldContactNil

func (s *IrQwebFieldContact) NilableType_() interface{} {
	return &IrQwebFieldContactNil{}
}

func (ns *IrQwebFieldContactNil) Type_() interface{} {
	s := &IrQwebFieldContact{}
	return load(ns, s)
}

func (s *IrQwebFieldContacts) NilableType_() interface{} {
	return &IrQwebFieldContactsNil{}
}

func (ns *IrQwebFieldContactsNil) Type_() interface{} {
	s := &IrQwebFieldContacts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldContact))
	}
	return s
}
