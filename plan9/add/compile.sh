$ go tool compile -S main.go >main.asm

#go tool compile命令用于调用Go语言提供的底层命令工具，其中-S参数表示输出汇编格式。