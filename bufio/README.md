# readstring-vs-readline.go 
## 此程序对readstring 和readline 的性能进行对比
## 结论:ReadLine 读取文件更快，原因是由于 ReadString 后端调用 ReadBytes，而 ReadBytes 多次使用 copy 方法造成大量耗时，因此对大文件的读操作优先使用ReadLine
