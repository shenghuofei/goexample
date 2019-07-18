package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

/******************************************************
信号 值 处理动作 发出信号的原因 
------------------------------------------------------
SIGHUP 1 A 终端挂起或者控制进程终止 
SIGINT 2 A 键盘中断（如break键被按下） 
SIGQUIT 3 C 键盘的退出键被按下 
SIGILL 4 C 非法指令 
SIGABRT 6 C 由abort(3)发出的退出指令 
SIGFPE 8 C 浮点异常 
SIGKILL 9 AEF Kill信号 
SIGSEGV 11 C 无效的内存引用 
SIGPIPE 13 A 管道破裂: 写一个没有读端口的管道 
SIGALRM 14 A 由alarm(2)发出的信号 
SIGTERM 15 A 终止信号 
SIGUSR1 30,10,16 A 用户自定义信号1 
SIGUSR2 31,12,17 A 用户自定义信号2 
SIGCHLD 20,17,18 B 子进程结束信号 
SIGCONT 19,18,25 进程继续（曾被停止的进程） 
SIGSTOP 17,19,23 DEF 终止进程 
SIGTSTP 18,20,24 D 控制终端（tty）上按下停止键 
SIGTTIN 21,21,26 D 后台进程企图从控制终端读 
SIGTTOU 22,22,27 D 后台进程企图从控制终端写 
********************************************************/
func handleSignal() {
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
    // 监听信号
    for {
        sig := <-ch
        log.Printf("signal receive: %v\n", sig)
        ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
        log.Println(ctx)
        switch sig {
        case syscall.SIGINT, syscall.SIGTERM: // 终止进程执行
            log.Println("shutdown")
            signal.Stop(ch)
            log.Println("graceful shutdown")
            os.Exit(0)
        case syscall.SIGUSR2: // 进程热重启
            log.Println("reload")
            reload() // 执行热重启函数
            log.Println("graceful reload")
        }
    }
}

func main() {
    reload()
    go handleSignal()
    for {
        log.Println("main",cfg)
        time.Sleep(1*time.Second)
    }
}
