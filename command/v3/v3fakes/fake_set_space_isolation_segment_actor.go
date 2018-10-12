// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeSetSpaceIsolationSegmentActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	AssignIsolationSegmentToSpaceByNameAndSpaceStub        func(isolationSegmentName string, spaceGUID string) (v7action.Warnings, error)
	assignIsolationSegmentToSpaceByNameAndSpaceMutex       sync.RWMutex
	assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall []struct {
		isolationSegmentName string
		spaceGUID            string
	}
	assignIsolationSegmentToSpaceByNameAndSpaceReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetSpaceIsolationSegmentActor) CloudControllerAPIVersion() string {
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

func (fake *FakeSetSpaceIsolationSegmentActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeSetSpaceIsolationSegmentActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSetSpaceIsolationSegmentActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpace(isolationSegmentName string, spaceGUID string) (v7action.Warnings, error) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall[len(fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall)]
	fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall = append(fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall, struct {
		isolationSegmentName string
		spaceGUID            string
	}{isolationSegmentName, spaceGUID})
	fake.recordInvocation("AssignIsolationSegmentToSpaceByNameAndSpace", []interface{}{isolationSegmentName, spaceGUID})
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Unlock()
	if fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub != nil {
		return fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub(isolationSegmentName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.assignIsolationSegmentToSpaceByNameAndSpaceReturns.result1, fake.assignIsolationSegmentToSpaceByNameAndSpaceReturns.result2
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceCallCount() int {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RLock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RUnlock()
	return len(fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall)
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RLock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RUnlock()
	return fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall[i].isolationSegmentName, fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceReturns(result1 v7action.Warnings, result2 error) {
	fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub = nil
	fake.assignIsolationSegmentToSpaceByNameAndSpaceReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub = nil
	if fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall == nil {
		fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetSpaceIsolationSegmentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RLock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetSpaceIsolationSegmentActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.SetSpaceIsolationSegmentActor = new(FakeSetSpaceIsolationSegmentActor)
