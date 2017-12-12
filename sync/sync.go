package main

import (
    "sync"
    "fmt"
    "time"
)

/*

Once是只执行一次动作的对象

func (o *Once) Do(f func())
Do方法当且仅当第一次被调用时才执行函数f

Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁

func (m *Mutex) Lock()
Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁

func (m *Mutex) Unlock()
Unlock方法解锁m，如果m未加锁会导致运行时错误

RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁

func (rw *RWMutex) Lock()
Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入

func (rw *RWMutex) Unlock()
Unlock方法解除rw的写入锁状态，如果m未加写入锁会导致运行时错误

func (rw *RWMutex) RLock()
RLock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取

func (rw *RWMutex) RUnlock()
Runlock方法解除rw的读取锁状态，如果m未加读取锁会导致运行时错误

Cond实现了一个条件变量，一个线程集合地，供线程等待或者宣布某事件的发生

func (c *Cond) Broadcast()
Broadcast唤醒所有等待c的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定

func (c *Cond) Signal()
Signal唤醒等待c的一个线程（如果存在）。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定

WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束

Pool是一个可以分别存取的临时对象的集合,Pool可以安全的被多个线程同时使用

*/

func condition() bool {
    time.Sleep(1*time.Second)
    return true
}

func main() {
    var once sync.Once
    onceBody := func() {
        fmt.Println("Only once")
    }
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i := 0; i < 10; i++ {
        <-done
    }

    var wg sync.WaitGroup
    var urls = []string{
        "http://www.golang.org/",
        "http://www.google.com/",
        "http://www.somestupidname.com/",
    }
    for _, url := range urls {
        // Increment the WaitGroup counter.
        wg.Add(1)
        // Launch a goroutine to fetch the URL.
        go func(url string) {
            // Decrement the counter when the goroutine completes.
            defer wg.Done()
            // Fetch the URL.
            time.Sleep(1*time.Second)
            fmt.Println(url)
        }(url)
    }
    // Wait for all HTTP fetches to complete.
    wg.Wait()

    pool := sync.Pool{}
    pool.Put("asdf")
    fmt.Println(pool.Get())
}
