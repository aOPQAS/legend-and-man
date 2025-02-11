package main

type DBReader interface {
	Read(string) error
}
