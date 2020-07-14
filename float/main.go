package main


import (
	"encoding/json"
	"fmt"
	//"reflect"
	"strconv"
	"math"
)


func main() {
	var v1 interface{}
	ev1:=json.Unmarshal([]byte("{\"key\":0}"),&v1)
	fmt.Println(ev1,v1)
	if m,ok:=v1.(map[string]interface{});ok{
		if v0,ok:=m["key"];ok{
			fmt.Println(v0)
		}
	}
	f0:=float32(94.0)  // MOVSS   $f32.42bc0000(SB), X0
	fmt.Println(f0)
	f := (93.0 - 20.0) / (140.0 - 20.0)
	f1:=f*300.0
	f2:=(93.0-20.0)/(140.0-20.0)*300.0
	fmt.Println(f1,f2, (93.0-20.0)/(140.0-20.0)*300.0)
	//fmt.Println(reflect.TypeOf(f1),reflect.TypeOf(f2))
	//182.49999999999997 182.5 182.5
	HexTofloat64("4066cfffffffffff")
	HexTofloat64("4066d00000000000")

}

func HexTofloat64(hex string)  {
	//
	n, _:= strconv.ParseUint(hex,16,64)
	//fmt.Sprintf("%016b", n)
	fmt.Println(n)
	//str1 := strconv.FormatFloat(n, 'f', 2, 64)
	//fmt.Println(str1) //3.14

	nf:= math.Float64frombits(uint64(n))
	fmt.Println(nf)

	//fmt.Printf("%x\n", math.Float32bits(-561.2863))
}
// go tool compile -N -l -S main.go
//MOVSD	传送（复制）双字
/**
        0x002f 00047 (main.go:11)       MOVSD   $f64.3fe3777777777777(SB), X1
        0x0037 00055 (main.go:11)       MOVSD   X1, "".f+64(SP)

        0x003d 00061 (main.go:12)       MOVSD   $f64.4066cfffffffffff(SB), X1
        0x0045 00069 (main.go:12)       MOVSD   X1, "".f1+56(SP)

        0x004b 00075 (main.go:13)       MOVSD   $f64.4066d00000000000(SB), X1
        0x0053 00083 (main.go:13)       MOVSD   X1, "".f2+48(SP)

 */