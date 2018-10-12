// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeV3PackagesActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetApplicationPackagesStub        func(appName string, spaceGUID string) ([]v7action.Package, v7action.Warnings, error)
	getApplicationPackagesMutex       sync.RWMutex
	getApplicationPackagesArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationPackagesReturns struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	getApplicationPackagesReturnsOnCall map[int]struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3PackagesActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3PackagesActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3PackagesActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3PackagesActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3PackagesActor) GetApplicationPackages(appName string, spaceGUID string) ([]v7action.Package, v7action.Warnings, error) {
	fake.getApplicationPackagesMutex.Lock()
	ret, specificReturn := fake.getApplicationPackagesReturnsOnCall[len(fake.getApplicationPackagesArgsForCall)]
	fake.getApplicationPackagesArgsForCall = append(fake.getApplicationPackagesArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationPackages", []interface{}{appName, spaceGUID})
	fake.getApplicationPackagesMutex.Unlock()
	if fake.GetApplicationPackagesStub != nil {
		return fake.GetApplicationPackagesStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationPackagesReturns.result1, fake.getApplicationPackagesReturns.result2, fake.getApplicationPackagesReturns.result3
}

func (fake *FakeV3PackagesActor) GetApplicationPackagesCallCount() int {
	fake.getApplicationPackagesMutex.RLock()
	defer fake.getApplicationPackagesMutex.RUnlock()
	return len(fake.getApplicationPackagesArgsForCall)
}

func (fake *FakeV3PackagesActor) GetApplicationPackagesArgsForCall(i int) (string, string) {
	fake.getApplicationPackagesMutex.RLock()
	defer fake.getApplicationPackagesMutex.RUnlock()
	return fake.getApplicationPackagesArgsForCall[i].appName, fake.getApplicationPackagesArgsForCall[i].spaceGUID
}

func (fake *FakeV3PackagesActor) GetApplicationPackagesReturns(result1 []v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.GetApplicationPackagesStub = nil
	fake.getApplicationPackagesReturns = struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PackagesActor) GetApplicationPackagesReturnsOnCall(i int, result1 []v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.GetApplicationPackagesStub = nil
	if fake.getApplicationPackagesReturnsOnCall == nil {
		fake.getApplicationPackagesReturnsOnCall = make(map[int]struct {
			result1 []v7action.Package
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationPackagesReturnsOnCall[i] = struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PackagesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getApplicationPackagesMutex.RLock()
	defer fake.getApplicationPackagesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3PackagesActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3PackagesActor = new(FakeV3PackagesActor)
