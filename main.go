package main

import "github.com/bantublockchain/bantupaysdk-go/funcs"

func NewPayment() (p *funcs.PaymentInfo) {
	return new(funcs.PaymentInfo)
}
