// This file was generated by counterfeiter
package execfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/exec"
)

type FakeTaskConfigSource struct {
	FetchConfigStub        func(*exec.SourceRepository) (atc.TaskConfig, error)
	fetchConfigMutex       sync.RWMutex
	fetchConfigArgsForCall []struct {
		arg1 *exec.SourceRepository
	}
	fetchConfigReturns struct {
		result1 atc.TaskConfig
		result2 error
	}
	fetchConfigReturnsOnCall map[int]struct {
		result1 atc.TaskConfig
		result2 error
	}
	WarningsStub        func() []string
	warningsMutex       sync.RWMutex
	warningsArgsForCall []struct{}
	warningsReturns     struct {
		result1 []string
	}
	warningsReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskConfigSource) FetchConfig(arg1 *exec.SourceRepository) (atc.TaskConfig, error) {
	fake.fetchConfigMutex.Lock()
	ret, specificReturn := fake.fetchConfigReturnsOnCall[len(fake.fetchConfigArgsForCall)]
	fake.fetchConfigArgsForCall = append(fake.fetchConfigArgsForCall, struct {
		arg1 *exec.SourceRepository
	}{arg1})
	fake.recordInvocation("FetchConfig", []interface{}{arg1})
	fake.fetchConfigMutex.Unlock()
	if fake.FetchConfigStub != nil {
		return fake.FetchConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchConfigReturns.result1, fake.fetchConfigReturns.result2
}

func (fake *FakeTaskConfigSource) FetchConfigCallCount() int {
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	return len(fake.fetchConfigArgsForCall)
}

func (fake *FakeTaskConfigSource) FetchConfigArgsForCall(i int) *exec.SourceRepository {
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	return fake.fetchConfigArgsForCall[i].arg1
}

func (fake *FakeTaskConfigSource) FetchConfigReturns(result1 atc.TaskConfig, result2 error) {
	fake.FetchConfigStub = nil
	fake.fetchConfigReturns = struct {
		result1 atc.TaskConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskConfigSource) FetchConfigReturnsOnCall(i int, result1 atc.TaskConfig, result2 error) {
	fake.FetchConfigStub = nil
	if fake.fetchConfigReturnsOnCall == nil {
		fake.fetchConfigReturnsOnCall = make(map[int]struct {
			result1 atc.TaskConfig
			result2 error
		})
	}
	fake.fetchConfigReturnsOnCall[i] = struct {
		result1 atc.TaskConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskConfigSource) Warnings() []string {
	fake.warningsMutex.Lock()
	ret, specificReturn := fake.warningsReturnsOnCall[len(fake.warningsArgsForCall)]
	fake.warningsArgsForCall = append(fake.warningsArgsForCall, struct{}{})
	fake.recordInvocation("Warnings", []interface{}{})
	fake.warningsMutex.Unlock()
	if fake.WarningsStub != nil {
		return fake.WarningsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.warningsReturns.result1
}

func (fake *FakeTaskConfigSource) WarningsCallCount() int {
	fake.warningsMutex.RLock()
	defer fake.warningsMutex.RUnlock()
	return len(fake.warningsArgsForCall)
}

func (fake *FakeTaskConfigSource) WarningsReturns(result1 []string) {
	fake.WarningsStub = nil
	fake.warningsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeTaskConfigSource) WarningsReturnsOnCall(i int, result1 []string) {
	fake.WarningsStub = nil
	if fake.warningsReturnsOnCall == nil {
		fake.warningsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.warningsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeTaskConfigSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	fake.warningsMutex.RLock()
	defer fake.warningsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTaskConfigSource) recordInvocation(key string, args []interface{}) {
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

var _ exec.TaskConfigSource = new(FakeTaskConfigSource)
