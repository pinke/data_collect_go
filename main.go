package main

/*
#cgo LDFLAGS: -lstdc++ -fPIC -L. -F. -Wl,-rpath,./ -framework MacDataCollect  -framework CoreFoundation -framework CoreFoundation  -framework IOKit  -framework Foundation
#cgo CPPFLAGS: -fPIC -I./ -I./include -g
#cgo CFLAGS: -fPIC -I./ -I./include -g
#include "stdio.h"
#include "stdlib.h"

extern int _Z28CTP_GetSystemInfoUnAesEncodePcRi(char* pSystemInfo, int* nLen);
extern int _Z17CTP_GetSystemInfoPcRi(char* pSystemInfo, int* nLen);
extern const char* _Z28CTP_GetDataCollectApiVersionv();

//重命令一个,相当于转换
char* CTP_GetSystemInfoUnAesEncode(int *length){
    char *result = (char *)malloc(270); // 分配堆内存
    int ret = _Z28CTP_GetSystemInfoUnAesEncodePcRi(result, length);
   printf("result : %d\n",ret);
    if (ret == 0) {
        result[*length] = '\0';
        return result; // 返回指向堆内存的指针
    }
    free(result); // 错误时释放内存
    return NULL;
}

//char* CTP_GetRealSystemInfo(int *length){
//    char *result = (char *)malloc(270); // 分配堆内存
//    int ret = _Z21CTP_GetRealSystemInfoPcRi(result, length);
//   printf("result : %d\n",ret);
//    if (ret == 0) {
//        result[*length] = '\0';
//        return result; // 返回指向堆内存的指针
//    }
//    free(result); // 错误时释放内存
//    return NULL;
//}


char* CTP_GetSystemInfo(int *length){
    char *result = (char *)malloc(270); // 分配堆内存
    int ret = _Z17CTP_GetSystemInfoPcRi(result, length);
   printf("result : %d\n",ret);
    if (ret == 0) {
        result[*length] = '\0';
        return result; // 返回指向堆内存的指针
    }
    free(result); // 错误时释放内存
    return NULL;
}

*/
import "C"

// GO 包一个
func GetSystemInfoUnAesEncode() (string, int, error) {
	var length C.int
	result := C.CTP_GetSystemInfoUnAesEncode(&length)

	if result != nil {
		//error for object 0x16ba266da: pointer being freed was not allocated
		//defer C.free(unsafe.Pointer(result)) // 使用C.free来释放内存

		// 将C字符串转换为Go字符串
		goStr := C.GoStringN(result, C.int(length))
		//println(goStr)
		return goStr, int(length), nil
	}
	return "", 0, nil
}
func CTP_GetSystemInfo() (string, int, error) {
	var length C.int
	result := C.CTP_GetSystemInfo(&length)

	if result != nil {
		//error for object 0x16ba266da: pointer being freed was not allocated
		//defer C.free(unsafe.Pointer(result)) // 使用C.free来释放内存

		// 将C字符串转换为Go字符串
		goStr := C.GoStringN(result, C.int(length))
		return goStr, int(length), nil
	}
	return "", 0, nil
}

func main() {

	verInfo := C._Z28CTP_GetDataCollectApiVersionv()
	println(C.GoString(verInfo))

	str, l, err := GetSystemInfoUnAesEncode()
	if err == nil {
		println(l, len(str))
	}
	str, l, err = CTP_GetSystemInfo()
	if err == nil {
		println(l, len(str))
	}

	//C.free(unsafe.Pointer(verInfo))

}
