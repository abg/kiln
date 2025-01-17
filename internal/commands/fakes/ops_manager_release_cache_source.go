// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/internal/commands"
	"github.com/pivotal-cf/om/api"
)

type OpsManagerReleaseCacheSource struct {
	GetBoshEnvironmentStub        func() (api.GetBoshEnvironmentOutput, error)
	getBoshEnvironmentMutex       sync.RWMutex
	getBoshEnvironmentArgsForCall []struct {
	}
	getBoshEnvironmentReturns struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}
	getBoshEnvironmentReturnsOnCall map[int]struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}
	GetSecurityRootCACertificateStub        func() (string, error)
	getSecurityRootCACertificateMutex       sync.RWMutex
	getSecurityRootCACertificateArgsForCall []struct {
	}
	getSecurityRootCACertificateReturns struct {
		result1 string
		result2 error
	}
	getSecurityRootCACertificateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetStagedProductByNameStub        func(string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		arg1 string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	GetStagedProductManifestStub        func(string) (string, error)
	getStagedProductManifestMutex       sync.RWMutex
	getStagedProductManifestArgsForCall []struct {
		arg1 string
	}
	getStagedProductManifestReturns struct {
		result1 string
		result2 error
	}
	getStagedProductManifestReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OpsManagerReleaseCacheSource) GetBoshEnvironment() (api.GetBoshEnvironmentOutput, error) {
	fake.getBoshEnvironmentMutex.Lock()
	ret, specificReturn := fake.getBoshEnvironmentReturnsOnCall[len(fake.getBoshEnvironmentArgsForCall)]
	fake.getBoshEnvironmentArgsForCall = append(fake.getBoshEnvironmentArgsForCall, struct {
	}{})
	stub := fake.GetBoshEnvironmentStub
	fakeReturns := fake.getBoshEnvironmentReturns
	fake.recordInvocation("GetBoshEnvironment", []interface{}{})
	fake.getBoshEnvironmentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OpsManagerReleaseCacheSource) GetBoshEnvironmentCallCount() int {
	fake.getBoshEnvironmentMutex.RLock()
	defer fake.getBoshEnvironmentMutex.RUnlock()
	return len(fake.getBoshEnvironmentArgsForCall)
}

func (fake *OpsManagerReleaseCacheSource) GetBoshEnvironmentCalls(stub func() (api.GetBoshEnvironmentOutput, error)) {
	fake.getBoshEnvironmentMutex.Lock()
	defer fake.getBoshEnvironmentMutex.Unlock()
	fake.GetBoshEnvironmentStub = stub
}

func (fake *OpsManagerReleaseCacheSource) GetBoshEnvironmentReturns(result1 api.GetBoshEnvironmentOutput, result2 error) {
	fake.getBoshEnvironmentMutex.Lock()
	defer fake.getBoshEnvironmentMutex.Unlock()
	fake.GetBoshEnvironmentStub = nil
	fake.getBoshEnvironmentReturns = struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetBoshEnvironmentReturnsOnCall(i int, result1 api.GetBoshEnvironmentOutput, result2 error) {
	fake.getBoshEnvironmentMutex.Lock()
	defer fake.getBoshEnvironmentMutex.Unlock()
	fake.GetBoshEnvironmentStub = nil
	if fake.getBoshEnvironmentReturnsOnCall == nil {
		fake.getBoshEnvironmentReturnsOnCall = make(map[int]struct {
			result1 api.GetBoshEnvironmentOutput
			result2 error
		})
	}
	fake.getBoshEnvironmentReturnsOnCall[i] = struct {
		result1 api.GetBoshEnvironmentOutput
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetSecurityRootCACertificate() (string, error) {
	fake.getSecurityRootCACertificateMutex.Lock()
	ret, specificReturn := fake.getSecurityRootCACertificateReturnsOnCall[len(fake.getSecurityRootCACertificateArgsForCall)]
	fake.getSecurityRootCACertificateArgsForCall = append(fake.getSecurityRootCACertificateArgsForCall, struct {
	}{})
	stub := fake.GetSecurityRootCACertificateStub
	fakeReturns := fake.getSecurityRootCACertificateReturns
	fake.recordInvocation("GetSecurityRootCACertificate", []interface{}{})
	fake.getSecurityRootCACertificateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OpsManagerReleaseCacheSource) GetSecurityRootCACertificateCallCount() int {
	fake.getSecurityRootCACertificateMutex.RLock()
	defer fake.getSecurityRootCACertificateMutex.RUnlock()
	return len(fake.getSecurityRootCACertificateArgsForCall)
}

func (fake *OpsManagerReleaseCacheSource) GetSecurityRootCACertificateCalls(stub func() (string, error)) {
	fake.getSecurityRootCACertificateMutex.Lock()
	defer fake.getSecurityRootCACertificateMutex.Unlock()
	fake.GetSecurityRootCACertificateStub = stub
}

func (fake *OpsManagerReleaseCacheSource) GetSecurityRootCACertificateReturns(result1 string, result2 error) {
	fake.getSecurityRootCACertificateMutex.Lock()
	defer fake.getSecurityRootCACertificateMutex.Unlock()
	fake.GetSecurityRootCACertificateStub = nil
	fake.getSecurityRootCACertificateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetSecurityRootCACertificateReturnsOnCall(i int, result1 string, result2 error) {
	fake.getSecurityRootCACertificateMutex.Lock()
	defer fake.getSecurityRootCACertificateMutex.Unlock()
	fake.GetSecurityRootCACertificateStub = nil
	if fake.getSecurityRootCACertificateReturnsOnCall == nil {
		fake.getSecurityRootCACertificateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getSecurityRootCACertificateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductByName(arg1 string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStagedProductByNameStub
	fakeReturns := fake.getStagedProductByNameReturns
	fake.recordInvocation("GetStagedProductByName", []interface{}{arg1})
	fake.getStagedProductByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductByNameCalls(stub func(string) (api.StagedProductsFindOutput, error)) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = stub
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	argsForCall := fake.getStagedProductByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductManifest(arg1 string) (string, error) {
	fake.getStagedProductManifestMutex.Lock()
	ret, specificReturn := fake.getStagedProductManifestReturnsOnCall[len(fake.getStagedProductManifestArgsForCall)]
	fake.getStagedProductManifestArgsForCall = append(fake.getStagedProductManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStagedProductManifestStub
	fakeReturns := fake.getStagedProductManifestReturns
	fake.recordInvocation("GetStagedProductManifest", []interface{}{arg1})
	fake.getStagedProductManifestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductManifestCallCount() int {
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	return len(fake.getStagedProductManifestArgsForCall)
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductManifestCalls(stub func(string) (string, error)) {
	fake.getStagedProductManifestMutex.Lock()
	defer fake.getStagedProductManifestMutex.Unlock()
	fake.GetStagedProductManifestStub = stub
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductManifestArgsForCall(i int) string {
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	argsForCall := fake.getStagedProductManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductManifestReturns(result1 string, result2 error) {
	fake.getStagedProductManifestMutex.Lock()
	defer fake.getStagedProductManifestMutex.Unlock()
	fake.GetStagedProductManifestStub = nil
	fake.getStagedProductManifestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) GetStagedProductManifestReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStagedProductManifestMutex.Lock()
	defer fake.getStagedProductManifestMutex.Unlock()
	fake.GetStagedProductManifestStub = nil
	if fake.getStagedProductManifestReturnsOnCall == nil {
		fake.getStagedProductManifestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStagedProductManifestReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OpsManagerReleaseCacheSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBoshEnvironmentMutex.RLock()
	defer fake.getBoshEnvironmentMutex.RUnlock()
	fake.getSecurityRootCACertificateMutex.RLock()
	defer fake.getSecurityRootCACertificateMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OpsManagerReleaseCacheSource) recordInvocation(key string, args []interface{}) {
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

var _ commands.OpsManagerReleaseCacheSource = new(OpsManagerReleaseCacheSource)
