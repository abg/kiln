// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/kiln/commands"
)

type PivnetProductFilesService struct {
	AddToReleaseStub        func(string, int, int) error
	addToReleaseMutex       sync.RWMutex
	addToReleaseArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	addToReleaseReturns struct {
		result1 error
	}
	addToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(string) ([]pivnet.ProductFile, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 string
	}
	listReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PivnetProductFilesService) AddToRelease(arg1 string, arg2 int, arg3 int) error {
	fake.addToReleaseMutex.Lock()
	ret, specificReturn := fake.addToReleaseReturnsOnCall[len(fake.addToReleaseArgsForCall)]
	fake.addToReleaseArgsForCall = append(fake.addToReleaseArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddToRelease", []interface{}{arg1, arg2, arg3})
	fake.addToReleaseMutex.Unlock()
	if fake.AddToReleaseStub != nil {
		return fake.AddToReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addToReleaseReturns
	return fakeReturns.result1
}

func (fake *PivnetProductFilesService) AddToReleaseCallCount() int {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return len(fake.addToReleaseArgsForCall)
}

func (fake *PivnetProductFilesService) AddToReleaseCalls(stub func(string, int, int) error) {
	fake.addToReleaseMutex.Lock()
	defer fake.addToReleaseMutex.Unlock()
	fake.AddToReleaseStub = stub
}

func (fake *PivnetProductFilesService) AddToReleaseArgsForCall(i int) (string, int, int) {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	argsForCall := fake.addToReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PivnetProductFilesService) AddToReleaseReturns(result1 error) {
	fake.addToReleaseMutex.Lock()
	defer fake.addToReleaseMutex.Unlock()
	fake.AddToReleaseStub = nil
	fake.addToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *PivnetProductFilesService) AddToReleaseReturnsOnCall(i int, result1 error) {
	fake.addToReleaseMutex.Lock()
	defer fake.addToReleaseMutex.Unlock()
	fake.AddToReleaseStub = nil
	if fake.addToReleaseReturnsOnCall == nil {
		fake.addToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PivnetProductFilesService) List(arg1 string) ([]pivnet.ProductFile, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetProductFilesService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *PivnetProductFilesService) ListCalls(stub func(string) ([]pivnet.ProductFile, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *PivnetProductFilesService) ListArgsForCall(i int) string {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PivnetProductFilesService) ListReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *PivnetProductFilesService) ListReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *PivnetProductFilesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PivnetProductFilesService) recordInvocation(key string, args []interface{}) {
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

var _ commands.PivnetProductFilesService = new(PivnetProductFilesService)