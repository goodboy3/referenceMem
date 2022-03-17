package dataInRedis

import (
	"log"
	"testing"

	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/plugin/redisClient"
	"github.com/goodboy3/referenceMem/plugin/reference"
)

func init() {
	basic.InitLogger()

	//redis
	err := redisClient.Init(redisClient.Config{
		Address:   "127.0.0.1",
		UserName:  "",
		Password:  "",
		Port:      6379,
		KeyPrefix: "userTest:",
		UseTLS:    false,
	})
	if err != nil {
		log.Fatalln("redis init err", err)
	}

	//reference
	err = reference.Init()
	if err != nil {
		log.Fatalln("reference init err", err)
	}
}

func Test_peer(t *testing.T) {
	//
	p := &PeerInfo{
		Tag:      "abcd",
		Location: "USA",
		IP:       "127.0.0.1",
	}
	tag := "abcd"

	err := SetPeer(p, tag)
	if err != nil {
		log.Fatalln("SetPeer err", err, "tag", tag)
	}

	pp := GetPeer(tag, false)
	log.Println(pp)

	DeletePeer(tag)

	pp = GetPeer(tag, false)
	log.Println(pp)
}
