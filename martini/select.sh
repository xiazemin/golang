$cat stack.txt |grep martini |grep '.go' |awk -F'at' '{print $2}' |awk -F'github.com' '{print $2}' |grep 'go-martini'

/go-martini/martini/router.go:418
/go-martini/martini/router.go:285
/go-martini/martini/router.go:132
/go-martini/martini/martini.go:125
/go-martini/martini/martini.go:179
/go-martini/martini/martini.go:170
/go-martini/martini/recovery.go:142
/go-martini/martini/martini.go:179
/go-martini/martini/martini.go:170
/go-martini/martini/logger.go:25
/go-martini/martini/martini.go:179
/go-martini/martini/martini.go:75