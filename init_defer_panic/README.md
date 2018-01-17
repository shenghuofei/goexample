# go语言init，defer，panic使用细节

1. 一个包中可以有多个init函数，执行顺序从前到后
2. defer执行顺序与语句顺序相反，且defer在return之前执行
3. panic 在return的时候执行，所以defer语句一定在panic前执行
