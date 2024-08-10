# MacOS 下用 go调用 MacDataCollect.framework 的例子

## 这里对framework做了个修改了,要么编译ld时报文件格式错

``` shell
mv MacDataCollect.framework/MacDataCollect MacDataCollect.framework/MacDataCollect_bak
cp MacDataCollect.framework/Versions/A/MacDataCollect  MacDataCollect.framework/MacDataCollect
```
接下来就是简单的go build go run即可
```shell
 go build -o bin/test .
 go run .
 ./bin/test
```
