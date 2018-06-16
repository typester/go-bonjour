package bonjour

import (
	"github.com/typester/go-cocoa-eventloop"
	"testing"
	"time"
)

func TestBrowserInit(t *testing.T) {
	b := NewBrowser()

	if b.ptr == nil {
		t.Errorf("ptr should not be nil\n")
	}

	b.Free()

	if b.ptr != nil {
		t.Errorf("ptr should be nil")
	}
}

func TestServiceInit(t *testing.T) {
	s := NewService("local.", "_ssh._tcp.", "GoTest", 4423)

	if s.ptr == nil {
		t.Errorf("ptr should not be nil\n")
	}

	name := s.Name()
	if name != "GoTest" {
		t.Errorf("name is wrong")
	}

	domain := s.Domain()
	if domain != "local." {
		t.Errorf("domain is wrong")
	}

	type_ := s.Type()
	if type_ != "_ssh._tcp." {
		t.Errorf("type is wrong")
	}

	host := s.HostName()
	if host != "" {
		t.Errorf("hostName is wrong")
	}

	port := s.Port()
	if port != 4423 {
		t.Errorf("port is wrong: %d != 4423", port)
	}

	s.Free()

	if s.ptr != nil {
		t.Errorf("ptr should be nil")
	}
}

func TestBrowserError(t *testing.T) {
	browser := NewBrowser()
	browser.Search("", "")

	go func() {
		switch event := (<-browser.Event).(type) {
		case *DidNotSearchEvent:
			if event.ErrCode != -72004 { // NSNetServicesBadArgumentError
				t.Errorf("error is wrong")
			}
		default:
			t.Errorf("unexpected event")
		}
		eventloop.Stop()
	}()

	eventloop.Run()
}

func TestBasic(t *testing.T) {
	service := NewService("local.", "_ssh._tcp.", "GoTest", 4423)
	service.Publish()

	browser := NewBrowser()
	browser.Search("_ssh._tcp.", "")

	timeout := time.After(30 * time.Second)

	go func() {
	LOOP:
		for {
			select {
			case <-timeout:
				t.Errorf("search time out")
				break LOOP
			case event := <-browser.Event:
				switch e := event.(type) {
				case *FindServiceEvent:
					if e.Service.Name() == "GoTest" {
						service.Stop()
					}
				case *RemoveServiceEvent:
					if e.Service.Name() == "GoTest" {
						// removed
						break LOOP
					}
				default:
					// ignore
				}
			}
		}
		eventloop.Stop()
	}()
	eventloop.Run()
}
