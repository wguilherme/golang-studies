package embedder

import "fmt"

type Embedder struct {
	Title string
}

func (e Embedder) GetTitle(title string) {
	fmt.Println(title)
}

func main() {

}
