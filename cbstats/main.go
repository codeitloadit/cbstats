package main

func main() {
	RunServer()
}

// HandleError is shorthand for the err -> panic idiom.
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
