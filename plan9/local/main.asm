"".main STEXT nosplit size=1 args=0x0 locals=0x0
	0x0000 00000 (main.go:3)	TEXT	"".main(SB), NOSPLIT, $0-0
	0x0000 00000 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:7)	RET
	0x0000 c3                                               .
"".add STEXT nosplit size=21 args=0x18 locals=0x0
	0x0000 00000 (main.go:9)	TEXT	"".add(SB), NOSPLIT, $0-24
	0x0000 00000 (main.go:9)	FUNCDATA	$0, gclocals·54241e171da8af6ae173d69da0236748(SB)
	0x0000 00000 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:9)	MOVQ	"".a+8(SP), AX
	0x0005 00005 (main.go:9)	MOVQ	"".b+16(SP), CX
	0x000a 00010 (main.go:11)	LEAQ	1(CX)(AX*1), AX
	0x000f 00015 (main.go:11)	MOVQ	AX, "".~r2+24(SP)
	0x0014 00020 (main.go:11)	RET
	0x0000 48 8b 44 24 08 48 8b 4c 24 10 48 8d 44 01 01 48  H.D$.H.L$.H.D..H
	0x0010 89 44 24 18 c3                                   .D$..
"".init STEXT size=79 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	72
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JLS	47
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	JNE	56
	0x0031 00049 (<autogenerated>:1)	PCDATA	$0, $0
	0x0031 00049 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0036 00054 (<autogenerated>:1)	UNDEF
	0x0038 00056 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0043 00067 (<autogenerated>:1)	ADDQ	$8, SP
	0x0047 00071 (<autogenerated>:1)	RET
	0x0048 00072 (<autogenerated>:1)	NOP
	0x0048 00072 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0048 00072 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004d 00077 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 39 48  eH..%....H;a.v9H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 02 48  ...............H
	0x0040 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb b1     .,$H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 73+4 t=8 runtime.morestack_noctxt+0
go.info."".main SDWARFINFO size=29
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 00           .............
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+1
go.range."".main SDWARFRANGE size=0
go.info."".add SDWARFINFO size=63
	0x0000 02 22 22 2e 61 64 64 00 00 00 00 00 00 00 00 00  ."".add.........
	0x0010 00 00 00 00 00 00 00 00 01 9c 01 05 61 00 01 9c  ............a...
	0x0020 00 00 00 00 05 62 00 04 9c 11 08 22 00 00 00 00  .....b....."....
	0x0030 05 7e 72 32 00 04 9c 11 10 22 00 00 00 00 00     .~r2.....".....
	rel 8+8 t=1 "".add+0
	rel 16+8 t=1 "".add+21
	rel 32+4 t=28 go.info.int+0
	rel 44+4 t=28 go.info.int+0
	rel 58+4 t=28 go.info.int+0
go.range."".add SDWARFRANGE size=0
go.info."".init SDWARFINFO size=29
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 00           .............
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+79
go.range."".init SDWARFRANGE size=0
"".initdone· SNOPTRBSS size=1
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·54241e171da8af6ae173d69da0236748 SRODATA dupok size=9
	0x0000 01 00 00 00 03 00 00 00 00                       .........
