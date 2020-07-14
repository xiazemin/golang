package main
import "fmt"

func main(){
defer func (){
echo()
}()

defer func (){
echo()
}()
}
func echo(){
fmt.Println("second");
}
