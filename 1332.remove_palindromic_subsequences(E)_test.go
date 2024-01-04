package main

import (
	"log"
	"testing"
)

func Test_removePalindromeSub(t *testing.T) {
	a := removePalindromeSub("baabb")
	if a != 2 {
		t.Error("Fail result: ", a)
		return
	}
	log.Println("success")
}
