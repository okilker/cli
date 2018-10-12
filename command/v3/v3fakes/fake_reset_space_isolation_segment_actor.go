// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeResetSpaceIsolationSegmentActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	ResetSpaceIsolationSegmentStub        func(orgGUID string, spaceGUID string) (string, v7action.Warnings, error)
	resetSpaceIsolationSegmentMutex       sync.RWMutex
	resetSpaceIsolationSegmentArgsForCall []struct {
		orgGUID   string
		spaceGUID string
	}
	resetSpaceIsolationSegmentReturns struct {
		result1 string
		result2 v7action.Warnings
		result3 error
	}
	resetSpaceIsolationSegmentReturnsOnCall map[int]struct {
		result1 string
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResetSpaceIsolationSegmentActor) CloudControllerAPIVersion() string {
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

func (fake *FakeResetSpaceIsolationSegmentActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeResetSpaceIsolationSegmentActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResetSpaceIsolationSegmentActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeResetSpaceIsolationSegmentActor) ResetSpaceIsolationSegment(orgGUID string, spaceGUID string) (string, v7action.Warnings, error) {
	fake.resetSpaceIsolationSegmentMutex.Lock()
	ret, specificReturn := fake.resetSpaceIsolationSegmentReturnsOnCall[len(fake.resetSpaceIsolationSegmentArgsForCall)]
	fake.resetSpaceIsolationSegmentArgsForCall = append(fake.resetSpaceIsolationSegmentArgsForCall, struct {
		orgGUID   string
		spaceGUID string
	}{orgGUID, spaceGUID})
	fake.recordInvocation("ResetSpaceIsolationSegment", []interface{}{orgGUID, spaceGUID})
	fake.resetSpaceIsolationSegmentMutex.Unlock()
	if fake.ResetSpaceIsolationSegmentStub != nil {
		return fake.ResetSpaceIsolationSegmentStub(orgGUID, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.resetSpaceIsolationSegmentReturns.result1, fake.resetSpaceIsolationSegmentReturns.result2, fake.resetSpaceIsolationSegmentReturns.result3
}

func (fake *FakeResetSpaceIsolationSegmentActor) ResetSpaceIsolationSegmentCallCount() int {
	fake.resetSpaceIsolationSegmentMutex.RLock()
	defer fake.resetSpaceIsolationSegmentMutex.RUnlock()
	return len(fake.resetSpaceIsolationSegmentArgsForCall)
}

func (fake *FakeResetSpaceIsolationSegmentActor) ResetSpaceIsolationSegmentArgsForCall(i int) (string, string) {
	fake.resetSpaceIsolationSegmentMutex.RLock()
	defer fake.resetSpaceIsolationSegmentMutex.RUnlock()
	return fake.resetSpaceIsolationSegmentArgsForCall[i].orgGUID, fake.resetSpaceIsolationSegmentArgsForCall[i].spaceGUID
}

func (fake *FakeResetSpaceIsolationSegmentActor) ResetSpaceIsolationSegmentReturns(result1 string, result2 v7action.Warnings, result3 error) {
	fake.ResetSpaceIsolationSegmentStub = nil
	fake.resetSpaceIsolationSegmentReturns = struct {
		result1 string
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResetSpaceIsolationSegmentActor) ResetSpaceIsolationSegmentReturnsOnCall(i int, result1 string, result2 v7action.Warnings, result3 error) {
	fake.ResetSpaceIsolationSegmentStub = nil
	if fake.resetSpaceIsolationSegmentReturnsOnCall == nil {
		fake.resetSpaceIsolationSegmentReturnsOnCall = make(map[int]struct {
			result1 string
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.resetSpaceIsolationSegmentReturnsOnCall[i] = struct {
		result1 string
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResetSpaceIsolationSegmentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.resetSpaceIsolationSegmentMutex.RLock()
	defer fake.resetSpaceIsolationSegmentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResetSpaceIsolationSegmentActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.ResetSpaceIsolationSegmentActor = new(FakeResetSpaceIsolationSegmentActor)
