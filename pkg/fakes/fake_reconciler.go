// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry/reconciler"
	"github.com/sirupsen/logrus"
)

type FakeRegistryReconciler struct {
	CheckRegistryServerStub        func(*logrus.Entry, *v1alpha1.CatalogSource) (bool, error)
	checkRegistryServerMutex       sync.RWMutex
	checkRegistryServerArgsForCall []struct {
		arg1 *logrus.Entry
		arg2 *v1alpha1.CatalogSource
	}
	checkRegistryServerReturns struct {
		result1 bool
		result2 error
	}
	checkRegistryServerReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	EnsureRegistryServerStub        func(*logrus.Entry, *v1alpha1.CatalogSource) error
	ensureRegistryServerMutex       sync.RWMutex
	ensureRegistryServerArgsForCall []struct {
		arg1 *logrus.Entry
		arg2 *v1alpha1.CatalogSource
	}
	ensureRegistryServerReturns struct {
		result1 error
	}
	ensureRegistryServerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistryReconciler) CheckRegistryServer(arg1 *logrus.Entry, arg2 *v1alpha1.CatalogSource) (bool, error) {
	fake.checkRegistryServerMutex.Lock()
	ret, specificReturn := fake.checkRegistryServerReturnsOnCall[len(fake.checkRegistryServerArgsForCall)]
	fake.checkRegistryServerArgsForCall = append(fake.checkRegistryServerArgsForCall, struct {
		arg1 *logrus.Entry
		arg2 *v1alpha1.CatalogSource
	}{arg1, arg2})
	fake.recordInvocation("CheckRegistryServer", []interface{}{arg1, arg2})
	fake.checkRegistryServerMutex.Unlock()
	if fake.CheckRegistryServerStub != nil {
		return fake.CheckRegistryServerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkRegistryServerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistryReconciler) CheckRegistryServerCallCount() int {
	fake.checkRegistryServerMutex.RLock()
	defer fake.checkRegistryServerMutex.RUnlock()
	return len(fake.checkRegistryServerArgsForCall)
}

func (fake *FakeRegistryReconciler) CheckRegistryServerCalls(stub func(*logrus.Entry, *v1alpha1.CatalogSource) (bool, error)) {
	fake.checkRegistryServerMutex.Lock()
	defer fake.checkRegistryServerMutex.Unlock()
	fake.CheckRegistryServerStub = stub
}

func (fake *FakeRegistryReconciler) CheckRegistryServerArgsForCall(i int) (*logrus.Entry, *v1alpha1.CatalogSource) {
	fake.checkRegistryServerMutex.RLock()
	defer fake.checkRegistryServerMutex.RUnlock()
	argsForCall := fake.checkRegistryServerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistryReconciler) CheckRegistryServerReturns(result1 bool, result2 error) {
	fake.checkRegistryServerMutex.Lock()
	defer fake.checkRegistryServerMutex.Unlock()
	fake.CheckRegistryServerStub = nil
	fake.checkRegistryServerReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryReconciler) CheckRegistryServerReturnsOnCall(i int, result1 bool, result2 error) {
	fake.checkRegistryServerMutex.Lock()
	defer fake.checkRegistryServerMutex.Unlock()
	fake.CheckRegistryServerStub = nil
	if fake.checkRegistryServerReturnsOnCall == nil {
		fake.checkRegistryServerReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkRegistryServerReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryReconciler) EnsureRegistryServer(arg1 *logrus.Entry, arg2 *v1alpha1.CatalogSource) error {
	fake.ensureRegistryServerMutex.Lock()
	ret, specificReturn := fake.ensureRegistryServerReturnsOnCall[len(fake.ensureRegistryServerArgsForCall)]
	fake.ensureRegistryServerArgsForCall = append(fake.ensureRegistryServerArgsForCall, struct {
		arg1 *logrus.Entry
		arg2 *v1alpha1.CatalogSource
	}{arg1, arg2})
	fake.recordInvocation("EnsureRegistryServer", []interface{}{arg1, arg2})
	fake.ensureRegistryServerMutex.Unlock()
	if fake.EnsureRegistryServerStub != nil {
		return fake.EnsureRegistryServerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ensureRegistryServerReturns
	return fakeReturns.result1
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerCallCount() int {
	fake.ensureRegistryServerMutex.RLock()
	defer fake.ensureRegistryServerMutex.RUnlock()
	return len(fake.ensureRegistryServerArgsForCall)
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerCalls(stub func(*logrus.Entry, *v1alpha1.CatalogSource) error) {
	fake.ensureRegistryServerMutex.Lock()
	defer fake.ensureRegistryServerMutex.Unlock()
	fake.EnsureRegistryServerStub = stub
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerArgsForCall(i int) (*logrus.Entry, *v1alpha1.CatalogSource) {
	fake.ensureRegistryServerMutex.RLock()
	defer fake.ensureRegistryServerMutex.RUnlock()
	argsForCall := fake.ensureRegistryServerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerReturns(result1 error) {
	fake.ensureRegistryServerMutex.Lock()
	defer fake.ensureRegistryServerMutex.Unlock()
	fake.EnsureRegistryServerStub = nil
	fake.ensureRegistryServerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerReturnsOnCall(i int, result1 error) {
	fake.ensureRegistryServerMutex.Lock()
	defer fake.ensureRegistryServerMutex.Unlock()
	fake.EnsureRegistryServerStub = nil
	if fake.ensureRegistryServerReturnsOnCall == nil {
		fake.ensureRegistryServerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ensureRegistryServerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistryReconciler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkRegistryServerMutex.RLock()
	defer fake.checkRegistryServerMutex.RUnlock()
	fake.ensureRegistryServerMutex.RLock()
	defer fake.ensureRegistryServerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistryReconciler) recordInvocation(key string, args []interface{}) {
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

var _ reconciler.RegistryReconciler = new(FakeRegistryReconciler)
