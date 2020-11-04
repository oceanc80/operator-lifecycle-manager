// Code generated by counterfeiter. DO NOT EDIT.
package listers

import (
	"sync"

	v1a "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/listers/rbac/v1"
)

type FakeRoleBindingLister struct {
	ListStub        func(labels.Selector) ([]*v1a.RoleBinding, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 labels.Selector
	}
	listReturns struct {
		result1 []*v1a.RoleBinding
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []*v1a.RoleBinding
		result2 error
	}
	RoleBindingsStub        func(string) v1.RoleBindingNamespaceLister
	roleBindingsMutex       sync.RWMutex
	roleBindingsArgsForCall []struct {
		arg1 string
	}
	roleBindingsReturns struct {
		result1 v1.RoleBindingNamespaceLister
	}
	roleBindingsReturnsOnCall map[int]struct {
		result1 v1.RoleBindingNamespaceLister
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoleBindingLister) List(arg1 labels.Selector) ([]*v1a.RoleBinding, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 labels.Selector
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRoleBindingLister) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeRoleBindingLister) ListCalls(stub func(labels.Selector) ([]*v1a.RoleBinding, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeRoleBindingLister) ListArgsForCall(i int) labels.Selector {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRoleBindingLister) ListReturns(result1 []*v1a.RoleBinding, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*v1a.RoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleBindingLister) ListReturnsOnCall(i int, result1 []*v1a.RoleBinding, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*v1a.RoleBinding
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*v1a.RoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleBindingLister) RoleBindings(arg1 string) v1.RoleBindingNamespaceLister {
	fake.roleBindingsMutex.Lock()
	ret, specificReturn := fake.roleBindingsReturnsOnCall[len(fake.roleBindingsArgsForCall)]
	fake.roleBindingsArgsForCall = append(fake.roleBindingsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RoleBindings", []interface{}{arg1})
	fake.roleBindingsMutex.Unlock()
	if fake.RoleBindingsStub != nil {
		return fake.RoleBindingsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.roleBindingsReturns
	return fakeReturns.result1
}

func (fake *FakeRoleBindingLister) RoleBindingsCallCount() int {
	fake.roleBindingsMutex.RLock()
	defer fake.roleBindingsMutex.RUnlock()
	return len(fake.roleBindingsArgsForCall)
}

func (fake *FakeRoleBindingLister) RoleBindingsCalls(stub func(string) v1.RoleBindingNamespaceLister) {
	fake.roleBindingsMutex.Lock()
	defer fake.roleBindingsMutex.Unlock()
	fake.RoleBindingsStub = stub
}

func (fake *FakeRoleBindingLister) RoleBindingsArgsForCall(i int) string {
	fake.roleBindingsMutex.RLock()
	defer fake.roleBindingsMutex.RUnlock()
	argsForCall := fake.roleBindingsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRoleBindingLister) RoleBindingsReturns(result1 v1.RoleBindingNamespaceLister) {
	fake.roleBindingsMutex.Lock()
	defer fake.roleBindingsMutex.Unlock()
	fake.RoleBindingsStub = nil
	fake.roleBindingsReturns = struct {
		result1 v1.RoleBindingNamespaceLister
	}{result1}
}

func (fake *FakeRoleBindingLister) RoleBindingsReturnsOnCall(i int, result1 v1.RoleBindingNamespaceLister) {
	fake.roleBindingsMutex.Lock()
	defer fake.roleBindingsMutex.Unlock()
	fake.RoleBindingsStub = nil
	if fake.roleBindingsReturnsOnCall == nil {
		fake.roleBindingsReturnsOnCall = make(map[int]struct {
			result1 v1.RoleBindingNamespaceLister
		})
	}
	fake.roleBindingsReturnsOnCall[i] = struct {
		result1 v1.RoleBindingNamespaceLister
	}{result1}
}

func (fake *FakeRoleBindingLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.roleBindingsMutex.RLock()
	defer fake.roleBindingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoleBindingLister) recordInvocation(key string, args []interface{}) {
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

var _ v1.RoleBindingLister = new(FakeRoleBindingLister)
