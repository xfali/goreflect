/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @date 2019/2/21
 * @time 10:54
 * @version V1.0
 * Description: 
 */

package reflectSlice

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	将 []T 切片转换为 []byte
	类似C语言中将其他类型的数组转换为char数组：
 */
func ToBytes(slice interface{}) (data []byte) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("ByteSlice called with non-slice value of type %T", slice))
	}
	h := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

/*
	将 []X 转换为 []Y 切片
	类似C语言中将不同类型的数组转之间的相互转换：
	注意：数据的底层结构需无变化
	该转换操作有一定的风险，需要自己保证安全。主要涉及以下几种类型：

	当结构体中含有指针时，转换会导致垃圾回收的问题。
	如果是 []byte 转 []T 可能会导致起始地址未对齐的问题 （[]byte 有可能从奇数位置切片）。
	该转换操作可能依赖当前系统，不同类型的处理器之间有差异。
	该转换操作的优势是性能和类似void*的泛型，与cgo接口配合使用会更加理想。
 */
func ToType(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}
