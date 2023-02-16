package main

import "fmt"



type securityCode struct{
	code int
}

func newSecurityCode(code int) *securityCode{
	return &securityCode{
		code: code,
	}
}

func (sc *securityCode) checkSecurityCode(code int) error{
	if sc.code != code{
		return fmt.Errorf("wrong code given")
	}

	fmt.Println("Security code verified")
	return nil
}