package main

import (
    "time"
    "fmt"
)

/*
type Weekday int ;Weekday代表一周的某一天

func (d Weekday) String() string
String返回该日（周几）的英文名（"Sunday"、"Monday"，……）

type Month int;Month代表一年的某个月

func (m Month) String() string
String返回月份的英文名（"January"，"February"，……）

Time代表一个纳秒精度的时间点

func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
Date返回一个时区为loc、当地时间为year-month-day hour:min:sec + nsec nanoseconds的时间点

func Now() Time
Now返回当前本地时间

func Parse(layout, value string) (Time, error)
Parse解析一个格式化的时间字符串并返回它代表的时间。layout定义了参考时间2006-01-02 15:04:05的格式

func ParseInLocation(layout, value string, loc *Location) (Time, error)
ParseInLocation类似Parse但有两个重要的不同之处,第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc；第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc

func Unix(sec int64, nsec int64) Time
Unix创建一个本地时间，对应sec和nsec表示的Unix时间（从January 1, 1970 UTC至该时间的秒数和纳秒数）

func (t Time) Location() *Location
Location返回t的地点和时区信息

func (t Time) Zone() (name string, offset int)
Zone计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）

func (t Time) Local() Time
Local返回采用本地和本地时区，但指向同一时间点的Time

func (t Time) UTC() Time
UTC返回采用UTC和零时区，但指向同一时间点的Time

func (t Time) Unix() int64
Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）

func (t Time) UnixNano() int64
UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）

func (t Time) Equal(u Time) bool
判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较

func (t Time) Before(u Time) bool
如果t代表的时间点在u之前，返回真；否则返回假

func (t Time) After(u Time) bool
如果t代表的时间点在u之后，返回真；否则返回假

func (t Time) Date() (year int, month Month, day int)
返回时间点t对应的年、月、日

func (t Time) Clock() (hour, min, sec int)
返回t对应的那一天的时、分、秒

func (t Time) Year() int
返回时间点t对应的年份

func (t Time) Month() Month
返回时间点t对应那一年的第几月

func (t Time) ISOWeek() (year, week int)
返回时间点t对应的ISO 9601标准下的年份和星期编号

func (t Time) YearDay() int
返回时间点t对应的那一年的第几天

func (t Time) Day() int
返回时间点t对应那一月的第几日

func (t Time) Weekday() Weekday
返回时间点t对应的那一周的周几

func (t Time) Hour() int
返回t对应的那一天的第几小时，范围[0, 23]

func (t Time) Minute() int
返回t对应的那一小时的第几分种，范围[0, 59]

func (t Time) Second() int
返回t对应的那一分钟的第几秒，范围[0, 59]

func (t Time) Nanosecond() int
返回t对应的那一秒内的纳秒偏移量，范围[0, 999999999]

func (t Time) Add(d Duration) Time
Add返回时间点t+d

func (t Time) AddDate(years int, months int, days int) Time
AddDate返回增加了给出的年份、月份和天数的时间点Time

func (t Time) Sub(u Time) Duration
返回一个时间段t-u

func (t Time) Round(d Duration) Time
返回距离t最近的时间点，该时间点应该满足从Time零值到该时间点的时间段能整除d；如果有两个满足要求的时间点，距离t相同，会向上舍入；如果d <= 0，会返回t的拷贝

func (t Time) Truncate(d Duration) Time
类似Round，但是返回的是最接近但早于t的时间点；如果d <= 0，会返回t的拷贝

func (t Time) Format(layout string) string
Format根据layout指定的格式返回t代表的时间点的格式化文本表示。layout定义了参考时间

func (t Time) String() string
String返回采用如下格式字符串的格式化时间

type Duration int64;Duration类型代表两个时间点之间经过的时间，以纳秒为单位

func (d Duration) Hours() float64
Hours将时间段表示为float64类型的小时数

func (d Duration) Minutes() float64
Minutes将时间段表示为float64类型的分钟数

func (d Duration) Seconds() float64
Seconds将时间段表示为float64类型的秒数

func (d Duration) Nanoseconds() int64
Nanoseconds将时间段表示为int64类型的纳秒数，等价于int64(d)

func (d Duration) String() string
返回时间段采用"72h3m0.5s"格式的字符串表示

func ParseDuration(s string) (Duration, error)
ParseDuration解析一个时间段字符串,如"300ms"、"-1.5h"、"2h45m"，合法的单位有"ns"、"us" /"µs"、"ms"、"s"、"m"、"h"

func Since(t Time) Duration
Since返回从t到现在经过的时间，等价于time.Now().Sub(t)

type Timer struct {
    C <-chan Time
    // 内含隐藏或非导出字段
}
Timer类型代表单次时间事件。当Timer到期时，当时的时间会被发送给C，除非Timer是被AfterFunc函数创建的

func NewTimer(d Duration) *Timer
NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间

func AfterFunc(d Duration, f func()) *Timer
AfterFunc另起一个go程等待时间段d过去，然后调用f。它返回一个Timer，可以通过调用其Stop方法来取消等待和对f的调用

func (t *Timer) Reset(d Duration) bool
Reset使t重新开始计时，（本方法返回后再）等待时间段d过去后到期。如果调用时t还在等待中会返回真；如果t已经到期或者被停止了会返回假

func (t *Timer) Stop() bool
Stop停止Timer的执行,如果停止了t会返回真；如果t已经被停止或者过期了会返回假

type Ticker struct {
    C <-chan Time // 周期性传递时间信息的通道
    // 内含隐藏或非导出字段
}
Ticker保管一个通道，并每隔一段时间向其传递"tick"

func NewTicker(d Duration) *Ticker
NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间

func (t *Ticker) Stop()
Stop关闭一个Ticker。在关闭后，将不会发送更多的tick信息

func Sleep(d Duration)
Sleep阻塞当前go程至少d代表的时间段

func After(d Duration) <-chan Time
After会在另一线程经过时间段d后向返回值发送当时的时间

func Tick(d Duration) <-chan Time
Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便

*/

func round(){
    t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
    round := []time.Duration{
        time.Nanosecond,
        time.Microsecond,
        time.Millisecond,
        time.Second,
        2 * time.Second,
        time.Minute,
        10 * time.Minute,
        time.Hour,
    }
    for _, d := range round {
        fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
    }
}

func timechan(){
    time.AfterFunc(1*time.Second,func(){
        fmt.Println("after func")
    })
 
    c := time.NewTimer(1*time.Second)
    select {
        case m := <- c.C:
            fmt.Println("timer",m)
    }

    ch := time.After(1*time.Second)
    select {
        case m := <- ch:
            fmt.Println("after",m)
    }

    tk := time.NewTicker(1*time.Second)
    count := 0 
    for {
        select {
            case t := <- tk.C:
                count += 1
                fmt.Println("tkr",count,t)
        }
        if count >= 3 { break }
    }

    tk1 := time.Tick(1*time.Second)
    count = 0 
    for now := range tk1 {
        count += 1
        fmt.Println("tk",count,now)
        if count >= 3 { break }
    }

}

func main() {
    round()
    timechan()
    var w time.Weekday = 1
    fmt.Println(w.String())
    var m time.Month = 1
    fmt.Println(m.String())
    t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    fmt.Printf("Go launched at %s\n", t.Local())
    year, month, day := time.Now().Date()
    fmt.Println(year,month,day)
    longForm := "Jan 2, 2006 at 3:04pm (MST)"
    t, _ = time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
    fmt.Println(t)
    shortForm := "2006-Jan-02"
    t, _ = time.Parse(shortForm, "2013-Feb-03")
    fmt.Println(t)

    loc, _ := time.LoadLocation("Europe/Berlin")
    longForm = "Jan 2, 2006 at 3:04pm (MST)"
    t, _ = time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
    fmt.Println(t)
    // Note: without explicit zone, returns time in given location.
    shortForm = "2006-Jan-02"
    t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
    fmt.Println(t)
    fmt.Println(time.Unix(1513063135, 0))
    t=time.Now()
    fmt.Println(t.Location().String())
    fmt.Println(t.Zone())
    fmt.Println(t.Local())
    fmt.Println(t.UTC())
    fmt.Println(t.Unix())
    fmt.Println(t.UnixNano())
    fmt.Println(t.Equal(time.Now()))
    fmt.Println(t.Before(time.Now()))
    fmt.Println(t.After(time.Now()))
    fmt.Println(t.Date())
    fmt.Println(t.Clock())
    fmt.Println(t.Month())
    fmt.Println(t.ISOWeek())
    fmt.Println(t.Add(-1*time.Hour))
    fmt.Println(time.Now().Sub(t))
    fmt.Println(t.Format("2006-01-02 15:04:05"))
    fmt.Println(t.String())
    fmt.Println(time.ParseDuration("100us"))
    fmt.Println((10*time.Hour).String())
}
