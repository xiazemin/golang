$  go run -race main.go
a is  3
==================
WARNING: DATA RACE
Write at 0x00c420018070 by goroutine 6:
  main.main.func1()
      /Users/didi/goLang/src/github.com/xiazemin/race/main.go:10 +0x3b

Previous write at 0x00c420018070 by main goroutine:
  main.main()
      /Users/didi/goLang/src/github.com/xiazemin/race/main.go:12 +0x8e

Goroutine 6 (running) created at:
  main.main()
      /Users/didi/goLang/src/github.com/xiazemin/race/main.go:9 +0x7d
==================
Found 1 data race(s)
exit status 66
