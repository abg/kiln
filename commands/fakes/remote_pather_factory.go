// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/commands"
	"github.com/pivotal-cf/kiln/fetcher"
	"github.com/pivotal-cf/kiln/internal/cargo"
)

type RemotePatherFactory struct {
	RemotePatherStub        func(string, cargo.Kilnfile) (fetcher.RemotePather, error)
	remotePatherMutex       sync.RWMutex
	remotePatherArgsForCall []struct {
		arg1 string
		arg2 cargo.Kilnfile
	}
	remotePatherReturns struct {
		result1 fetcher.RemotePather
		result2 error
	}
	remotePatherReturnsOnCall map[int]struct {
		result1 fetcher.RemotePather
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RemotePatherFactory) RemotePather(arg1 string, arg2 cargo.Kilnfile) (fetcher.RemotePather, error) {
	fake.remotePatherMutex.Lock()
	ret, specificReturn := fake.remotePatherReturnsOnCall[len(fake.remotePatherArgsForCall)]
	fake.remotePatherArgsForCall = append(fake.remotePatherArgsForCall, struct {
		arg1 string
		arg2 cargo.Kilnfile
	}{arg1, arg2})
	fake.recordInvocation("RemotePather", []interface{}{arg1, arg2})
	fake.remotePatherMutex.Unlock()
	if fake.RemotePatherStub != nil {
		return fake.RemotePatherStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.remotePatherReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *RemotePatherFactory) RemotePatherCallCount() int {
	fake.remotePatherMutex.RLock()
	defer fake.remotePatherMutex.RUnlock()
	return len(fake.remotePatherArgsForCall)
}

func (fake *RemotePatherFactory) RemotePatherCalls(stub func(string, cargo.Kilnfile) (fetcher.RemotePather, error)) {
	fake.remotePatherMutex.Lock()
	defer fake.remotePatherMutex.Unlock()
	fake.RemotePatherStub = stub
}

func (fake *RemotePatherFactory) RemotePatherArgsForCall(i int) (string, cargo.Kilnfile) {
	fake.remotePatherMutex.RLock()
	defer fake.remotePatherMutex.RUnlock()
	argsForCall := fake.remotePatherArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *RemotePatherFactory) RemotePatherReturns(result1 fetcher.RemotePather, result2 error) {
	fake.remotePatherMutex.Lock()
	defer fake.remotePatherMutex.Unlock()
	fake.RemotePatherStub = nil
	fake.remotePatherReturns = struct {
		result1 fetcher.RemotePather
		result2 error
	}{result1, result2}
}

func (fake *RemotePatherFactory) RemotePatherReturnsOnCall(i int, result1 fetcher.RemotePather, result2 error) {
	fake.remotePatherMutex.Lock()
	defer fake.remotePatherMutex.Unlock()
	fake.RemotePatherStub = nil
	if fake.remotePatherReturnsOnCall == nil {
		fake.remotePatherReturnsOnCall = make(map[int]struct {
			result1 fetcher.RemotePather
			result2 error
		})
	}
	fake.remotePatherReturnsOnCall[i] = struct {
		result1 fetcher.RemotePather
		result2 error
	}{result1, result2}
}

func (fake *RemotePatherFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.remotePatherMutex.RLock()
	defer fake.remotePatherMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RemotePatherFactory) recordInvocation(key string, args []interface{}) {
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

var _ commands.RemotePatherFactory = new(RemotePatherFactory)