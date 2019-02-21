/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @date 2019/2/21
 * @time 11:04
 * @version V1.0
 * Description: 
 */

package test

import (
	"fmt"
	"goreflect/src/reflectSlice"
	"reflect"
	"testing"
)

type RGB struct {
	R, G, B uint8
}

type BGR struct {
	B, G, R uint8
}

func (bgr BGR) String() string {
	return fmt.Sprintf("B%d G%d R%d", bgr.B, bgr.G, bgr.R)
}

func TestSlice(t *testing.T) {
	arr := []RGB { {0,0,0},{1,1,1}, {2,2,2} }
	bytes := reflectSlice.ToBytes(arr)
	for _, v := range bytes {
		fmt.Println(v)
	}

	BGRs := reflectSlice.ToType(arr, reflect.TypeOf([]BGR(nil)))
	for _, v := range BGRs.([]BGR) {
		fmt.Println(v)
	}
}
