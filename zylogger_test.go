package zylog

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestWarm(t *testing.T) {
	Warm(time.Now().String(),"%v","shuai")
	log.Fatal(fmt.Errorf("cao"))
}