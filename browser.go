package bonjour
/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#include "browser.h"
*/
import "C"
import "unsafe"

import (
	"sync"
	"runtime"
	"github.com/typester/go-cocoa-eventloop"
)

type Browser struct {
	mu sync.Mutex
	ptr *_Ctype_struct_BonjourBrowser
	Event chan interface{}
}

type FindServiceEvent struct {
	Service *Service
}

type RemoveServiceEvent struct {
	Service *Service
}

type DidNotSearchEvent struct {
	ErrCode int64
}

func NewBrowser() *Browser {
	b := &Browser{}
	b.Event = make(chan interface{})
	b.ptr = C.BonjourBrowserNew(unsafe.Pointer(b))
	runtime.SetFinalizer(b, func (b *Browser) { b.Free() })
	return b
}

func (b *Browser) Free() {
	b.mu.Lock()
	if b.ptr != nil {
		C.BonjourBrowserFree(b.ptr)
		b.ptr = nil
	}
	b.mu.Unlock()
}

func (b *Browser) Search(_type string, domain string) {
	go b.doSearch(_type, domain)
}

func (b *Browser) doSearch(_type string, domain string) {
	b.mu.Lock()
	if b.ptr == nil {
		panic("object is already deallocated")
	}

	eventloop.Do(func () {
		C.BonjourBrowserSearch(b.ptr, C.CString(_type), C.CString(domain))
	});
	
	b.mu.Unlock()
}

func (b *Browser) didFoundService(s *Service) {
	b.Event <- &FindServiceEvent{s}
}

func (b *Browser) didRemoveService(s *Service) {
	b.Event <- &RemoveServiceEvent{s}
}

func (b *Browser) DidNotSearch(errCode int64) {
	b.Event <- &DidNotSearchEvent{errCode}
}

//export bonjourDidFoundService
func bonjourDidFoundService(ptr unsafe.Pointer, servicePtr unsafe.Pointer) {
	b := (*Browser)(ptr)
	service := newServiceFromPtr(servicePtr)
	b.didFoundService(service)
}

//export bonjourDidRemoveService
func bonjourDidRemoveService(ptr unsafe.Pointer, servicePtr unsafe.Pointer) {
	b := (*Browser)(ptr)
	service := newServiceFromPtr(servicePtr)
	b.didRemoveService(service)
}

//export bonjourDidNotSearch
func bonjourDidNotSearch(ptr unsafe.Pointer, code int64) {
	b := (*Browser)(ptr)
	b.DidNotSearch(code)
}







