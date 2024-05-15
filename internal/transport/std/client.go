package std

import (
	"log"
	"sync"
)

var chanID int = 0

type Client struct {
	chans map[int]chan struct{}
	mut   sync.RWMutex
}

var clients = Client{
	chans: make(map[int]chan struct{}, 6),
}

func (c *Client) AddChan(ch chan struct{}) int {
	c.mut.Lock()
	defer func() {
		chanID++
		c.mut.Unlock()
	}()
	c.chans[chanID] = ch
	log.Printf("new connection with chanel id: %v", chanID)
	return chanID
}

func (c *Client) DeleteChan(id int) {
	c.mut.Lock()
	defer c.mut.Unlock()
	delete(c.chans, id)
	log.Printf("close connection id: %v", id)
}

func (c *Client) Update() {
	for _, c := range c.chans {
		c <- struct{}{}
	}
}
