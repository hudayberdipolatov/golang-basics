package main

import "fmt"

// Panic
// Programmanyň ýerine ýetirilişini derrew gutarmak
// üçin howsala beýanyny ulanýarys. Programmamyz käbir möhüm
// ýalňyşlyklar sebäpli dikeldilip bilinmejek derejä ýeten bolsa,
// howsala ulanmak iň gowusydyr.
// panic yazylandan son kod setirleri yerine yetirilmeyar

func main() {

	fmt.Println("Help! Something bad is happening.")
	panic("Ending the program")
	fmt.Println("Waiting to execute")

}
