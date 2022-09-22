package pkg

import "fmt"

type Subscriber struct {
	Name string
}

func (cons *Subscriber) Update(pubName string) {
	fmt.Println("Sending to subscribe from publisher\n", cons.GetName(), pubName)
}
func (cons *Subscriber) GetName() string {
	return cons.Name
}
