package bonjour

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#include "service.h"
*/
import "C"
import "unsafe"

import (
	"github.com/typester/go-cocoa-eventloop"
	"runtime"
	"sync"
)

type Service struct {
	mu  sync.Mutex
	ptr *_Ctype_struct_BonjourService
}

func NewService(domain, _type, name string, port int) *Service {
	s := &Service{}
	s.ptr = C.BonjourServiceNew(
		unsafe.Pointer(s),
		C.CString(domain),
		C.CString(_type),
		C.CString(name),
		C.int(port),
	)
	runtime.SetFinalizer(s, func(s *Service) { s.Free() })
	return s
}

func newServiceFromPtr(ptr unsafe.Pointer) *Service {
	s := &Service{}
	s.ptr = C.BonjourServiceNewFromPtr(unsafe.Pointer(s), ptr)
	runtime.SetFinalizer(s, func(s *Service) { s.Free() })
	return s
}

func (s *Service) Free() {
	s.mu.Lock()
	if s.ptr != nil {
		C.BonjourServiceFree(s.ptr)
		s.ptr = nil
	}
	s.mu.Unlock()
}

func (s *Service) Name() string {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	name := C.GoString(C.BonjourServiceGetName(s.ptr))

	s.mu.Unlock()

	return name
}

func (s *Service) Type() string {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	_type := C.GoString(C.BonjourServiceGetType(s.ptr))

	s.mu.Unlock()

	return _type
}

func (s *Service) Domain() string {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	domain := C.GoString(C.BonjourServiceGetDomain(s.ptr))

	s.mu.Unlock()

	return domain
}

func (s *Service) HostName() string {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	name := C.GoString(C.BonjourServiceGetHostName(s.ptr))

	s.mu.Unlock()

	return name
}

func (s *Service) Port() int {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	port := int(C.BonjourServiceGetPort(s.ptr))

	s.mu.Unlock()

	return port
}

func (s *Service) Publish() {
	go s.doPublish()
}

func (s *Service) doPublish() {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	eventloop.Do(func() {
		C.BonjourServicePublish(s.ptr)
	})

	s.mu.Unlock()
}

func (s *Service) Stop() {
	go s.doStop()
}

func (s *Service) doStop() {
	s.mu.Lock()
	if s.ptr == nil {
		panic("object is already deallocated")
	}

	eventloop.Do(func() {
		C.BonjourServiceStop(s.ptr)
	})

	s.mu.Unlock()
}
