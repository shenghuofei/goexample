package main

import (
    "github.com/robfig/cron"
    "fmt"
    "time"
)

/*

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

const (
    Second      ParseOption = 1 << iota // Seconds field, default 0
    Minute                              // Minutes field, default 0
    Hour                                // Hours field, default 0
    Dom                                 // Day of month field, default *
    Month                               // Month field, default *
    Dow                                 // Day of week field, default *
    DowOptional                         // Optional day of week field, default *
    Descriptor                          // Allow descriptors such as @monthly, @weekly, etc.
)

*/

func now() time.Time {
	return time.Now()
}

var specParser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

func do_job(spec string) {
    sched,_ := specParser.Parse(spec)

    //sched,_ := cron.Parse("1 */1 * * *")
    now := now()
    pre := now
    next := sched.Next(now)
    for {
        fmt.Println("pre",pre,"next",next,"now",now,"spec",spec)
        select {
            case now = <- time.After(next.Sub(now)):
                fmt.Println("to run",now,"spec",spec)
                pre = next  
                next = sched.Next(next)
                fmt.Println("pre",pre,"next",next,"spec",spec)
        }
    }
}

func main() {
    cronmap := map[int]string{}
    for i:=1;i<5;i++ {
        cronmap[i] = fmt.Sprintf("*/%d * * * *",i)
    }
    fmt.Println(cronmap)
    for _,v := range cronmap {
        go do_job(v)
    }
    select {}
}
