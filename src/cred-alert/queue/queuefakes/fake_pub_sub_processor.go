// This file was generated by counterfeiter
package queuefakes

import (
	"cred-alert/queue"
	"sync"

	"cloud.google.com/go/pubsub"
	"code.cloudfoundry.org/lager"
)

type FakePubSubProcessor struct {
	ProcessStub        func(lager.Logger, *pubsub.Message) (bool, error)
	processMutex       sync.RWMutex
	processArgsForCall []struct {
		arg1 lager.Logger
		arg2 *pubsub.Message
	}
	processReturns struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePubSubProcessor) Process(arg1 lager.Logger, arg2 *pubsub.Message) (bool, error) {
	fake.processMutex.Lock()
	fake.processArgsForCall = append(fake.processArgsForCall, struct {
		arg1 lager.Logger
		arg2 *pubsub.Message
	}{arg1, arg2})
	fake.recordInvocation("Process", []interface{}{arg1, arg2})
	fake.processMutex.Unlock()
	if fake.ProcessStub != nil {
		return fake.ProcessStub(arg1, arg2)
	} else {
		return fake.processReturns.result1, fake.processReturns.result2
	}
}

func (fake *FakePubSubProcessor) ProcessCallCount() int {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return len(fake.processArgsForCall)
}

func (fake *FakePubSubProcessor) ProcessArgsForCall(i int) (lager.Logger, *pubsub.Message) {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return fake.processArgsForCall[i].arg1, fake.processArgsForCall[i].arg2
}

func (fake *FakePubSubProcessor) ProcessReturns(result1 bool, result2 error) {
	fake.ProcessStub = nil
	fake.processReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePubSubProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePubSubProcessor) recordInvocation(key string, args []interface{}) {
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

var _ queue.PubSubProcessor = new(FakePubSubProcessor)
