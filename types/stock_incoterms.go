package types

import (
	"time"
)

type StockIncoterms struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	Code        string    `xmlrpc:"code"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockIncotermsNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	Code        interface{} `xmlrpc:"code"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockIncotermsModel string = "stock.incoterms"

type StockIncotermss []StockIncoterms

type StockIncotermssNil []StockIncotermsNil

func (s *StockIncoterms) NilableType_() interface{} {
	return &StockIncotermsNil{}
}

func (ns *StockIncotermsNil) Type_() interface{} {
	s := &StockIncoterms{}
	return load(ns, s)
}

func (s *StockIncotermss) NilableType_() interface{} {
	return &StockIncotermssNil{}
}

func (ns *StockIncotermssNil) Type_() interface{} {
	s := &StockIncotermss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockIncoterms))
	}
	return s
}
