package main

type config struct {
	Next     string
	Previous string
}

func getConfig() config {
	return config{
		Next:     "",
		Previous: "",
	}
}
