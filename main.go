package main

func main() {
	employee := Employee{}
	e := employee.Init()
	if e.Resign == "y" || e.Resign == "Y" {
		e.GetInformation()
	} else {
		e.Decline()
	}
}
