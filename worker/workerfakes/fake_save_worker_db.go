// This file was generated by counterfeiter
package workerfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeSaveWorkerDB struct {
	SaveWorkerStub        func(db.WorkerInfo, time.Duration) (db.SavedWorker, error)
	saveWorkerMutex       sync.RWMutex
	saveWorkerArgsForCall []struct {
		arg1 db.WorkerInfo
		arg2 time.Duration
	}
	saveWorkerReturns struct {
		result1 db.SavedWorker
		result2 error
	}
	saveWorkerReturnsOnCall map[int]struct {
		result1 db.SavedWorker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSaveWorkerDB) SaveWorker(arg1 db.WorkerInfo, arg2 time.Duration) (db.SavedWorker, error) {
	fake.saveWorkerMutex.Lock()
	ret, specificReturn := fake.saveWorkerReturnsOnCall[len(fake.saveWorkerArgsForCall)]
	fake.saveWorkerArgsForCall = append(fake.saveWorkerArgsForCall, struct {
		arg1 db.WorkerInfo
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("SaveWorker", []interface{}{arg1, arg2})
	fake.saveWorkerMutex.Unlock()
	if fake.SaveWorkerStub != nil {
		return fake.SaveWorkerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.saveWorkerReturns.result1, fake.saveWorkerReturns.result2
}

func (fake *FakeSaveWorkerDB) SaveWorkerCallCount() int {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return len(fake.saveWorkerArgsForCall)
}

func (fake *FakeSaveWorkerDB) SaveWorkerArgsForCall(i int) (db.WorkerInfo, time.Duration) {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return fake.saveWorkerArgsForCall[i].arg1, fake.saveWorkerArgsForCall[i].arg2
}

func (fake *FakeSaveWorkerDB) SaveWorkerReturns(result1 db.SavedWorker, result2 error) {
	fake.SaveWorkerStub = nil
	fake.saveWorkerReturns = struct {
		result1 db.SavedWorker
		result2 error
	}{result1, result2}
}

func (fake *FakeSaveWorkerDB) SaveWorkerReturnsOnCall(i int, result1 db.SavedWorker, result2 error) {
	fake.SaveWorkerStub = nil
	if fake.saveWorkerReturnsOnCall == nil {
		fake.saveWorkerReturnsOnCall = make(map[int]struct {
			result1 db.SavedWorker
			result2 error
		})
	}
	fake.saveWorkerReturnsOnCall[i] = struct {
		result1 db.SavedWorker
		result2 error
	}{result1, result2}
}

func (fake *FakeSaveWorkerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSaveWorkerDB) recordInvocation(key string, args []interface{}) {
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

var _ worker.SaveWorkerDB = new(FakeSaveWorkerDB)
