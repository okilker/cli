// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeDeleteActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	DeleteApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v7action.Warnings, error)
	deleteApplicationByNameAndSpaceMutex       sync.RWMutex
	deleteApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	deleteApplicationByNameAndSpaceReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteActor) CloudControllerAPIVersion() string {
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

func (fake *FakeDeleteActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeDeleteActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeDeleteActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeDeleteActor) DeleteApplicationByNameAndSpace(name string, spaceGUID string) (v7action.Warnings, error) {
	fake.deleteApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.deleteApplicationByNameAndSpaceReturnsOnCall[len(fake.deleteApplicationByNameAndSpaceArgsForCall)]
	fake.deleteApplicationByNameAndSpaceArgsForCall = append(fake.deleteApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("DeleteApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.deleteApplicationByNameAndSpaceMutex.Unlock()
	if fake.DeleteApplicationByNameAndSpaceStub != nil {
		return fake.DeleteApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteApplicationByNameAndSpaceReturns.result1, fake.deleteApplicationByNameAndSpaceReturns.result2
}

func (fake *FakeDeleteActor) DeleteApplicationByNameAndSpaceCallCount() int {
	fake.deleteApplicationByNameAndSpaceMutex.RLock()
	defer fake.deleteApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.deleteApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeDeleteActor) DeleteApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.deleteApplicationByNameAndSpaceMutex.RLock()
	defer fake.deleteApplicationByNameAndSpaceMutex.RUnlock()
	return fake.deleteApplicationByNameAndSpaceArgsForCall[i].name, fake.deleteApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeDeleteActor) DeleteApplicationByNameAndSpaceReturns(result1 v7action.Warnings, result2 error) {
	fake.DeleteApplicationByNameAndSpaceStub = nil
	fake.deleteApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteActor) DeleteApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.DeleteApplicationByNameAndSpaceStub = nil
	if fake.deleteApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.deleteApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.deleteApplicationByNameAndSpaceMutex.RLock()
	defer fake.deleteApplicationByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.DeleteActor = new(FakeDeleteActor)
