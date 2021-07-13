package tests

import (
	"Chizuru-GO/chizuru/nekosAPI"
	"log"
	"testing"
)

func TestNekos(t *testing.T){
	r, err := nekosAPI.DoRequest("neko")
	if err!= nil{
		log.Fatal(err.Error())
	}
	log.Println(r.URL)
}
