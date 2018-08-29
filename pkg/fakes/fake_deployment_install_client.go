// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/wrappers"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/ownerutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1rbac "k8s.io/api/rbac/v1beta1"
)

type FakeInstallStrategyDeploymentInterface struct {
	CreateRoleStub        func(role *v1beta1rbac.Role) (*v1beta1rbac.Role, error)
	createRoleMutex       sync.RWMutex
	createRoleArgsForCall []struct {
		role *v1beta1rbac.Role
	}
	createRoleReturns struct {
		result1 *v1beta1rbac.Role
		result2 error
	}
	createRoleReturnsOnCall map[int]struct {
		result1 *v1beta1rbac.Role
		result2 error
	}
	CreateRoleBindingStub        func(roleBinding *v1beta1rbac.RoleBinding) (*v1beta1rbac.RoleBinding, error)
	createRoleBindingMutex       sync.RWMutex
	createRoleBindingArgsForCall []struct {
		roleBinding *v1beta1rbac.RoleBinding
	}
	createRoleBindingReturns struct {
		result1 *v1beta1rbac.RoleBinding
		result2 error
	}
	createRoleBindingReturnsOnCall map[int]struct {
		result1 *v1beta1rbac.RoleBinding
		result2 error
	}
	EnsureServiceAccountStub        func(serviceAccount *corev1.ServiceAccount, owner ownerutil.Owner) (*corev1.ServiceAccount, error)
	ensureServiceAccountMutex       sync.RWMutex
	ensureServiceAccountArgsForCall []struct {
		serviceAccount *corev1.ServiceAccount
		owner          ownerutil.Owner
	}
	ensureServiceAccountReturns struct {
		result1 *corev1.ServiceAccount
		result2 error
	}
	ensureServiceAccountReturnsOnCall map[int]struct {
		result1 *corev1.ServiceAccount
		result2 error
	}
	CreateDeploymentStub        func(deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	createDeploymentMutex       sync.RWMutex
	createDeploymentArgsForCall []struct {
		deployment *appsv1.Deployment
	}
	createDeploymentReturns struct {
		result1 *appsv1.Deployment
		result2 error
	}
	createDeploymentReturnsOnCall map[int]struct {
		result1 *appsv1.Deployment
		result2 error
	}
	CreateOrUpdateDeploymentStub        func(deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	createOrUpdateDeploymentMutex       sync.RWMutex
	createOrUpdateDeploymentArgsForCall []struct {
		deployment *appsv1.Deployment
	}
	createOrUpdateDeploymentReturns struct {
		result1 *appsv1.Deployment
		result2 error
	}
	createOrUpdateDeploymentReturnsOnCall map[int]struct {
		result1 *appsv1.Deployment
		result2 error
	}
	DeleteDeploymentStub        func(name string) error
	deleteDeploymentMutex       sync.RWMutex
	deleteDeploymentArgsForCall []struct {
		name string
	}
	deleteDeploymentReturns struct {
		result1 error
	}
	deleteDeploymentReturnsOnCall map[int]struct {
		result1 error
	}
	GetServiceAccountByNameStub        func(serviceAccountName string) (*corev1.ServiceAccount, error)
	getServiceAccountByNameMutex       sync.RWMutex
	getServiceAccountByNameArgsForCall []struct {
		serviceAccountName string
	}
	getServiceAccountByNameReturns struct {
		result1 *corev1.ServiceAccount
		result2 error
	}
	getServiceAccountByNameReturnsOnCall map[int]struct {
		result1 *corev1.ServiceAccount
		result2 error
	}
	FindAnyDeploymentsMatchingNamesStub        func(depNames []string) ([]*appsv1.Deployment, error)
	findAnyDeploymentsMatchingNamesMutex       sync.RWMutex
	findAnyDeploymentsMatchingNamesArgsForCall []struct {
		depNames []string
	}
	findAnyDeploymentsMatchingNamesReturns struct {
		result1 []*appsv1.Deployment
		result2 error
	}
	findAnyDeploymentsMatchingNamesReturnsOnCall map[int]struct {
		result1 []*appsv1.Deployment
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRole(role *v1beta1rbac.Role) (*v1beta1rbac.Role, error) {
	fake.createRoleMutex.Lock()
	ret, specificReturn := fake.createRoleReturnsOnCall[len(fake.createRoleArgsForCall)]
	fake.createRoleArgsForCall = append(fake.createRoleArgsForCall, struct {
		role *v1beta1rbac.Role
	}{role})
	fake.recordInvocation("CreateRole", []interface{}{role})
	fake.createRoleMutex.Unlock()
	if fake.CreateRoleStub != nil {
		return fake.CreateRoleStub(role)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createRoleReturns.result1, fake.createRoleReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleCallCount() int {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	return len(fake.createRoleArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleArgsForCall(i int) *v1beta1rbac.Role {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	return fake.createRoleArgsForCall[i].role
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleReturns(result1 *v1beta1rbac.Role, result2 error) {
	fake.CreateRoleStub = nil
	fake.createRoleReturns = struct {
		result1 *v1beta1rbac.Role
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleReturnsOnCall(i int, result1 *v1beta1rbac.Role, result2 error) {
	fake.CreateRoleStub = nil
	if fake.createRoleReturnsOnCall == nil {
		fake.createRoleReturnsOnCall = make(map[int]struct {
			result1 *v1beta1rbac.Role
			result2 error
		})
	}
	fake.createRoleReturnsOnCall[i] = struct {
		result1 *v1beta1rbac.Role
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBinding(roleBinding *v1beta1rbac.RoleBinding) (*v1beta1rbac.RoleBinding, error) {
	fake.createRoleBindingMutex.Lock()
	ret, specificReturn := fake.createRoleBindingReturnsOnCall[len(fake.createRoleBindingArgsForCall)]
	fake.createRoleBindingArgsForCall = append(fake.createRoleBindingArgsForCall, struct {
		roleBinding *v1beta1rbac.RoleBinding
	}{roleBinding})
	fake.recordInvocation("CreateRoleBinding", []interface{}{roleBinding})
	fake.createRoleBindingMutex.Unlock()
	if fake.CreateRoleBindingStub != nil {
		return fake.CreateRoleBindingStub(roleBinding)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createRoleBindingReturns.result1, fake.createRoleBindingReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingCallCount() int {
	fake.createRoleBindingMutex.RLock()
	defer fake.createRoleBindingMutex.RUnlock()
	return len(fake.createRoleBindingArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingArgsForCall(i int) *v1beta1rbac.RoleBinding {
	fake.createRoleBindingMutex.RLock()
	defer fake.createRoleBindingMutex.RUnlock()
	return fake.createRoleBindingArgsForCall[i].roleBinding
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingReturns(result1 *v1beta1rbac.RoleBinding, result2 error) {
	fake.CreateRoleBindingStub = nil
	fake.createRoleBindingReturns = struct {
		result1 *v1beta1rbac.RoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingReturnsOnCall(i int, result1 *v1beta1rbac.RoleBinding, result2 error) {
	fake.CreateRoleBindingStub = nil
	if fake.createRoleBindingReturnsOnCall == nil {
		fake.createRoleBindingReturnsOnCall = make(map[int]struct {
			result1 *v1beta1rbac.RoleBinding
			result2 error
		})
	}
	fake.createRoleBindingReturnsOnCall[i] = struct {
		result1 *v1beta1rbac.RoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccount(serviceAccount *corev1.ServiceAccount, owner ownerutil.Owner) (*corev1.ServiceAccount, error) {
	fake.ensureServiceAccountMutex.Lock()
	ret, specificReturn := fake.ensureServiceAccountReturnsOnCall[len(fake.ensureServiceAccountArgsForCall)]
	fake.ensureServiceAccountArgsForCall = append(fake.ensureServiceAccountArgsForCall, struct {
		serviceAccount *corev1.ServiceAccount
		owner          ownerutil.Owner
	}{serviceAccount, owner})
	fake.recordInvocation("EnsureServiceAccount", []interface{}{serviceAccount, owner})
	fake.ensureServiceAccountMutex.Unlock()
	if fake.EnsureServiceAccountStub != nil {
		return fake.EnsureServiceAccountStub(serviceAccount, owner)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ensureServiceAccountReturns.result1, fake.ensureServiceAccountReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountCallCount() int {
	fake.ensureServiceAccountMutex.RLock()
	defer fake.ensureServiceAccountMutex.RUnlock()
	return len(fake.ensureServiceAccountArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountArgsForCall(i int) (*corev1.ServiceAccount, ownerutil.Owner) {
	fake.ensureServiceAccountMutex.RLock()
	defer fake.ensureServiceAccountMutex.RUnlock()
	return fake.ensureServiceAccountArgsForCall[i].serviceAccount, fake.ensureServiceAccountArgsForCall[i].owner
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountReturns(result1 *corev1.ServiceAccount, result2 error) {
	fake.EnsureServiceAccountStub = nil
	fake.ensureServiceAccountReturns = struct {
		result1 *corev1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountReturnsOnCall(i int, result1 *corev1.ServiceAccount, result2 error) {
	fake.EnsureServiceAccountStub = nil
	if fake.ensureServiceAccountReturnsOnCall == nil {
		fake.ensureServiceAccountReturnsOnCall = make(map[int]struct {
			result1 *corev1.ServiceAccount
			result2 error
		})
	}
	fake.ensureServiceAccountReturnsOnCall[i] = struct {
		result1 *corev1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeployment(deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	fake.createDeploymentMutex.Lock()
	ret, specificReturn := fake.createDeploymentReturnsOnCall[len(fake.createDeploymentArgsForCall)]
	fake.createDeploymentArgsForCall = append(fake.createDeploymentArgsForCall, struct {
		deployment *appsv1.Deployment
	}{deployment})
	fake.recordInvocation("CreateDeployment", []interface{}{deployment})
	fake.createDeploymentMutex.Unlock()
	if fake.CreateDeploymentStub != nil {
		return fake.CreateDeploymentStub(deployment)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createDeploymentReturns.result1, fake.createDeploymentReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentCallCount() int {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return len(fake.createDeploymentArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentArgsForCall(i int) *appsv1.Deployment {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return fake.createDeploymentArgsForCall[i].deployment
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentReturns(result1 *appsv1.Deployment, result2 error) {
	fake.CreateDeploymentStub = nil
	fake.createDeploymentReturns = struct {
		result1 *appsv1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentReturnsOnCall(i int, result1 *appsv1.Deployment, result2 error) {
	fake.CreateDeploymentStub = nil
	if fake.createDeploymentReturnsOnCall == nil {
		fake.createDeploymentReturnsOnCall = make(map[int]struct {
			result1 *appsv1.Deployment
			result2 error
		})
	}
	fake.createDeploymentReturnsOnCall[i] = struct {
		result1 *appsv1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeployment(deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	fake.createOrUpdateDeploymentMutex.Lock()
	ret, specificReturn := fake.createOrUpdateDeploymentReturnsOnCall[len(fake.createOrUpdateDeploymentArgsForCall)]
	fake.createOrUpdateDeploymentArgsForCall = append(fake.createOrUpdateDeploymentArgsForCall, struct {
		deployment *appsv1.Deployment
	}{deployment})
	fake.recordInvocation("CreateOrUpdateDeployment", []interface{}{deployment})
	fake.createOrUpdateDeploymentMutex.Unlock()
	if fake.CreateOrUpdateDeploymentStub != nil {
		return fake.CreateOrUpdateDeploymentStub(deployment)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createOrUpdateDeploymentReturns.result1, fake.createOrUpdateDeploymentReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentCallCount() int {
	fake.createOrUpdateDeploymentMutex.RLock()
	defer fake.createOrUpdateDeploymentMutex.RUnlock()
	return len(fake.createOrUpdateDeploymentArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentArgsForCall(i int) *appsv1.Deployment {
	fake.createOrUpdateDeploymentMutex.RLock()
	defer fake.createOrUpdateDeploymentMutex.RUnlock()
	return fake.createOrUpdateDeploymentArgsForCall[i].deployment
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentReturns(result1 *appsv1.Deployment, result2 error) {
	fake.CreateOrUpdateDeploymentStub = nil
	fake.createOrUpdateDeploymentReturns = struct {
		result1 *appsv1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentReturnsOnCall(i int, result1 *appsv1.Deployment, result2 error) {
	fake.CreateOrUpdateDeploymentStub = nil
	if fake.createOrUpdateDeploymentReturnsOnCall == nil {
		fake.createOrUpdateDeploymentReturnsOnCall = make(map[int]struct {
			result1 *appsv1.Deployment
			result2 error
		})
	}
	fake.createOrUpdateDeploymentReturnsOnCall[i] = struct {
		result1 *appsv1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeployment(name string) error {
	fake.deleteDeploymentMutex.Lock()
	ret, specificReturn := fake.deleteDeploymentReturnsOnCall[len(fake.deleteDeploymentArgsForCall)]
	fake.deleteDeploymentArgsForCall = append(fake.deleteDeploymentArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("DeleteDeployment", []interface{}{name})
	fake.deleteDeploymentMutex.Unlock()
	if fake.DeleteDeploymentStub != nil {
		return fake.DeleteDeploymentStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteDeploymentReturns.result1
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentCallCount() int {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return len(fake.deleteDeploymentArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentArgsForCall(i int) string {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return fake.deleteDeploymentArgsForCall[i].name
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentReturns(result1 error) {
	fake.DeleteDeploymentStub = nil
	fake.deleteDeploymentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentReturnsOnCall(i int, result1 error) {
	fake.DeleteDeploymentStub = nil
	if fake.deleteDeploymentReturnsOnCall == nil {
		fake.deleteDeploymentReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteDeploymentReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByName(serviceAccountName string) (*corev1.ServiceAccount, error) {
	fake.getServiceAccountByNameMutex.Lock()
	ret, specificReturn := fake.getServiceAccountByNameReturnsOnCall[len(fake.getServiceAccountByNameArgsForCall)]
	fake.getServiceAccountByNameArgsForCall = append(fake.getServiceAccountByNameArgsForCall, struct {
		serviceAccountName string
	}{serviceAccountName})
	fake.recordInvocation("GetServiceAccountByName", []interface{}{serviceAccountName})
	fake.getServiceAccountByNameMutex.Unlock()
	if fake.GetServiceAccountByNameStub != nil {
		return fake.GetServiceAccountByNameStub(serviceAccountName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceAccountByNameReturns.result1, fake.getServiceAccountByNameReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameCallCount() int {
	fake.getServiceAccountByNameMutex.RLock()
	defer fake.getServiceAccountByNameMutex.RUnlock()
	return len(fake.getServiceAccountByNameArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameArgsForCall(i int) string {
	fake.getServiceAccountByNameMutex.RLock()
	defer fake.getServiceAccountByNameMutex.RUnlock()
	return fake.getServiceAccountByNameArgsForCall[i].serviceAccountName
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameReturns(result1 *corev1.ServiceAccount, result2 error) {
	fake.GetServiceAccountByNameStub = nil
	fake.getServiceAccountByNameReturns = struct {
		result1 *corev1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameReturnsOnCall(i int, result1 *corev1.ServiceAccount, result2 error) {
	fake.GetServiceAccountByNameStub = nil
	if fake.getServiceAccountByNameReturnsOnCall == nil {
		fake.getServiceAccountByNameReturnsOnCall = make(map[int]struct {
			result1 *corev1.ServiceAccount
			result2 error
		})
	}
	fake.getServiceAccountByNameReturnsOnCall[i] = struct {
		result1 *corev1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNames(depNames []string) ([]*appsv1.Deployment, error) {
	var depNamesCopy []string
	if depNames != nil {
		depNamesCopy = make([]string, len(depNames))
		copy(depNamesCopy, depNames)
	}
	fake.findAnyDeploymentsMatchingNamesMutex.Lock()
	ret, specificReturn := fake.findAnyDeploymentsMatchingNamesReturnsOnCall[len(fake.findAnyDeploymentsMatchingNamesArgsForCall)]
	fake.findAnyDeploymentsMatchingNamesArgsForCall = append(fake.findAnyDeploymentsMatchingNamesArgsForCall, struct {
		depNames []string
	}{depNamesCopy})
	fake.recordInvocation("FindAnyDeploymentsMatchingNames", []interface{}{depNamesCopy})
	fake.findAnyDeploymentsMatchingNamesMutex.Unlock()
	if fake.FindAnyDeploymentsMatchingNamesStub != nil {
		return fake.FindAnyDeploymentsMatchingNamesStub(depNames)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findAnyDeploymentsMatchingNamesReturns.result1, fake.findAnyDeploymentsMatchingNamesReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesCallCount() int {
	fake.findAnyDeploymentsMatchingNamesMutex.RLock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.RUnlock()
	return len(fake.findAnyDeploymentsMatchingNamesArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesArgsForCall(i int) []string {
	fake.findAnyDeploymentsMatchingNamesMutex.RLock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.RUnlock()
	return fake.findAnyDeploymentsMatchingNamesArgsForCall[i].depNames
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesReturns(result1 []*appsv1.Deployment, result2 error) {
	fake.FindAnyDeploymentsMatchingNamesStub = nil
	fake.findAnyDeploymentsMatchingNamesReturns = struct {
		result1 []*appsv1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesReturnsOnCall(i int, result1 []*appsv1.Deployment, result2 error) {
	fake.FindAnyDeploymentsMatchingNamesStub = nil
	if fake.findAnyDeploymentsMatchingNamesReturnsOnCall == nil {
		fake.findAnyDeploymentsMatchingNamesReturnsOnCall = make(map[int]struct {
			result1 []*appsv1.Deployment
			result2 error
		})
	}
	fake.findAnyDeploymentsMatchingNamesReturnsOnCall[i] = struct {
		result1 []*appsv1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	fake.createRoleBindingMutex.RLock()
	defer fake.createRoleBindingMutex.RUnlock()
	fake.ensureServiceAccountMutex.RLock()
	defer fake.ensureServiceAccountMutex.RUnlock()
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	fake.createOrUpdateDeploymentMutex.RLock()
	defer fake.createOrUpdateDeploymentMutex.RUnlock()
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	fake.getServiceAccountByNameMutex.RLock()
	defer fake.getServiceAccountByNameMutex.RUnlock()
	fake.findAnyDeploymentsMatchingNamesMutex.RLock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstallStrategyDeploymentInterface) recordInvocation(key string, args []interface{}) {
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

var _ wrappers.InstallStrategyDeploymentInterface = new(FakeInstallStrategyDeploymentInterface)
