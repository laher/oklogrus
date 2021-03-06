package main

import (
	"io"
	"log"
	"net/url"
	"time"

	"github.com/oklog/oklog/pkg/forward"
	"github.com/sirupsen/logrus"
)

// for the moment this only works with https://github.com/laher/oklog, branch feature/pkg-forward
// PR to follow.
func main() {
	log.Println("logrus -> oklog")
	u, err := url.Parse("tcp://localhost:7651")
	if err != nil {
		log.Fatal(err)
	}
	urls := []*url.URL{u}
	l := logrus.New()

	f := forward.NewBufferedForwarder(urls, "oklogrus:", 5)
	r, w := io.Pipe()
	go func() {
		err := f.Forward(r)
		if err != nil {
			log.Fatalf("Error forwarding: %v", err)
		}
	}()
	l.Out = w
	l.Infof("hi")
	for {
		time.Sleep(time.Second)
		l.Infof("now %v", time.Now())
	}
}
