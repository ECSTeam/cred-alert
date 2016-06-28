// This file was generated by counterfeiter
package loggingfakes

import (
	"cred-alert/logging"
	"sync"

	"github.com/pivotal-golang/lager"
)

type FakeEmitter struct {
	CountViolationStub        func(logger lager.Logger, count int)
	countViolationMutex       sync.RWMutex
	countViolationArgsForCall []struct {
		logger lager.Logger
		count  int
	}
	CountAPIRequestStub        func(logger lager.Logger)
	countAPIRequestMutex       sync.RWMutex
	countAPIRequestArgsForCall []struct {
		logger lager.Logger
	}
	CounterStub        func(name string) logging.Counter
	counterMutex       sync.RWMutex
	counterArgsForCall []struct {
		name string
	}
	counterReturns struct {
		result1 logging.Counter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) CountViolation(logger lager.Logger, count int) {
	fake.countViolationMutex.Lock()
	fake.countViolationArgsForCall = append(fake.countViolationArgsForCall, struct {
		logger lager.Logger
		count  int
	}{logger, count})
	fake.recordInvocation("CountViolation", []interface{}{logger, count})
	fake.countViolationMutex.Unlock()
	if fake.CountViolationStub != nil {
		fake.CountViolationStub(logger, count)
	}
}

func (fake *FakeEmitter) CountViolationCallCount() int {
	fake.countViolationMutex.RLock()
	defer fake.countViolationMutex.RUnlock()
	return len(fake.countViolationArgsForCall)
}

func (fake *FakeEmitter) CountViolationArgsForCall(i int) (lager.Logger, int) {
	fake.countViolationMutex.RLock()
	defer fake.countViolationMutex.RUnlock()
	return fake.countViolationArgsForCall[i].logger, fake.countViolationArgsForCall[i].count
}

func (fake *FakeEmitter) CountAPIRequest(logger lager.Logger) {
	fake.countAPIRequestMutex.Lock()
	fake.countAPIRequestArgsForCall = append(fake.countAPIRequestArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("CountAPIRequest", []interface{}{logger})
	fake.countAPIRequestMutex.Unlock()
	if fake.CountAPIRequestStub != nil {
		fake.CountAPIRequestStub(logger)
	}
}

func (fake *FakeEmitter) CountAPIRequestCallCount() int {
	fake.countAPIRequestMutex.RLock()
	defer fake.countAPIRequestMutex.RUnlock()
	return len(fake.countAPIRequestArgsForCall)
}

func (fake *FakeEmitter) CountAPIRequestArgsForCall(i int) lager.Logger {
	fake.countAPIRequestMutex.RLock()
	defer fake.countAPIRequestMutex.RUnlock()
	return fake.countAPIRequestArgsForCall[i].logger
}

func (fake *FakeEmitter) Counter(name string) logging.Counter {
	fake.counterMutex.Lock()
	fake.counterArgsForCall = append(fake.counterArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Counter", []interface{}{name})
	fake.counterMutex.Unlock()
	if fake.CounterStub != nil {
		return fake.CounterStub(name)
	} else {
		return fake.counterReturns.result1
	}
}

func (fake *FakeEmitter) CounterCallCount() int {
	fake.counterMutex.RLock()
	defer fake.counterMutex.RUnlock()
	return len(fake.counterArgsForCall)
}

func (fake *FakeEmitter) CounterArgsForCall(i int) string {
	fake.counterMutex.RLock()
	defer fake.counterMutex.RUnlock()
	return fake.counterArgsForCall[i].name
}

func (fake *FakeEmitter) CounterReturns(result1 logging.Counter) {
	fake.CounterStub = nil
	fake.counterReturns = struct {
		result1 logging.Counter
	}{result1}
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.countViolationMutex.RLock()
	defer fake.countViolationMutex.RUnlock()
	fake.countAPIRequestMutex.RLock()
	defer fake.countAPIRequestMutex.RUnlock()
	fake.counterMutex.RLock()
	defer fake.counterMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
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

var _ logging.Emitter = new(FakeEmitter)
