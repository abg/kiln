// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pivotal-cf/kiln/internal/providers"
)

type S3Provider struct {
	GetS3ClientStub        func(string, string, string) s3iface.S3API
	getS3ClientMutex       sync.RWMutex
	getS3ClientArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getS3ClientReturns struct {
		result1 s3iface.S3API
	}
	getS3ClientReturnsOnCall map[int]struct {
		result1 s3iface.S3API
	}
	GetS3DownloaderStub        func(string, string, string) providers.S3Downloader
	getS3DownloaderMutex       sync.RWMutex
	getS3DownloaderArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getS3DownloaderReturns struct {
		result1 providers.S3Downloader
	}
	getS3DownloaderReturnsOnCall map[int]struct {
		result1 providers.S3Downloader
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *S3Provider) GetS3Client(arg1 string, arg2 string, arg3 string) s3iface.S3API {
	fake.getS3ClientMutex.Lock()
	ret, specificReturn := fake.getS3ClientReturnsOnCall[len(fake.getS3ClientArgsForCall)]
	fake.getS3ClientArgsForCall = append(fake.getS3ClientArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetS3Client", []interface{}{arg1, arg2, arg3})
	fake.getS3ClientMutex.Unlock()
	if fake.GetS3ClientStub != nil {
		return fake.GetS3ClientStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getS3ClientReturns
	return fakeReturns.result1
}

func (fake *S3Provider) GetS3ClientCallCount() int {
	fake.getS3ClientMutex.RLock()
	defer fake.getS3ClientMutex.RUnlock()
	return len(fake.getS3ClientArgsForCall)
}

func (fake *S3Provider) GetS3ClientCalls(stub func(string, string, string) s3iface.S3API) {
	fake.getS3ClientMutex.Lock()
	defer fake.getS3ClientMutex.Unlock()
	fake.GetS3ClientStub = stub
}

func (fake *S3Provider) GetS3ClientArgsForCall(i int) (string, string, string) {
	fake.getS3ClientMutex.RLock()
	defer fake.getS3ClientMutex.RUnlock()
	argsForCall := fake.getS3ClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *S3Provider) GetS3ClientReturns(result1 s3iface.S3API) {
	fake.getS3ClientMutex.Lock()
	defer fake.getS3ClientMutex.Unlock()
	fake.GetS3ClientStub = nil
	fake.getS3ClientReturns = struct {
		result1 s3iface.S3API
	}{result1}
}

func (fake *S3Provider) GetS3ClientReturnsOnCall(i int, result1 s3iface.S3API) {
	fake.getS3ClientMutex.Lock()
	defer fake.getS3ClientMutex.Unlock()
	fake.GetS3ClientStub = nil
	if fake.getS3ClientReturnsOnCall == nil {
		fake.getS3ClientReturnsOnCall = make(map[int]struct {
			result1 s3iface.S3API
		})
	}
	fake.getS3ClientReturnsOnCall[i] = struct {
		result1 s3iface.S3API
	}{result1}
}

func (fake *S3Provider) GetS3Downloader(arg1 string, arg2 string, arg3 string) providers.S3Downloader {
	fake.getS3DownloaderMutex.Lock()
	ret, specificReturn := fake.getS3DownloaderReturnsOnCall[len(fake.getS3DownloaderArgsForCall)]
	fake.getS3DownloaderArgsForCall = append(fake.getS3DownloaderArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetS3Downloader", []interface{}{arg1, arg2, arg3})
	fake.getS3DownloaderMutex.Unlock()
	if fake.GetS3DownloaderStub != nil {
		return fake.GetS3DownloaderStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getS3DownloaderReturns
	return fakeReturns.result1
}

func (fake *S3Provider) GetS3DownloaderCallCount() int {
	fake.getS3DownloaderMutex.RLock()
	defer fake.getS3DownloaderMutex.RUnlock()
	return len(fake.getS3DownloaderArgsForCall)
}

func (fake *S3Provider) GetS3DownloaderCalls(stub func(string, string, string) providers.S3Downloader) {
	fake.getS3DownloaderMutex.Lock()
	defer fake.getS3DownloaderMutex.Unlock()
	fake.GetS3DownloaderStub = stub
}

func (fake *S3Provider) GetS3DownloaderArgsForCall(i int) (string, string, string) {
	fake.getS3DownloaderMutex.RLock()
	defer fake.getS3DownloaderMutex.RUnlock()
	argsForCall := fake.getS3DownloaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *S3Provider) GetS3DownloaderReturns(result1 providers.S3Downloader) {
	fake.getS3DownloaderMutex.Lock()
	defer fake.getS3DownloaderMutex.Unlock()
	fake.GetS3DownloaderStub = nil
	fake.getS3DownloaderReturns = struct {
		result1 providers.S3Downloader
	}{result1}
}

func (fake *S3Provider) GetS3DownloaderReturnsOnCall(i int, result1 providers.S3Downloader) {
	fake.getS3DownloaderMutex.Lock()
	defer fake.getS3DownloaderMutex.Unlock()
	fake.GetS3DownloaderStub = nil
	if fake.getS3DownloaderReturnsOnCall == nil {
		fake.getS3DownloaderReturnsOnCall = make(map[int]struct {
			result1 providers.S3Downloader
		})
	}
	fake.getS3DownloaderReturnsOnCall[i] = struct {
		result1 providers.S3Downloader
	}{result1}
}

func (fake *S3Provider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getS3ClientMutex.RLock()
	defer fake.getS3ClientMutex.RUnlock()
	fake.getS3DownloaderMutex.RLock()
	defer fake.getS3DownloaderMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *S3Provider) recordInvocation(key string, args []interface{}) {
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