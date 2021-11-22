// +build !ignore_autogenerated

/*
** Copyright (c) 2021 Oracle and/or its affiliates.
**
** The Universal Permissive License (UPL), Version 1.0
**
** Subject to the condition set forth below, permission is hereby granted to any
** person obtaining a copy of this software, associated documentation and/or data
** (collectively the "Software"), free of charge and under any and all copyright
** rights in the Software, and any and all patent rights owned or freely
** licensable by each licensor hereunder covering either (i) the unmodified
** Software as contributed to or provided by such licensor, or (ii) the Larger
** Works (as defined below), to deal in both
**
** (a) the Software, and
** (b) any piece of software and/or hardware listed in the lrgrwrks.txt file if
** one is included with the Software (each a "Larger Work" to which the Software
** is contributed by such licensors),
**
** without restriction, including without limitation the rights to copy, create
** derivative works of, display, perform, and distribute the Software and make,
** use, sell, offer for sale, import, export, have made, and have sold the
** Software and the Larger Work(s), and to sublicense the foregoing rights on
** either these or other terms.
**
** This license is subject to the following condition:
** The above copyright notice and either this complete permission notice or at
** a minimum a reference to the UPL must be included in all copies or
** substantial portions of the Software.
**
** THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
** IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
** FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
** AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
** LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
** OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
** SOFTWARE.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabase) DeepCopyInto(out *AutonomousDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabase.
func (in *AutonomousDatabase) DeepCopy() *AutonomousDatabase {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutonomousDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseBackup) DeepCopyInto(out *AutonomousDatabaseBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseBackup.
func (in *AutonomousDatabaseBackup) DeepCopy() *AutonomousDatabaseBackup {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutonomousDatabaseBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseBackupList) DeepCopyInto(out *AutonomousDatabaseBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutonomousDatabaseBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseBackupList.
func (in *AutonomousDatabaseBackupList) DeepCopy() *AutonomousDatabaseBackupList {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutonomousDatabaseBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseBackupSpec) DeepCopyInto(out *AutonomousDatabaseBackupSpec) {
	*out = *in
	in.OCIConfig.DeepCopyInto(&out.OCIConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseBackupSpec.
func (in *AutonomousDatabaseBackupSpec) DeepCopy() *AutonomousDatabaseBackupSpec {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseBackupStatus) DeepCopyInto(out *AutonomousDatabaseBackupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseBackupStatus.
func (in *AutonomousDatabaseBackupStatus) DeepCopy() *AutonomousDatabaseBackupStatus {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseDetails) DeepCopyInto(out *AutonomousDatabaseDetails) {
	*out = *in
	if in.AutonomousDatabaseOCID != nil {
		in, out := &in.AutonomousDatabaseOCID, &out.AutonomousDatabaseOCID
		*out = new(string)
		**out = **in
	}
	if in.CompartmentOCID != nil {
		in, out := &in.CompartmentOCID, &out.CompartmentOCID
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DbName != nil {
		in, out := &in.DbName, &out.DbName
		*out = new(string)
		**out = **in
	}
	if in.IsDedicated != nil {
		in, out := &in.IsDedicated, &out.IsDedicated
		*out = new(bool)
		**out = **in
	}
	if in.DbVersion != nil {
		in, out := &in.DbVersion, &out.DbVersion
		*out = new(string)
		**out = **in
	}
	if in.DataStorageSizeInTBs != nil {
		in, out := &in.DataStorageSizeInTBs, &out.DataStorageSizeInTBs
		*out = new(int)
		**out = **in
	}
	if in.CPUCoreCount != nil {
		in, out := &in.CPUCoreCount, &out.CPUCoreCount
		*out = new(int)
		**out = **in
	}
	in.AdminPassword.DeepCopyInto(&out.AdminPassword)
	if in.IsAutoScalingEnabled != nil {
		in, out := &in.IsAutoScalingEnabled, &out.IsAutoScalingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SubnetOCID != nil {
		in, out := &in.SubnetOCID, &out.SubnetOCID
		*out = new(string)
		**out = **in
	}
	if in.NsgOCIDs != nil {
		in, out := &in.NsgOCIDs, &out.NsgOCIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointLabel != nil {
		in, out := &in.PrivateEndpointLabel, &out.PrivateEndpointLabel
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointIP != nil {
		in, out := &in.PrivateEndpointIP, &out.PrivateEndpointIP
		*out = new(string)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Wallet.DeepCopyInto(&out.Wallet)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseDetails.
func (in *AutonomousDatabaseDetails) DeepCopy() *AutonomousDatabaseDetails {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseList) DeepCopyInto(out *AutonomousDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutonomousDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseList.
func (in *AutonomousDatabaseList) DeepCopy() *AutonomousDatabaseList {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutonomousDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseRestore) DeepCopyInto(out *AutonomousDatabaseRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseRestore.
func (in *AutonomousDatabaseRestore) DeepCopy() *AutonomousDatabaseRestore {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutonomousDatabaseRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseRestoreList) DeepCopyInto(out *AutonomousDatabaseRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutonomousDatabaseRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseRestoreList.
func (in *AutonomousDatabaseRestoreList) DeepCopy() *AutonomousDatabaseRestoreList {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutonomousDatabaseRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseRestoreSource) DeepCopyInto(out *AutonomousDatabaseRestoreSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseRestoreSource.
func (in *AutonomousDatabaseRestoreSource) DeepCopy() *AutonomousDatabaseRestoreSource {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseRestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseRestoreSpec) DeepCopyInto(out *AutonomousDatabaseRestoreSpec) {
	*out = *in
	out.Source = in.Source
	in.OCIConfig.DeepCopyInto(&out.OCIConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseRestoreSpec.
func (in *AutonomousDatabaseRestoreSpec) DeepCopy() *AutonomousDatabaseRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseRestoreStatus) DeepCopyInto(out *AutonomousDatabaseRestoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseRestoreStatus.
func (in *AutonomousDatabaseRestoreStatus) DeepCopy() *AutonomousDatabaseRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseSpec) DeepCopyInto(out *AutonomousDatabaseSpec) {
	*out = *in
	in.Details.DeepCopyInto(&out.Details)
	in.OCIConfig.DeepCopyInto(&out.OCIConfig)
	if in.HardLink != nil {
		in, out := &in.HardLink, &out.HardLink
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseSpec.
func (in *AutonomousDatabaseSpec) DeepCopy() *AutonomousDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutonomousDatabaseStatus) DeepCopyInto(out *AutonomousDatabaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutonomousDatabaseStatus.
func (in *AutonomousDatabaseStatus) DeepCopy() *AutonomousDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(AutonomousDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpec) DeepCopyInto(out *CatalogSpec) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]EnvironmentVariable, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PvAnnotations != nil {
		in, out := &in.PvAnnotations, &out.PvAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PvMatchLabels != nil {
		in, out := &in.PvMatchLabels, &out.PvMatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePulllPolicy != nil {
		in, out := &in.ImagePulllPolicy, &out.ImagePulllPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpec.
func (in *CatalogSpec) DeepCopy() *CatalogSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentVariable) DeepCopyInto(out *EnvironmentVariable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentVariable.
func (in *EnvironmentVariable) DeepCopy() *EnvironmentVariable {
	if in == nil {
		return nil
	}
	out := new(EnvironmentVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsmSpec) DeepCopyInto(out *GsmSpec) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]EnvironmentVariable, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PvMatchLabels != nil {
		in, out := &in.PvMatchLabels, &out.PvMatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePulllPolicy != nil {
		in, out := &in.ImagePulllPolicy, &out.ImagePulllPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsmSpec.
func (in *GsmSpec) DeepCopy() *GsmSpec {
	if in == nil {
		return nil
	}
	out := new(GsmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsmStatus) DeepCopyInto(out *GsmStatus) {
	*out = *in
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsmStatus.
func (in *GsmStatus) DeepCopy() *GsmStatus {
	if in == nil {
		return nil
	}
	out := new(GsmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsmStatusDetails) DeepCopyInto(out *GsmStatusDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsmStatusDetails.
func (in *GsmStatusDetails) DeepCopy() *GsmStatusDetails {
	if in == nil {
		return nil
	}
	out := new(GsmStatusDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIConfigSpec) DeepCopyInto(out *OCIConfigSpec) {
	*out = *in
	if in.ConfigMapName != nil {
		in, out := &in.ConfigMapName, &out.ConfigMapName
		*out = new(string)
		**out = **in
	}
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIConfigSpec.
func (in *OCIConfigSpec) DeepCopy() *OCIConfigSpec {
	if in == nil {
		return nil
	}
	out := new(OCIConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSpec) DeepCopyInto(out *PasswordSpec) {
	*out = *in
	if in.K8sSecretName != nil {
		in, out := &in.K8sSecretName, &out.K8sSecretName
		*out = new(string)
		**out = **in
	}
	if in.OCISecretOCID != nil {
		in, out := &in.OCISecretOCID, &out.OCISecretOCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSpec.
func (in *PasswordSpec) DeepCopy() *PasswordSpec {
	if in == nil {
		return nil
	}
	out := new(PasswordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortMapping) DeepCopyInto(out *PortMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortMapping.
func (in *PortMapping) DeepCopy() *PortMapping {
	if in == nil {
		return nil
	}
	out := new(PortMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardSpec) DeepCopyInto(out *ShardSpec) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]EnvironmentVariable, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PvAnnotations != nil {
		in, out := &in.PvAnnotations, &out.PvAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PvMatchLabels != nil {
		in, out := &in.PvMatchLabels, &out.PvMatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePulllPolicy != nil {
		in, out := &in.ImagePulllPolicy, &out.ImagePulllPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardSpec.
func (in *ShardSpec) DeepCopy() *ShardSpec {
	if in == nil {
		return nil
	}
	out := new(ShardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingDatabase) DeepCopyInto(out *ShardingDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingDatabase.
func (in *ShardingDatabase) DeepCopy() *ShardingDatabase {
	if in == nil {
		return nil
	}
	out := new(ShardingDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShardingDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingDatabaseList) DeepCopyInto(out *ShardingDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShardingDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingDatabaseList.
func (in *ShardingDatabaseList) DeepCopy() *ShardingDatabaseList {
	if in == nil {
		return nil
	}
	out := new(ShardingDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShardingDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingDatabaseSpec) DeepCopyInto(out *ShardingDatabaseSpec) {
	*out = *in
	if in.Shard != nil {
		in, out := &in.Shard, &out.Shard
		*out = make([]ShardSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Catalog != nil {
		in, out := &in.Catalog, &out.Catalog
		*out = make([]CatalogSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Gsm != nil {
		in, out := &in.Gsm, &out.Gsm
		*out = make([]GsmSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PortMappings != nil {
		in, out := &in.PortMappings, &out.PortMappings
		*out = make([]PortMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingDatabaseSpec.
func (in *ShardingDatabaseSpec) DeepCopy() *ShardingDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(ShardingDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingDatabaseStatus) DeepCopyInto(out *ShardingDatabaseStatus) {
	*out = *in
	if in.Shard != nil {
		in, out := &in.Shard, &out.Shard
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Catalog != nil {
		in, out := &in.Catalog, &out.Catalog
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Gsm.DeepCopyInto(&out.Gsm)
	if in.CrdStatus != nil {
		in, out := &in.CrdStatus, &out.CrdStatus
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingDatabaseStatus.
func (in *ShardingDatabaseStatus) DeepCopy() *ShardingDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(ShardingDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabase) DeepCopyInto(out *SingleInstanceDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabase.
func (in *SingleInstanceDatabase) DeepCopy() *SingleInstanceDatabase {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SingleInstanceDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabaseAdminPassword) DeepCopyInto(out *SingleInstanceDatabaseAdminPassword) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabaseAdminPassword.
func (in *SingleInstanceDatabaseAdminPassword) DeepCopy() *SingleInstanceDatabaseAdminPassword {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabaseAdminPassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabaseImage) DeepCopyInto(out *SingleInstanceDatabaseImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabaseImage.
func (in *SingleInstanceDatabaseImage) DeepCopy() *SingleInstanceDatabaseImage {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabaseImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabaseInitParams) DeepCopyInto(out *SingleInstanceDatabaseInitParams) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabaseInitParams.
func (in *SingleInstanceDatabaseInitParams) DeepCopy() *SingleInstanceDatabaseInitParams {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabaseInitParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabaseList) DeepCopyInto(out *SingleInstanceDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SingleInstanceDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabaseList.
func (in *SingleInstanceDatabaseList) DeepCopy() *SingleInstanceDatabaseList {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SingleInstanceDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabasePersistence) DeepCopyInto(out *SingleInstanceDatabasePersistence) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabasePersistence.
func (in *SingleInstanceDatabasePersistence) DeepCopy() *SingleInstanceDatabasePersistence {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabasePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabaseSpec) DeepCopyInto(out *SingleInstanceDatabaseSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.AdminPassword = in.AdminPassword
	out.Image = in.Image
	out.Persistence = in.Persistence
	out.InitParams = in.InitParams
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabaseSpec.
func (in *SingleInstanceDatabaseSpec) DeepCopy() *SingleInstanceDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleInstanceDatabaseStatus) DeepCopyInto(out *SingleInstanceDatabaseStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StandbyDatabases != nil {
		in, out := &in.StandbyDatabases, &out.StandbyDatabases
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.InitParams = in.InitParams
	out.Persistence = in.Persistence
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleInstanceDatabaseStatus.
func (in *SingleInstanceDatabaseStatus) DeepCopy() *SingleInstanceDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(SingleInstanceDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WalletSpec) DeepCopyInto(out *WalletSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WalletSpec.
func (in *WalletSpec) DeepCopy() *WalletSpec {
	if in == nil {
		return nil
	}
	out := new(WalletSpec)
	in.DeepCopyInto(out)
	return out
}
