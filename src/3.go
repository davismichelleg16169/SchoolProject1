

package main
import "fmt"
func main(){
	random := rand.Intn(10)
	if random > 5 {
		fmt.Println("It's a high number!")
	} else {
		fmt.Println("It's a low number.")
	}
}