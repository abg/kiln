// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"sync"

	"github.com/pivotal-cf/kiln/internal/component"
	"github.com/pivotal-cf/kiln/pkg/cargo"
)

type ReleaseUploader struct {
	GetMatchedReleaseStub        func(cargo.ComponentSpec) (cargo.ComponentLock, error)
	getMatchedReleaseMutex       sync.RWMutex
	getMatchedReleaseArgsForCall []struct {
		arg1 cargo.ComponentSpec
	}
	getMatchedReleaseReturns struct {
		result1 cargo.ComponentLock
		result2 error
	}
	getMatchedReleaseReturnsOnCall map[int]struct {
		result1 cargo.ComponentLock
		result2 error
	}
	UploadReleaseStub        func(cargo.ComponentSpec, io.Reader) (cargo.ComponentLock, error)
	uploadReleaseMutex       sync.RWMutex
	uploadReleaseArgsForCall []struct {
		arg1 cargo.ComponentSpec
		arg2 io.Reader
	}
	uploadReleaseReturns struct {
		result1 cargo.ComponentLock
		result2 error
	}
	uploadReleaseReturnsOnCall map[int]struct {
		result1 cargo.ComponentLock
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseUploader) GetMatchedRelease(arg1 cargo.ComponentSpec) (cargo.ComponentLock, error) {
	fake.getMatchedReleaseMutex.Lock()
	ret, specificReturn := fake.getMatchedReleaseReturnsOnCall[len(fake.getMatchedReleaseArgsForCall)]
	fake.getMatchedReleaseArgsForCall = append(fake.getMatchedReleaseArgsForCall, struct {
		arg1 cargo.ComponentSpec
	}{arg1})
	stub := fake.GetMatchedReleaseStub
	fakeReturns := fake.getMatchedReleaseReturns
	fake.recordInvocation("GetMatchedRelease", []interface{}{arg1})
	fake.getMatchedReleaseMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseUploader) GetMatchedReleaseCallCount() int {
	fake.getMatchedReleaseMutex.RLock()
	defer fake.getMatchedReleaseMutex.RUnlock()
	return len(fake.getMatchedReleaseArgsForCall)
}

func (fake *ReleaseUploader) GetMatchedReleaseCalls(stub func(cargo.ComponentSpec) (cargo.ComponentLock, error)) {
	fake.getMatchedReleaseMutex.Lock()
	defer fake.getMatchedReleaseMutex.Unlock()
	fake.GetMatchedReleaseStub = stub
}

func (fake *ReleaseUploader) GetMatchedReleaseArgsForCall(i int) cargo.ComponentSpec {
	fake.getMatchedReleaseMutex.RLock()
	defer fake.getMatchedReleaseMutex.RUnlock()
	argsForCall := fake.getMatchedReleaseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReleaseUploader) GetMatchedReleaseReturns(result1 cargo.ComponentLock, result2 error) {
	fake.getMatchedReleaseMutex.Lock()
	defer fake.getMatchedReleaseMutex.Unlock()
	fake.GetMatchedReleaseStub = nil
	fake.getMatchedReleaseReturns = struct {
		result1 cargo.ComponentLock
		result2 error
	}{result1, result2}
}

func (fake *ReleaseUploader) GetMatchedReleaseReturnsOnCall(i int, result1 cargo.ComponentLock, result2 error) {
	fake.getMatchedReleaseMutex.Lock()
	defer fake.getMatchedReleaseMutex.Unlock()
	fake.GetMatchedReleaseStub = nil
	if fake.getMatchedReleaseReturnsOnCall == nil {
		fake.getMatchedReleaseReturnsOnCall = make(map[int]struct {
			result1 cargo.ComponentLock
			result2 error
		})
	}
	fake.getMatchedReleaseReturnsOnCall[i] = struct {
		result1 cargo.ComponentLock
		result2 error
	}{result1, result2}
}

func (fake *ReleaseUploader) UploadRelease(arg1 cargo.ComponentSpec, arg2 io.Reader) (cargo.ComponentLock, error) {
	fake.uploadReleaseMutex.Lock()
	ret, specificReturn := fake.uploadReleaseReturnsOnCall[len(fake.uploadReleaseArgsForCall)]
	fake.uploadReleaseArgsForCall = append(fake.uploadReleaseArgsForCall, struct {
		arg1 cargo.ComponentSpec
		arg2 io.Reader
	}{arg1, arg2})
	stub := fake.UploadReleaseStub
	fakeReturns := fake.uploadReleaseReturns
	fake.recordInvocation("UploadRelease", []interface{}{arg1, arg2})
	fake.uploadReleaseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseUploader) UploadReleaseCallCount() int {
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	return len(fake.uploadReleaseArgsForCall)
}

func (fake *ReleaseUploader) UploadReleaseCalls(stub func(cargo.ComponentSpec, io.Reader) (cargo.ComponentLock, error)) {
	fake.uploadReleaseMutex.Lock()
	defer fake.uploadReleaseMutex.Unlock()
	fake.UploadReleaseStub = stub
}

func (fake *ReleaseUploader) UploadReleaseArgsForCall(i int) (cargo.ComponentSpec, io.Reader) {
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	argsForCall := fake.uploadReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ReleaseUploader) UploadReleaseReturns(result1 cargo.ComponentLock, result2 error) {
	fake.uploadReleaseMutex.Lock()
	defer fake.uploadReleaseMutex.Unlock()
	fake.UploadReleaseStub = nil
	fake.uploadReleaseReturns = struct {
		result1 cargo.ComponentLock
		result2 error
	}{result1, result2}
}

func (fake *ReleaseUploader) UploadReleaseReturnsOnCall(i int, result1 cargo.ComponentLock, result2 error) {
	fake.uploadReleaseMutex.Lock()
	defer fake.uploadReleaseMutex.Unlock()
	fake.UploadReleaseStub = nil
	if fake.uploadReleaseReturnsOnCall == nil {
		fake.uploadReleaseReturnsOnCall = make(map[int]struct {
			result1 cargo.ComponentLock
			result2 error
		})
	}
	fake.uploadReleaseReturnsOnCall[i] = struct {
		result1 cargo.ComponentLock
		result2 error
	}{result1, result2}
}

func (fake *ReleaseUploader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMatchedReleaseMutex.RLock()
	defer fake.getMatchedReleaseMutex.RUnlock()
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseUploader) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ component.ReleaseUploader = new(ReleaseUploader)
