// This file was generated by counterfeiter
package workerfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
)

type FakeImageFactory struct {
	NewImageStub        func(logger lager.Logger, cancel <-chan os.Signal, imageResource atc.ImageResource, id worker.Identifier, metadata worker.Metadata, tags atc.Tags, teamID int, resourceTypes atc.ResourceTypes, workerClient worker.Client, delegate worker.ImageFetchingDelegate, privileged bool) worker.Image
	newImageMutex       sync.RWMutex
	newImageArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		imageResource atc.ImageResource
		id            worker.Identifier
		metadata      worker.Metadata
		tags          atc.Tags
		teamID        int
		resourceTypes atc.ResourceTypes
		workerClient  worker.Client
		delegate      worker.ImageFetchingDelegate
		privileged    bool
	}
	newImageReturns struct {
		result1 worker.Image
	}
	newImageReturnsOnCall map[int]struct {
		result1 worker.Image
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFactory) NewImage(logger lager.Logger, cancel <-chan os.Signal, imageResource atc.ImageResource, id worker.Identifier, metadata worker.Metadata, tags atc.Tags, teamID int, resourceTypes atc.ResourceTypes, workerClient worker.Client, delegate worker.ImageFetchingDelegate, privileged bool) worker.Image {
	fake.newImageMutex.Lock()
	ret, specificReturn := fake.newImageReturnsOnCall[len(fake.newImageArgsForCall)]
	fake.newImageArgsForCall = append(fake.newImageArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		imageResource atc.ImageResource
		id            worker.Identifier
		metadata      worker.Metadata
		tags          atc.Tags
		teamID        int
		resourceTypes atc.ResourceTypes
		workerClient  worker.Client
		delegate      worker.ImageFetchingDelegate
		privileged    bool
	}{logger, cancel, imageResource, id, metadata, tags, teamID, resourceTypes, workerClient, delegate, privileged})
	fake.recordInvocation("NewImage", []interface{}{logger, cancel, imageResource, id, metadata, tags, teamID, resourceTypes, workerClient, delegate, privileged})
	fake.newImageMutex.Unlock()
	if fake.NewImageStub != nil {
		return fake.NewImageStub(logger, cancel, imageResource, id, metadata, tags, teamID, resourceTypes, workerClient, delegate, privileged)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newImageReturns.result1
}

func (fake *FakeImageFactory) NewImageCallCount() int {
	fake.newImageMutex.RLock()
	defer fake.newImageMutex.RUnlock()
	return len(fake.newImageArgsForCall)
}

func (fake *FakeImageFactory) NewImageArgsForCall(i int) (lager.Logger, <-chan os.Signal, atc.ImageResource, worker.Identifier, worker.Metadata, atc.Tags, int, atc.ResourceTypes, worker.Client, worker.ImageFetchingDelegate, bool) {
	fake.newImageMutex.RLock()
	defer fake.newImageMutex.RUnlock()
	return fake.newImageArgsForCall[i].logger, fake.newImageArgsForCall[i].cancel, fake.newImageArgsForCall[i].imageResource, fake.newImageArgsForCall[i].id, fake.newImageArgsForCall[i].metadata, fake.newImageArgsForCall[i].tags, fake.newImageArgsForCall[i].teamID, fake.newImageArgsForCall[i].resourceTypes, fake.newImageArgsForCall[i].workerClient, fake.newImageArgsForCall[i].delegate, fake.newImageArgsForCall[i].privileged
}

func (fake *FakeImageFactory) NewImageReturns(result1 worker.Image) {
	fake.NewImageStub = nil
	fake.newImageReturns = struct {
		result1 worker.Image
	}{result1}
}

func (fake *FakeImageFactory) NewImageReturnsOnCall(i int, result1 worker.Image) {
	fake.NewImageStub = nil
	if fake.newImageReturnsOnCall == nil {
		fake.newImageReturnsOnCall = make(map[int]struct {
			result1 worker.Image
		})
	}
	fake.newImageReturnsOnCall[i] = struct {
		result1 worker.Image
	}{result1}
}

func (fake *FakeImageFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newImageMutex.RLock()
	defer fake.newImageMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImageFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.ImageFactory = new(FakeImageFactory)
