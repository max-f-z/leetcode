package main

import "testing"

func Test678(t *testing.T) {
	isValid := checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())")
	t.Log(isValid)
}
