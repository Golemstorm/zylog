package zylog

import (
	"fmt"
	"github.com/Golemstorm/zylog/color"
)

func Warm(tag string, msg string, args ...interface{}){
	if len(args) > 0 {
		if msg==""{

		}else {
			msg = fmt.Sprintf(msg,args...)
		}

	}

	fmt.Println(color.Yellow(fmt.Sprintf("[WARM]-[%v]:",tag)),msg)
}