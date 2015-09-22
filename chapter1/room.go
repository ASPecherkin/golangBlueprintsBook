package main

type room struct {
	messages
	forward chan []byte
}
