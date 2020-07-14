 go build -o main main.go
 go tool objdump  -s "main\.\w+" main

package main
func main() {
  go func() {}()
}


TEXT runtime.main.func1(SB) /usr/local/go/src/runtime/proc.go
  proc.go:128           0x10466b0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:128           0x10466b9               483b6110                CMPQ 0x10(CX), SP
  proc.go:128           0x10466bd               7631                    JBE 0x10466f0
  proc.go:128           0x10466bf               4883ec18                SUBQ $0x18, SP
  proc.go:128           0x10466c3               48896c2410              MOVQ BP, 0x10(SP)
  proc.go:128           0x10466c8               488d6c2410              LEAQ 0x10(SP), BP
  proc.go:128           0x10466cd               488d05e4920200          LEAQ 0x292e4(IP), AX
  proc.go:129           0x10466d4               48890424                MOVQ AX, 0(SP)
  proc.go:129           0x10466d8               48c744240800000000      MOVQ $0x0, 0x8(SP)
  proc.go:129           0x10466e1               e8fa17feff              CALL runtime.newm(SB)
  proc.go:130           0x10466e6               488b6c2410              MOVQ 0x10(SP), BP
  proc.go:130           0x10466eb               4883c418                ADDQ $0x18, SP
  proc.go:130           0x10466ef               c3                      RET
  proc.go:128           0x10466f0               e87b100000              CALL runtime.morestack_noctxt(SB)
  proc.go:128           0x10466f5               ebb9                    JMP runtime.main.func1(SB)
  :-1                   0x10466f7               cc                      INT $0x3
  :-1                   0x10466f8               cc                      INT $0x3
  :-1                   0x10466f9               cc                      INT $0x3
  :-1                   0x10466fa               cc                      INT $0x3
  :-1                   0x10466fb               cc                      INT $0x3
  :-1                   0x10466fc               cc                      INT $0x3
  :-1                   0x10466fd               cc                      INT $0x3
  :-1                   0x10466fe               cc                      INT $0x3
  :-1                   0x10466ff               cc                      INT $0x3

TEXT runtime.main.func2(SB) /usr/local/go/src/runtime/proc.go
  proc.go:151           0x1046700               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:151           0x1046709               483b6110                CMPQ 0x10(CX), SP
  proc.go:151           0x104670d               7628                    JBE 0x1046737
  proc.go:151           0x104670f               4883ec08                SUBQ $0x8, SP
  proc.go:151           0x1046713               48892c24                MOVQ BP, 0(SP)
  proc.go:151           0x1046717               488d2c24                LEAQ 0(SP), BP
  proc.go:151           0x104671b               488b442410              MOVQ 0x10(SP), AX
  proc.go:152           0x1046720               0fb600                  MOVZX 0(AX), AX
  proc.go:152           0x1046723               84c0                    TESTL AL, AL
  proc.go:152           0x1046725               7509                    JNE 0x1046730
  proc.go:155           0x1046727               488b2c24                MOVQ 0(SP), BP
  proc.go:155           0x104672b               4883c408                ADDQ $0x8, SP
  proc.go:155           0x104672f               c3                      RET
  proc.go:153           0x1046730               e81b5efeff              CALL runtime.unlockOSThread(SB)
  proc.go:152           0x1046735               ebf0                    JMP 0x1046727
  proc.go:151           0x1046737               e834100000              CALL runtime.morestack_noctxt(SB)
  proc.go:151           0x104673c               ebc2                    JMP runtime.main.func2(SB)
  :-1                   0x104673e               cc                      INT $0x3
  :-1                   0x104673f               cc                      INT $0x3

TEXT main.main(SB) /Users/didi/goLang/src/github.com/xiazemin/goroutine/main.go
  main.go:2             0x104f360               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  main.go:2             0x104f369               483b6110                CMPQ 0x10(CX), SP
  main.go:2             0x104f36d               7630                    JBE 0x104f39f
  main.go:2             0x104f36f               4883ec18                SUBQ $0x18, SP
  main.go:2             0x104f373               48896c2410              MOVQ BP, 0x10(SP)
  main.go:2             0x104f378               488d6c2410              LEAQ 0x10(SP), BP
  main.go:3             0x104f37d               c7042400000000          MOVL $0x0, 0(SP)
  main.go:3             0x104f384               488d0545030200          LEAQ 0x20345(IP), AX
  main.go:3             0x104f38b               4889442408              MOVQ AX, 0x8(SP)
  main.go:3             0x104f390               e85bc5fdff              CALL runtime.newproc(SB)
  main.go:4             0x104f395               488b6c2410              MOVQ 0x10(SP), BP
  main.go:4             0x104f39a               4883c418                ADDQ $0x18, SP
  main.go:4             0x104f39e               c3                      RET
  main.go:2             0x104f39f               e8cc83ffff              CALL runtime.morestack_noctxt(SB)
  main.go:2             0x104f3a4               ebba                    JMP main.main(SB)
  :-1                   0x104f3a6               cc                      INT $0x3
  :-1                   0x104f3a7               cc                      INT $0x3
  :-1                   0x104f3a8               cc                      INT $0x3
  :-1                   0x104f3a9               cc                      INT $0x3
  :-1                   0x104f3aa               cc                      INT $0x3
  :-1                   0x104f3ab               cc                      INT $0x3
  :-1                   0x104f3ac               cc                      INT $0x3
  :-1                   0x104f3ad               cc                      INT $0x3
  :-1                   0x104f3ae               cc                      INT $0x3
  :-1                   0x104f3af               cc                      INT $0x3

TEXT main.main.func1(SB) /Users/didi/goLang/src/github.com/xiazemin/goroutine/main.go
  main.go:3             0x104f3b0               c3                      RET
  :-1                   0x104f3b1               cc                      INT $0x3
  :-1                   0x104f3b2               cc                      INT $0x3
  :-1                   0x104f3b3               cc                      INT $0x3
  :-1                   0x104f3b4               cc                      INT $0x3
  :-1                   0x104f3b5               cc                      INT $0x3
  :-1                   0x104f3b6               cc                      INT $0x3
  :-1                   0x104f3b7               cc                      INT $0x3
  :-1                   0x104f3b8               cc                      INT $0x3
  :-1                   0x104f3b9               cc                      INT $0x3
  :-1                   0x104f3ba               cc                      INT $0x3
  :-1                   0x104f3bb               cc                      INT $0x3
  :-1                   0x104f3bc               cc                      INT $0x3
  :-1                   0x104f3bd               cc                      INT $0x3
  :-1                   0x104f3be               cc                      INT $0x3
  :-1                   0x104f3bf               cc                      INT $0x3

TEXT main.init(SB) <autogenerated>
  <autogenerated>:1     0x104f3c0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  <autogenerated>:1     0x104f3c9               483b6110                CMPQ 0x10(CX), SP
  <autogenerated>:1     0x104f3cd               7639                    JBE 0x104f408
  <autogenerated>:1     0x104f3cf               4883ec08                SUBQ $0x8, SP
  <autogenerated>:1     0x104f3d3               48892c24                MOVQ BP, 0(SP)
  <autogenerated>:1     0x104f3d7               488d2c24                LEAQ 0(SP), BP
  <autogenerated>:1     0x104f3db               0fb605de5f0700          MOVZX 0x75fde(IP), AX
  <autogenerated>:1     0x104f3e2               3c01                    CMPL $0x1, AL
  <autogenerated>:1     0x104f3e4               7609                    JBE 0x104f3ef
  <autogenerated>:1     0x104f3e6               488b2c24                MOVQ 0(SP), BP
  <autogenerated>:1     0x104f3ea               4883c408                ADDQ $0x8, SP
  <autogenerated>:1     0x104f3ee               c3                      RET
  <autogenerated>:1     0x104f3ef               7507                    JNE 0x104f3f8
  <autogenerated>:1     0x104f3f1               e8da1dfdff              CALL runtime.throwinit(SB)
  <autogenerated>:1     0x104f3f6               0f0b                    UD2
  <autogenerated>:1     0x104f3f8               c605c15f070002          MOVB $0x2, 0x75fc1(IP)
  <autogenerated>:1     0x104f3ff               488b2c24                MOVQ 0(SP), BP
  <autogenerated>:1     0x104f403               4883c408                ADDQ $0x8, SP
  <autogenerated>:1     0x104f407               c3                      RET
  <autogenerated>:1     0x104f408               e86383ffff              CALL runtime.morestack_noctxt(SB)
  <autogenerated>:1     0x104f40d               ebb1                    JMP main.init(SB)





package main
func main() {
  go Test()
}

func Test () {}


TEXT runtime.main.func1(SB) /usr/local/go/src/runtime/proc.go
  proc.go:128           0x10466b0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:128           0x10466b9               483b6110                CMPQ 0x10(CX), SP
  proc.go:128           0x10466bd               7631                    JBE 0x10466f0
  proc.go:128           0x10466bf               4883ec18                SUBQ $0x18, SP
  proc.go:128           0x10466c3               48896c2410              MOVQ BP, 0x10(SP)
  proc.go:128           0x10466c8               488d6c2410              LEAQ 0x10(SP), BP
  proc.go:128           0x10466cd               488d05e4920200          LEAQ 0x292e4(IP), AX
  proc.go:129           0x10466d4               48890424                MOVQ AX, 0(SP)
  proc.go:129           0x10466d8               48c744240800000000      MOVQ $0x0, 0x8(SP)
  proc.go:129           0x10466e1               e8fa17feff              CALL runtime.newm(SB)
  proc.go:130           0x10466e6               488b6c2410              MOVQ 0x10(SP), BP
  proc.go:130           0x10466eb               4883c418                ADDQ $0x18, SP
  proc.go:130           0x10466ef               c3                      RET
  proc.go:128           0x10466f0               e87b100000              CALL runtime.morestack_noctxt(SB)
  proc.go:128           0x10466f5               ebb9                    JMP runtime.main.func1(SB)
  :-1                   0x10466f7               cc                      INT $0x3
  :-1                   0x10466f8               cc                      INT $0x3
  :-1                   0x10466f9               cc                      INT $0x3
  :-1                   0x10466fa               cc                      INT $0x3
  :-1                   0x10466fb               cc                      INT $0x3
  :-1                   0x10466fc               cc                      INT $0x3
  :-1                   0x10466fd               cc                      INT $0x3
  :-1                   0x10466fe               cc                      INT $0x3
  :-1                   0x10466ff               cc                      INT $0x3

TEXT runtime.main.func2(SB) /usr/local/go/src/runtime/proc.go
  proc.go:151           0x1046700               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:151           0x1046709               483b6110                CMPQ 0x10(CX), SP
  proc.go:151           0x104670d               7628                    JBE 0x1046737
  proc.go:151           0x104670f               4883ec08                SUBQ $0x8, SP
  proc.go:151           0x1046713               48892c24                MOVQ BP, 0(SP)
  proc.go:151           0x1046717               488d2c24                LEAQ 0(SP), BP
  proc.go:151           0x104671b               488b442410              MOVQ 0x10(SP), AX
  proc.go:152           0x1046720               0fb600                  MOVZX 0(AX), AX
  proc.go:152           0x1046723               84c0                    TESTL AL, AL
  proc.go:152           0x1046725               7509                    JNE 0x1046730
  proc.go:155           0x1046727               488b2c24                MOVQ 0(SP), BP
  proc.go:155           0x104672b               4883c408                ADDQ $0x8, SP
  proc.go:155           0x104672f               c3                      RET
  proc.go:153           0x1046730               e81b5efeff              CALL runtime.unlockOSThread(SB)
  proc.go:152           0x1046735               ebf0                    JMP 0x1046727
  proc.go:151           0x1046737               e834100000              CALL runtime.morestack_noctxt(SB)
  proc.go:151           0x104673c               ebc2                    JMP runtime.main.func2(SB)
  :-1                   0x104673e               cc                      INT $0x3
  :-1                   0x104673f               cc                      INT $0x3

TEXT main.main(SB) /Users/didi/goLang/src/github.com/xiazemin/goroutine/main.go
  main.go:2             0x104f360               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  main.go:2             0x104f369               483b6110                CMPQ 0x10(CX), SP
  main.go:2             0x104f36d               7630                    JBE 0x104f39f
  main.go:2             0x104f36f               4883ec18                SUBQ $0x18, SP
  main.go:2             0x104f373               48896c2410              MOVQ BP, 0x10(SP)
  main.go:2             0x104f378               488d6c2410              LEAQ 0x10(SP), BP
  main.go:3             0x104f37d               c7042400000000          MOVL $0x0, 0(SP)
  main.go:3             0x104f384               488d0545030200          LEAQ 0x20345(IP), AX
  main.go:3             0x104f38b               4889442408              MOVQ AX, 0x8(SP)
  main.go:3             0x104f390               e85bc5fdff              CALL runtime.newproc(SB)
  main.go:4             0x104f395               488b6c2410              MOVQ 0x10(SP), BP
  main.go:4             0x104f39a               4883c418                ADDQ $0x18, SP
  main.go:4             0x104f39e               c3                      RET
  main.go:2             0x104f39f               e8cc83ffff              CALL runtime.morestack_noctxt(SB)
  main.go:2             0x104f3a4               ebba                    JMP main.main(SB)
  :-1                   0x104f3a6               cc                      INT $0x3
  :-1                   0x104f3a7               cc                      INT $0x3
  :-1                   0x104f3a8               cc                      INT $0x3
  :-1                   0x104f3a9               cc                      INT $0x3
  :-1                   0x104f3aa               cc                      INT $0x3
  :-1                   0x104f3ab               cc                      INT $0x3
  :-1                   0x104f3ac               cc                      INT $0x3
  :-1                   0x104f3ad               cc                      INT $0x3
  :-1                   0x104f3ae               cc                      INT $0x3
  :-1                   0x104f3af               cc                      INT $0x3

TEXT main.Test(SB) /Users/didi/goLang/src/github.com/xiazemin/goroutine/main.go
  main.go:6             0x104f3b0               c3                      RET
  :-1                   0x104f3b1               cc                      INT $0x3
  :-1                   0x104f3b2               cc                      INT $0x3
  :-1                   0x104f3b3               cc                      INT $0x3
  :-1                   0x104f3b4               cc                      INT $0x3
  :-1                   0x104f3b5               cc                      INT $0x3
  :-1                   0x104f3b6               cc                      INT $0x3
  :-1                   0x104f3b7               cc                      INT $0x3
  :-1                   0x104f3b8               cc                      INT $0x3
  :-1                   0x104f3b9               cc                      INT $0x3
  :-1                   0x104f3ba               cc                      INT $0x3
  :-1                   0x104f3bb               cc                      INT $0x3
  :-1                   0x104f3bc               cc                      INT $0x3
  :-1                   0x104f3bd               cc                      INT $0x3
  :-1                   0x104f3be               cc                      INT $0x3
  :-1                   0x104f3bf               cc                      INT $0x3

TEXT main.init(SB) <autogenerated>
  <autogenerated>:1     0x104f3c0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  <autogenerated>:1     0x104f3c9               483b6110                CMPQ 0x10(CX), SP
  <autogenerated>:1     0x104f3cd               7639                    JBE 0x104f408
  <autogenerated>:1     0x104f3cf               4883ec08                SUBQ $0x8, SP
  <autogenerated>:1     0x104f3d3               48892c24                MOVQ BP, 0(SP)
  <autogenerated>:1     0x104f3d7               488d2c24                LEAQ 0(SP), BP
  <autogenerated>:1     0x104f3db               0fb605de5f0700          MOVZX 0x75fde(IP), AX
  <autogenerated>:1     0x104f3e2               3c01                    CMPL $0x1, AL
  <autogenerated>:1     0x104f3e4               7609                    JBE 0x104f3ef
  <autogenerated>:1     0x104f3e6               488b2c24                MOVQ 0(SP), BP
  <autogenerated>:1     0x104f3ea               4883c408                ADDQ $0x8, SP
  <autogenerated>:1     0x104f3ee               c3                      RET
  <autogenerated>:1     0x104f3ef               7507                    JNE 0x104f3f8
  <autogenerated>:1     0x104f3f1               e8da1dfdff              CALL runtime.throwinit(SB)
  <autogenerated>:1     0x104f3f6               0f0b                    UD2
  <autogenerated>:1     0x104f3f8               c605c15f070002          MOVB $0x2, 0x75fc1(IP)
  <autogenerated>:1     0x104f3ff               488b2c24                MOVQ 0(SP), BP
  <autogenerated>:1     0x104f403               4883c408                ADDQ $0x8, SP
  <autogenerated>:1     0x104f407               c3                      RET
  <autogenerated>:1     0x104f408               e86383ffff              CALL runtime.morestack_noctxt(SB)
  <autogenerated>:1     0x104f40d               ebb1                    JMP main.init(SB)



 go tool objdump  -s "main\.\w+|\w+\.main" main

TEXT runtime.main(SB) /usr/local/go/src/runtime/proc.go
  proc.go:109           0x1023e70               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:109           0x1023e79               483b6110                CMPQ 0x10(CX), SP
  proc.go:109           0x1023e7d               0f86ac030000            JBE 0x102422f
  proc.go:109           0x1023e83               4883ec58                SUBQ $0x58, SP
  proc.go:109           0x1023e87               48896c2450              MOVQ BP, 0x50(SP)
  proc.go:109           0x1023e8c               488d6c2450              LEAQ 0x50(SP), BP
  proc.go:110           0x1023e91               65488b0425a0080000      MOVQ GS:0x8a0, AX
  proc.go:110           0x1023e9a               4889442440              MOVQ AX, 0x40(SP)
  proc.go:114           0x1023e9f               488b4830                MOVQ 0x30(AX), CX
  proc.go:114           0x1023ea3               488b09                  MOVQ 0(CX), CX
  proc.go:114           0x1023ea6               48c7813001000000000000  MOVQ $0x0, 0x130(CX)
  proc.go:120           0x1023eb1               48c7058421080000ca9a3b  MOVQ $0x3b9aca00, 0x82184(IP)
  proc.go:126           0x1023ebc               c6050a150a0001          MOVB $0x1, 0xa150a(IP)
  proc.go:126           0x1023ec3               488d0d8eb90400          LEAQ 0x4b98e(IP), CX
  proc.go:128           0x1023eca               48890c24                MOVQ CX, 0(SP)
  proc.go:128           0x1023ece               e84d370200              CALL runtime.systemstack(SB)
  proc.go:3179          0x1023ed3               65488b0425a0080000      MOVQ GS:0x8a0, AX
  proc.go:3179          0x1023edc               488b4030                MOVQ 0x30(AX), AX
  proc.go:3179          0x1023ee0               8b88bc020000            MOVL 0x2bc(AX), CX
  proc.go:3179          0x1023ee6               83c102                  ADDL $0x2, CX
  proc.go:3179          0x1023ee9               8988bc020000            MOVL CX, 0x2bc(AX)
  proc.go:3162          0x1023eef               65488b0425a0080000      MOVQ GS:0x8a0, AX
  proc.go:3163          0x1023ef8               488b4830                MOVQ 0x30(AX), CX
  proc.go:3163          0x1023efc               8401                    TESTB AL, 0(CX)
  proc.go:3163          0x1023efe               8b159c160a00            MOVL 0xa169c(IP), DX
  proc.go:3163          0x1023f04               488d9930010000          LEAQ 0x130(CX), BX
  proc.go:3163          0x1023f0b               85d2                    TESTL DX, DX
  proc.go:3163          0x1023f0d               0f855d020000            JNE 0x1024170
  proc.go:3163          0x1023f13               48898130010000          MOVQ AX, 0x130(CX)
  proc.go:3164          0x1023f1a               8b0d80160a00            MOVL 0xa1680(IP), CX
  proc.go:3164          0x1023f20               488d90e0000000          LEAQ 0xe0(AX), DX
  proc.go:3164          0x1023f27               488b5830                MOVQ 0x30(AX), BX
  proc.go:3164          0x1023f2b               85c9                    TESTL CX, CX
  proc.go:3164          0x1023f2d               0f852a020000            JNE 0x102415d
  proc.go:3164          0x1023f33               488998e0000000          MOVQ BX, 0xe0(AX)
  proc.go:3164          0x1023f3a               488b442440              MOVQ 0x40(SP), AX
  proc.go:140           0x1023f3f               488b4030                MOVQ 0x30(AX), AX
  proc.go:140           0x1023f43               488d0dd63b0800          LEAQ 0x83bd6(IP), CX
  proc.go:140           0x1023f4a               4839c8                  CMPQ CX, AX
  proc.go:140           0x1023f4d               0f85c1020000            JNE 0x1024214
  proc.go:144           0x1023f53               e858310200              CALL runtime.init(SB)
  proc.go:145           0x1023f58               e8f36c0200              CALL runtime.nanotime(SB)
  proc.go:145           0x1023f5d               488b0424                MOVQ 0(SP), AX
  proc.go:145           0x1023f61               4885c0                  TESTQ AX, AX
  proc.go:145           0x1023f64               0f848f020000            JE 0x10241f9
  proc.go:150           0x1023f6a               c644243701              MOVB $0x1, 0x37(SP)
  proc.go:150           0x1023f6f               488d442437              LEAQ 0x37(SP), AX
  proc.go:155           0x1023f74               4889442410              MOVQ AX, 0x10(SP)
  proc.go:151           0x1023f79               c7042408000000          MOVL $0x8, 0(SP)
  proc.go:151           0x1023f80               488d05d9b80400          LEAQ 0x4b8d9(IP), AX
  proc.go:151           0x1023f87               4889442408              MOVQ AX, 0x8(SP)
  proc.go:151           0x1023f8c               e87fd2ffff              CALL runtime.deferproc(SB)
  proc.go:151           0x1023f91               85c0                    TESTL AX, AX
  proc.go:151           0x1023f93               0f85b4010000            JNE 0x102414d
  proc.go:159           0x1023f99               e8b26c0200              CALL runtime.nanotime(SB)
  proc.go:159           0x1023f9e               488b0424                MOVQ 0(SP), AX
  proc.go:159           0x1023fa2               48890587150a00          MOVQ AX, 0xa1587(IP)
  proc.go:161           0x1023fa9               e822cffeff              CALL runtime.gcenable(SB)
  proc.go:161           0x1023fae               488d05cb4d0300          LEAQ 0x34dcb(IP), AX
  proc.go:163           0x1023fb5               48890424                MOVQ AX, 0(SP)
  proc.go:163           0x1023fb9               48c744240800000000      MOVQ $0x0, 0x8(SP)
  proc.go:163           0x1023fc2               e879f0fdff              CALL runtime.makechan(SB)
  proc.go:163           0x1023fc7               488b442410              MOVQ 0x10(SP), AX
  proc.go:163           0x1023fcc               8b0dce150a00            MOVL 0xa15ce(IP), CX
  proc.go:163           0x1023fd2               85c9                    TESTL CX, CX
  proc.go:163           0x1023fd4               0f8559010000            JNE 0x1024133
  proc.go:163           0x1023fda               48890567320800          MOVQ AX, 0x83267(IP)
  proc.go:164           0x1023fe1               0fb605e2130a00          MOVZX 0xa13e2(IP), AX
  proc.go:164           0x1023fe8               84c0                    TESTL AL, AL
  proc.go:164           0x1023fea               7452                    JE 0x102403e
  proc.go:165           0x1023fec               488b05dd310800          MOVQ 0x831dd(IP), AX
  proc.go:165           0x1023ff3               4885c0                  TESTQ AX, AX
  proc.go:165           0x1023ff6               0f84e2010000            JE 0x10241de
  proc.go:169           0x1023ffc               488b05dd310800          MOVQ 0x831dd(IP), AX
  proc.go:169           0x1024003               4885c0                  TESTQ AX, AX
  proc.go:169           0x1024006               0f84b7010000            JE 0x10241c3
  proc.go:172           0x102400c               488b05d5310800          MOVQ 0x831d5(IP), AX
  proc.go:172           0x1024013               4885c0                  TESTQ AX, AX
  proc.go:172           0x1024016               0f848c010000            JE 0x10241a8
  proc.go:176           0x102401c               488b05a5310800          MOVQ 0x831a5(IP), AX
  proc.go:176           0x1024023               4885c0                  TESTQ AX, AX
  proc.go:176           0x1024026               0f8461010000            JE 0x102418d
  proc.go:179           0x102402c               48890424                MOVQ AX, 0(SP)
  proc.go:179           0x1024030               48c744240800000000      MOVQ $0x0, 0x8(SP)
  proc.go:179           0x1024039               e8e2dffdff              CALL runtime.cgocall(SB)
  proc.go:183           0x102403e               488b0523b80400          MOVQ 0x4b823(IP), AX
  proc.go:183           0x1024045               488d151cb80400          LEAQ 0x4b81c(IP), DX
  proc.go:183           0x102404c               ffd0                    CALL AX
  proc.go:184           0x102404e               488b05f3310800          MOVQ 0x831f3(IP), AX
  proc.go:184           0x1024055               48890424                MOVQ AX, 0(SP)
  proc.go:184           0x1024059               e8d2fcfdff              CALL runtime.closechan(SB)
  proc.go:186           0x102405e               c644243700              MOVB $0x0, 0x37(SP)
  proc.go:187           0x1024063               e8e8840000              CALL runtime.unlockOSThread(SB)
  proc.go:189           0x1024068               0fb6055a130a00          MOVZX 0xa135a(IP), AX
  proc.go:189           0x102406f               84c0                    TESTL AL, AL
  proc.go:189           0x1024071               0f85ac000000            JNE 0x1024123
  proc.go:189           0x1024077               0fb6054d130a00          MOVZX 0xa134d(IP), AX
  proc.go:189           0x102407e               84c0                    TESTL AL, AL
  proc.go:189           0x1024080               0f859d000000            JNE 0x1024123
  proc.go:195           0x1024086               488b05e3b70400          MOVQ 0x4b7e3(IP), AX
  proc.go:195           0x102408d               488d15dcb70400          LEAQ 0x4b7dc(IP), DX
  proc.go:195           0x1024094               ffd0                    CALL AX
  proc.go:204           0x1024096               8b0584130a00            MOVL 0xa1384(IP), AX
  proc.go:204           0x102409c               85c0                    TESTL AX, AX
  proc.go:204           0x102409e               7428                    JE 0x10240c8
  proc.go:204           0x10240a0               31c0                    XORL AX, AX
  proc.go:206           0x10240a2               eb12                    JMP 0x10240b6
  proc.go:206           0x10240a4               4889442438              MOVQ AX, 0x38(SP)
  proc.go:210           0x10240a9               e832030000              CALL runtime.Gosched(SB)
  proc.go:210           0x10240ae               488b442438              MOVQ 0x38(SP), AX
  proc.go:206           0x10240b3               48ffc0                  INCQ AX
  proc.go:206           0x10240b6               483de8030000            CMPQ $0x3e8, AX
  proc.go:206           0x10240bc               7d0a                    JGE 0x10240c8
  proc.go:207           0x10240be               8b0d5c130a00            MOVL 0xa135c(IP), CX
  proc.go:207           0x10240c4               85c9                    TESTL CX, CX
  proc.go:207           0x10240c6               75dc                    JNE 0x10240a4
  proc.go:213           0x10240c8               8b054a130a00            MOVL 0xa134a(IP), AX
  proc.go:213           0x10240ce               85c0                    TESTL AX, AX
  proc.go:213           0x10240d0               7516                    JNE 0x10240e8
  proc.go:217           0x10240d2               c7042400000000          MOVL $0x0, 0(SP)
  proc.go:217           0x10240d9               e8126a0200              CALL runtime.exit(SB)
  proc.go:218           0x10240de               31c0                    XORL AX, AX
  proc.go:220           0x10240e0               c70000000000            MOVL $0x0, 0(AX)
  proc.go:218           0x10240e6               ebf6                    JMP 0x10240de
  proc.go:214           0x10240e8               48c7042400000000        MOVQ $0x0, 0(SP)
  proc.go:214           0x10240f0               48c744240800000000      MOVQ $0x0, 0x8(SP)
  proc.go:214           0x10240f9               488d053c780400          LEAQ 0x4783c(IP), AX
  proc.go:214           0x1024100               4889442410              MOVQ AX, 0x10(SP)
  proc.go:214           0x1024105               48c744241809000000      MOVQ $0x9, 0x18(SP)
  proc.go:214           0x102410e               c644242010              MOVB $0x10, 0x20(SP)
  proc.go:214           0x1024113               48c744242801000000      MOVQ $0x1, 0x28(SP)
  proc.go:214           0x102411c               e8ef020000              CALL runtime.gopark(SB)
  proc.go:213           0x1024121               ebaf                    JMP 0x10240d2
  proc.go:192           0x1024123               90                      NOPL
  proc.go:192           0x1024124               e877daffff              CALL runtime.deferreturn(SB)
  proc.go:192           0x1024129               488b6c2450              MOVQ 0x50(SP), BP
  proc.go:192           0x102412e               4883c458                ADDQ $0x58, SP
  proc.go:192           0x1024132               c3                      RET
  proc.go:192           0x1024133               488d0d0e310800          LEAQ 0x8310e(IP), CX
  proc.go:163           0x102413a               48890c24                MOVQ CX, 0(SP)
  proc.go:163           0x102413e               4889442408              MOVQ AX, 0x8(SP)
  proc.go:163           0x1024143               e8f889feff              CALL runtime.writebarrierptr(SB)
  proc.go:163           0x1024148               e994feffff              JMP 0x1023fe1
  proc.go:151           0x102414d               90                      NOPL
  proc.go:151           0x102414e               e84ddaffff              CALL runtime.deferreturn(SB)
  proc.go:151           0x1024153               488b6c2450              MOVQ 0x50(SP), BP
  proc.go:151           0x1024158               4883c458                ADDQ $0x58, SP
  proc.go:151           0x102415c               c3                      RET
  proc.go:3164          0x102415d               48891424                MOVQ DX, 0(SP)
  proc.go:3164          0x1024161               48895c2408              MOVQ BX, 0x8(SP)
  proc.go:3164          0x1024166               e8d589feff              CALL runtime.writebarrierptr(SB)
  proc.go:3164          0x102416b               e9cafdffff              JMP 0x1023f3a
  proc.go:3164          0x1024170               4889442448              MOVQ AX, 0x48(SP)
  proc.go:3163          0x1024175               48891c24                MOVQ BX, 0(SP)
  proc.go:3163          0x1024179               4889442408              MOVQ AX, 0x8(SP)
  proc.go:3163          0x102417e               e8bd89feff              CALL runtime.writebarrierptr(SB)
  proc.go:3163          0x1024183               488b442448              MOVQ 0x48(SP), AX
  proc.go:3163          0x1024188               e98dfdffff              JMP 0x1023f1a
  proc.go:3163          0x102418d               488d055ba50400          LEAQ 0x4a55b(IP), AX
  proc.go:177           0x1024194               48890424                MOVQ AX, 0(SP)
  proc.go:177           0x1024198               48c744240825000000      MOVQ $0x25, 0x8(SP)
  proc.go:177           0x10241a1               e86ae7ffff              CALL runtime.throw(SB)
  proc.go:177           0x10241a6               0f0b                    UD2
  proc.go:177           0x10241a8               488d05ea870400          LEAQ 0x487ea(IP), AX
  proc.go:173           0x10241af               48890424                MOVQ AX, 0(SP)
  proc.go:173           0x10241b3               48c744240815000000      MOVQ $0x15, 0x8(SP)
  proc.go:173           0x10241bc               e84fe7ffff              CALL runtime.throw(SB)
  proc.go:173           0x10241c1               0f0b                    UD2
  proc.go:173           0x10241c3               488d0594830400          LEAQ 0x48394(IP), AX
  proc.go:170           0x10241ca               48890424                MOVQ AX, 0(SP)
  proc.go:170           0x10241ce               48c744240813000000      MOVQ $0x13, 0x8(SP)
  proc.go:170           0x10241d7               e834e7ffff              CALL runtime.throw(SB)
  proc.go:170           0x10241dc               0f0b                    UD2
  proc.go:170           0x10241de               488d05c08e0400          LEAQ 0x48ec0(IP), AX
  proc.go:166           0x10241e5               48890424                MOVQ AX, 0(SP)
  proc.go:166           0x10241e9               48c744240819000000      MOVQ $0x19, 0x8(SP)
  proc.go:166           0x10241f2               e819e7ffff              CALL runtime.throw(SB)
  proc.go:166           0x10241f7               0f0b                    UD2
  proc.go:166           0x10241f9               488d05118c0400          LEAQ 0x48c11(IP), AX
  proc.go:146           0x1024200               48890424                MOVQ AX, 0(SP)
  proc.go:146           0x1024204               48c744240817000000      MOVQ $0x17, 0x8(SP)
  proc.go:146           0x102420d               e8fee6ffff              CALL runtime.throw(SB)
  proc.go:146           0x1024212               0f0b                    UD2
  proc.go:146           0x1024214               488d0595890400          LEAQ 0x48995(IP), AX
  proc.go:141           0x102421b               48890424                MOVQ AX, 0(SP)
  proc.go:141           0x102421f               48c744240816000000      MOVQ $0x16, 0x8(SP)
  proc.go:141           0x1024228               e8e3e6ffff              CALL runtime.throw(SB)
  proc.go:141           0x102422d               0f0b                    UD2
  proc.go:109           0x102422f               e83c350200              CALL runtime.morestack_noctxt(SB)
  proc.go:109           0x1024234               e937fcffff              JMP runtime.main(SB)
  :-1                   0x1024239               cc                      INT $0x3
  :-1                   0x102423a               cc                      INT $0x3
  :-1                   0x102423b               cc                      INT $0x3
  :-1                   0x102423c               cc                      INT $0x3
  :-1                   0x102423d               cc                      INT $0x3
  :-1                   0x102423e               cc                      INT $0x3
  :-1                   0x102423f               cc                      INT $0x3

TEXT runtime.main.func1(SB) /usr/local/go/src/runtime/proc.go
  proc.go:128           0x10466b0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:128           0x10466b9               483b6110                CMPQ 0x10(CX), SP
  proc.go:128           0x10466bd               7631                    JBE 0x10466f0
  proc.go:128           0x10466bf               4883ec18                SUBQ $0x18, SP
  proc.go:128           0x10466c3               48896c2410              MOVQ BP, 0x10(SP)
  proc.go:128           0x10466c8               488d6c2410              LEAQ 0x10(SP), BP
  proc.go:128           0x10466cd               488d05e4920200          LEAQ 0x292e4(IP), AX
  proc.go:129           0x10466d4               48890424                MOVQ AX, 0(SP)
  proc.go:129           0x10466d8               48c744240800000000      MOVQ $0x0, 0x8(SP)
  proc.go:129           0x10466e1               e8fa17feff              CALL runtime.newm(SB)
  proc.go:130           0x10466e6               488b6c2410              MOVQ 0x10(SP), BP
  proc.go:130           0x10466eb               4883c418                ADDQ $0x18, SP
  proc.go:130           0x10466ef               c3                      RET
  proc.go:128           0x10466f0               e87b100000              CALL runtime.morestack_noctxt(SB)
  proc.go:128           0x10466f5               ebb9                    JMP runtime.main.func1(SB)
  :-1                   0x10466f7               cc                      INT $0x3
  :-1                   0x10466f8               cc                      INT $0x3
  :-1                   0x10466f9               cc                      INT $0x3
  :-1                   0x10466fa               cc                      INT $0x3
  :-1                   0x10466fb               cc                      INT $0x3
  :-1                   0x10466fc               cc                      INT $0x3
  :-1                   0x10466fd               cc                      INT $0x3
  :-1                   0x10466fe               cc                      INT $0x3
  :-1                   0x10466ff               cc                      INT $0x3

TEXT runtime.main.func2(SB) /usr/local/go/src/runtime/proc.go
  proc.go:151           0x1046700               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  proc.go:151           0x1046709               483b6110                CMPQ 0x10(CX), SP
  proc.go:151           0x104670d               7628                    JBE 0x1046737
  proc.go:151           0x104670f               4883ec08                SUBQ $0x8, SP
  proc.go:151           0x1046713               48892c24                MOVQ BP, 0(SP)
  proc.go:151           0x1046717               488d2c24                LEAQ 0(SP), BP
  proc.go:151           0x104671b               488b442410              MOVQ 0x10(SP), AX
  proc.go:152           0x1046720               0fb600                  MOVZX 0(AX), AX
  proc.go:152           0x1046723               84c0                    TESTL AL, AL
  proc.go:152           0x1046725               7509                    JNE 0x1046730
  proc.go:155           0x1046727               488b2c24                MOVQ 0(SP), BP
  proc.go:155           0x104672b               4883c408                ADDQ $0x8, SP
  proc.go:155           0x104672f               c3                      RET
  proc.go:153           0x1046730               e81b5efeff              CALL runtime.unlockOSThread(SB)
  proc.go:152           0x1046735               ebf0                    JMP 0x1046727
  proc.go:151           0x1046737               e834100000              CALL runtime.morestack_noctxt(SB)
  proc.go:151           0x104673c               ebc2                    JMP runtime.main.func2(SB)
  :-1                   0x104673e               cc                      INT $0x3
  :-1                   0x104673f               cc                      INT $0x3

TEXT main.main(SB) /Users/didi/goLang/src/github.com/xiazemin/goroutine/main.go
  main.go:2             0x104f360               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  main.go:2             0x104f369               483b6110                CMPQ 0x10(CX), SP
  main.go:2             0x104f36d               7630                    JBE 0x104f39f
  main.go:2             0x104f36f               4883ec18                SUBQ $0x18, SP
  main.go:2             0x104f373               48896c2410              MOVQ BP, 0x10(SP)
  main.go:2             0x104f378               488d6c2410              LEAQ 0x10(SP), BP
  main.go:3             0x104f37d               c7042400000000          MOVL $0x0, 0(SP)
  main.go:3             0x104f384               488d0545030200          LEAQ 0x20345(IP), AX
  main.go:3             0x104f38b               4889442408              MOVQ AX, 0x8(SP)
  main.go:3             0x104f390               e85bc5fdff              CALL runtime.newproc(SB)
  main.go:4             0x104f395               488b6c2410              MOVQ 0x10(SP), BP
  main.go:4             0x104f39a               4883c418                ADDQ $0x18, SP
  main.go:4             0x104f39e               c3                      RET
  main.go:2             0x104f39f               e8cc83ffff              CALL runtime.morestack_noctxt(SB)
  main.go:2             0x104f3a4               ebba                    JMP main.main(SB)
  :-1                   0x104f3a6               cc                      INT $0x3
  :-1                   0x104f3a7               cc                      INT $0x3
  :-1                   0x104f3a8               cc                      INT $0x3
  :-1                   0x104f3a9               cc                      INT $0x3
  :-1                   0x104f3aa               cc                      INT $0x3
  :-1                   0x104f3ab               cc                      INT $0x3
  :-1                   0x104f3ac               cc                      INT $0x3
  :-1                   0x104f3ad               cc                      INT $0x3
  :-1                   0x104f3ae               cc                      INT $0x3
  :-1                   0x104f3af               cc                      INT $0x3

TEXT main.Test(SB) /Users/didi/goLang/src/github.com/xiazemin/goroutine/main.go
  main.go:6             0x104f3b0               c3                      RET
  :-1                   0x104f3b1               cc                      INT $0x3
  :-1                   0x104f3b2               cc                      INT $0x3
  :-1                   0x104f3b3               cc                      INT $0x3
  :-1                   0x104f3b4               cc                      INT $0x3
  :-1                   0x104f3b5               cc                      INT $0x3
  :-1                   0x104f3b6               cc                      INT $0x3
  :-1                   0x104f3b7               cc                      INT $0x3
  :-1                   0x104f3b8               cc                      INT $0x3
  :-1                   0x104f3b9               cc                      INT $0x3
  :-1                   0x104f3ba               cc                      INT $0x3
  :-1                   0x104f3bb               cc                      INT $0x3
  :-1                   0x104f3bc               cc                      INT $0x3
  :-1                   0x104f3bd               cc                      INT $0x3
  :-1                   0x104f3be               cc                      INT $0x3
  :-1                   0x104f3bf               cc                      INT $0x3

TEXT main.init(SB) <autogenerated>
  <autogenerated>:1     0x104f3c0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  <autogenerated>:1     0x104f3c9               483b6110                CMPQ 0x10(CX), SP
  <autogenerated>:1     0x104f3cd               7639                    JBE 0x104f408
  <autogenerated>:1     0x104f3cf               4883ec08                SUBQ $0x8, SP
  <autogenerated>:1     0x104f3d3               48892c24                MOVQ BP, 0(SP)
  <autogenerated>:1     0x104f3d7               488d2c24                LEAQ 0(SP), BP
  <autogenerated>:1     0x104f3db               0fb605de5f0700          MOVZX 0x75fde(IP), AX
  <autogenerated>:1     0x104f3e2               3c01                    CMPL $0x1, AL
  <autogenerated>:1     0x104f3e4               7609                    JBE 0x104f3ef
  <autogenerated>:1     0x104f3e6               488b2c24                MOVQ 0(SP), BP
  <autogenerated>:1     0x104f3ea               4883c408                ADDQ $0x8, SP
  <autogenerated>:1     0x104f3ee               c3                      RET
  <autogenerated>:1     0x104f3ef               7507                    JNE 0x104f3f8
  <autogenerated>:1     0x104f3f1               e8da1dfdff              CALL runtime.throwinit(SB)
  <autogenerated>:1     0x104f3f6               0f0b                    UD2
  <autogenerated>:1     0x104f3f8               c605c15f070002          MOVB $0x2, 0x75fc1(IP)
  <autogenerated>:1     0x104f3ff               488b2c24                MOVQ 0(SP), BP
  <autogenerated>:1     0x104f403               4883c408                ADDQ $0x8, SP
  <autogenerated>:1     0x104f407               c3                      RET
  <autogenerated>:1     0x104f408               e86383ffff              CALL runtime.morestack_noctxt(SB)
  <autogenerated>:1     0x104f40d               ebb1                    JMP main.init(SB)
