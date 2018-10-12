// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeSpaceActorV3 struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetEffectiveIsolationSegmentBySpaceStub        func(spaceGUID string, orgDefaultIsolationSegmentGUID string) (v7action.IsolationSegment, v7action.Warnings, error)
	getEffectiveIsolationSegmentBySpaceMutex       sync.RWMutex
	getEffectiveIsolationSegmentBySpaceArgsForCall []struct {
		spaceGUID                      string
		orgDefaultIsolationSegmentGUID string
	}
	getEffectiveIsolationSegmentBySpaceReturns struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}
	getEffectiveIsolationSegmentBySpaceReturnsOnCall map[int]struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersion() string {
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

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpaceActorV3) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpace(spaceGUID string, orgDefaultIsolationSegmentGUID string) (v7action.IsolationSegment, v7action.Warnings, error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	ret, specificReturn := fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall[len(fake.getEffectiveIsolationSegmentBySpaceArgsForCall)]
	fake.getEffectiveIsolationSegmentBySpaceArgsForCall = append(fake.getEffectiveIsolationSegmentBySpaceArgsForCall, struct {
		spaceGUID                      string
		orgDefaultIsolationSegmentGUID string
	}{spaceGUID, orgDefaultIsolationSegmentGUID})
	fake.recordInvocation("GetEffectiveIsolationSegmentBySpace", []interface{}{spaceGUID, orgDefaultIsolationSegmentGUID})
	fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	if fake.GetEffectiveIsolationSegmentBySpaceStub != nil {
		return fake.GetEffectiveIsolationSegmentBySpaceStub(spaceGUID, orgDefaultIsolationSegmentGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getEffectiveIsolationSegmentBySpaceReturns.result1, fake.getEffectiveIsolationSegmentBySpaceReturns.result2, fake.getEffectiveIsolationSegmentBySpaceReturns.result3
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceCallCount() int {
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	return len(fake.getEffectiveIsolationSegmentBySpaceArgsForCall)
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceArgsForCall(i int) (string, string) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	return fake.getEffectiveIsolationSegmentBySpaceArgsForCall[i].spaceGUID, fake.getEffectiveIsolationSegmentBySpaceArgsForCall[i].orgDefaultIsolationSegmentGUID
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceReturns(result1 v7action.IsolationSegment, result2 v7action.Warnings, result3 error) {
	fake.GetEffectiveIsolationSegmentBySpaceStub = nil
	fake.getEffectiveIsolationSegmentBySpaceReturns = struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActorV3) GetEffectiveIsolationSegmentBySpaceReturnsOnCall(i int, result1 v7action.IsolationSegment, result2 v7action.Warnings, result3 error) {
	fake.GetEffectiveIsolationSegmentBySpaceStub = nil
	if fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall == nil {
		fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.IsolationSegment
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall[i] = struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActorV3) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceActorV3) recordInvocation(key string, args []interface{}) {
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

var _ v2.SpaceActorV3 = new(FakeSpaceActorV3)
