package main

import "fmt"

func minInteger(num string, k int) string {
	sl:=[]byte(num)
	for count:=0;count<k;count++{
		flag:=false
		for i:=0;i<len(num)-1;i++{
			if sl[i]>sl[i+1]{
				sl[i],sl[i+1]=sl[i+1],sl[i]
				fmt.Println("count ++",string(sl),i,count)
				break
			}
			if i==len(num)-2{
				fmt.Println(sl,i)
				flag=true
			}
		}
		fmt.Println("count :",string(sl),count)
		if flag{
			break;
		}
	}
	return string(sl)
}

func main(){
	fmt.Println("4321")
	fmt.Println(minInteger("4321", 4))
		//minInteger("9438957234785635408",23))
}