package main
import (
    "os"
    "os/exec"
    "path/filepath"
    "path"
)

//https://wizardforcel.gitbooks.io/golang-stdlib-ref/content/105.html  go标准库说明

/* path包基本函数说明
Base 函数  : 返回路径的最后一个元素
Dir 函数   : 返回路径里面除最后一个元素之外的其他所有元素
Ext 函数   : 返回路径使用的文件扩展名
IsAbs 函数 : 检查路径是否为绝对路径
Join 函数  : 将任意数量的路径元素拼接为单个路径， 并在有需要时添加用于分割的斜杠
Match 函数 : 报告文件名是否与给定的 shell 文件名模式相匹配
Split 函数 : 根据路径的最后一个斜杠， 将路径划分为目录部分和文件名部分
*/

/* filepath包基本函数说明(filepath包含以上path包中的函数,作用相同)
Abs函数    : 获取 path 的绝对路径
*/

func main() {
    curpath,_ := os.Getwd()
    println("current path:",curpath)
    file, _ := exec.LookPath(os.Args[0])
    execpath, _ := filepath.Abs(file)
    println("exec abs path:",execpath)
    dir,execname := path.Split(execpath)
    println("exec dir:",dir,"exec program name:",execname)
    execname = path.Base(execpath)
    println("exec program name:",execname)
}
