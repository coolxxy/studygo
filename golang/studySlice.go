package golang

// go 语言中的 slice 是一种可变数组，其底层实现是
// type SliceHeader struct {
//    Data uintptr
//    Len  int
//    Cap  int
//	}
// 根据其内存分布特点，有一些内存友好的删除元素写法

// 尽可能少的内存复制实现 byte slice 中的空格删除
func TrimSpace(s []byte) []byte {
	if s == nil {
		return nil
	}
	ans := s[:0]
	for _, x := range s {
		if x != ' ' {
			ans = append(ans, x)
		}
	}
	return ans
}

// 可以借助该思路实现任意条件下字符的删除
func TrimByte(s []byte, fn func(byte) bool) []byte {
	if s == nil {
		return nil
	}
	ans := s[:0]
	for _, x := range s {
		if fn(x) == false {
			ans = append(ans, x)
		}
	}
	return ans
}

// 扩展到任意 interface{} slice
func TrimElem(s []interface{}, fn func(interface{}) bool) []interface{} {
	if s == nil {
		return nil
	}
	ans := s[:0]
	for _, x := range s {
		if fn(x) == false {
			ans = append(ans, x)
		}
	}
	return ans
}
