package operatorclient

import (
	"context"
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	acv1 "k8s.io/client-go/applyconfigurations/rbac/v1"
	"k8s.io/klog"
)

// ApplyRoleBinding applies the roleBinding.
func (c *Client) ApplyRoleBinding(applyConfig *acv1.RoleBindingApplyConfiguration, applyOptions metav1.ApplyOptions) (*rbacv1.RoleBinding, error) {
	return c.RbacV1().RoleBindings(*applyConfig.Namespace).Apply(context.TODO(), applyConfig, applyOptions)
}

// CreateRoleBinding creates the roleBinding or Updates if it already exists.
func (c *Client) CreateRoleBinding(ig *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {
	createdRB, err := c.RbacV1().RoleBindings(ig.GetNamespace()).Create(context.TODO(), ig, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return c.UpdateRoleBinding(ig)
	}
	return createdRB, err
}

// GetRoleBinding returns the existing roleBinding.
func (c *Client) GetRoleBinding(namespace, name string) (*rbacv1.RoleBinding, error) {
	return c.RbacV1().RoleBindings(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

// DeleteRoleBinding deletes the roleBinding.
func (c *Client) DeleteRoleBinding(namespace, name string, options *metav1.DeleteOptions) error {
	return c.RbacV1().RoleBindings(namespace).Delete(context.TODO(), name, *options)
}

// UpdateRoleBinding will update the given RoleBinding resource.
func (c *Client) UpdateRoleBinding(crb *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {
	klog.V(4).Infof("[UPDATE RoleBinding]: %s", crb.GetName())
	oldCrb, err := c.GetRoleBinding(crb.GetNamespace(), crb.GetName())
	if err != nil {
		return nil, err
	}
	patchBytes, err := createPatch(oldCrb, crb)
	if err != nil {
		return nil, fmt.Errorf("error creating patch for RoleBinding: %v", err)
	}
	return c.RbacV1().RoleBindings(crb.GetNamespace()).Patch(context.TODO(), crb.GetName(), types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
}
