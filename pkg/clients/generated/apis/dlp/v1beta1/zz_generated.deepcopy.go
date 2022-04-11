//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLPStoredInfoType) DeepCopyInto(out *DLPStoredInfoType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLPStoredInfoType.
func (in *DLPStoredInfoType) DeepCopy() *DLPStoredInfoType {
	if in == nil {
		return nil
	}
	out := new(DLPStoredInfoType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DLPStoredInfoType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLPStoredInfoTypeList) DeepCopyInto(out *DLPStoredInfoTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DLPStoredInfoType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLPStoredInfoTypeList.
func (in *DLPStoredInfoTypeList) DeepCopy() *DLPStoredInfoTypeList {
	if in == nil {
		return nil
	}
	out := new(DLPStoredInfoTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DLPStoredInfoTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLPStoredInfoTypeSpec) DeepCopyInto(out *DLPStoredInfoTypeSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dictionary != nil {
		in, out := &in.Dictionary, &out.Dictionary
		*out = new(StoredinfotypeDictionary)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.LargeCustomDictionary != nil {
		in, out := &in.LargeCustomDictionary, &out.LargeCustomDictionary
		*out = new(StoredinfotypeLargeCustomDictionary)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(StoredinfotypeRegex)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLPStoredInfoTypeSpec.
func (in *DLPStoredInfoTypeSpec) DeepCopy() *DLPStoredInfoTypeSpec {
	if in == nil {
		return nil
	}
	out := new(DLPStoredInfoTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLPStoredInfoTypeStatus) DeepCopyInto(out *DLPStoredInfoTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLPStoredInfoTypeStatus.
func (in *DLPStoredInfoTypeStatus) DeepCopy() *DLPStoredInfoTypeStatus {
	if in == nil {
		return nil
	}
	out := new(DLPStoredInfoTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeBigQueryField) DeepCopyInto(out *StoredinfotypeBigQueryField) {
	*out = *in
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(StoredinfotypeField)
		(*in).DeepCopyInto(*out)
	}
	if in.Table != nil {
		in, out := &in.Table, &out.Table
		*out = new(StoredinfotypeTable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeBigQueryField.
func (in *StoredinfotypeBigQueryField) DeepCopy() *StoredinfotypeBigQueryField {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeBigQueryField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeCloudStorageFileSet) DeepCopyInto(out *StoredinfotypeCloudStorageFileSet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeCloudStorageFileSet.
func (in *StoredinfotypeCloudStorageFileSet) DeepCopy() *StoredinfotypeCloudStorageFileSet {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeCloudStorageFileSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeCloudStoragePath) DeepCopyInto(out *StoredinfotypeCloudStoragePath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeCloudStoragePath.
func (in *StoredinfotypeCloudStoragePath) DeepCopy() *StoredinfotypeCloudStoragePath {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeCloudStoragePath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeDictionary) DeepCopyInto(out *StoredinfotypeDictionary) {
	*out = *in
	if in.CloudStoragePath != nil {
		in, out := &in.CloudStoragePath, &out.CloudStoragePath
		*out = new(StoredinfotypeCloudStoragePath)
		**out = **in
	}
	if in.WordList != nil {
		in, out := &in.WordList, &out.WordList
		*out = new(StoredinfotypeWordList)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeDictionary.
func (in *StoredinfotypeDictionary) DeepCopy() *StoredinfotypeDictionary {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeDictionary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeField) DeepCopyInto(out *StoredinfotypeField) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeField.
func (in *StoredinfotypeField) DeepCopy() *StoredinfotypeField {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeLargeCustomDictionary) DeepCopyInto(out *StoredinfotypeLargeCustomDictionary) {
	*out = *in
	if in.BigQueryField != nil {
		in, out := &in.BigQueryField, &out.BigQueryField
		*out = new(StoredinfotypeBigQueryField)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudStorageFileSet != nil {
		in, out := &in.CloudStorageFileSet, &out.CloudStorageFileSet
		*out = new(StoredinfotypeCloudStorageFileSet)
		**out = **in
	}
	if in.OutputPath != nil {
		in, out := &in.OutputPath, &out.OutputPath
		*out = new(StoredinfotypeOutputPath)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeLargeCustomDictionary.
func (in *StoredinfotypeLargeCustomDictionary) DeepCopy() *StoredinfotypeLargeCustomDictionary {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeLargeCustomDictionary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeOutputPath) DeepCopyInto(out *StoredinfotypeOutputPath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeOutputPath.
func (in *StoredinfotypeOutputPath) DeepCopy() *StoredinfotypeOutputPath {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeOutputPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeRegex) DeepCopyInto(out *StoredinfotypeRegex) {
	*out = *in
	if in.GroupIndexes != nil {
		in, out := &in.GroupIndexes, &out.GroupIndexes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeRegex.
func (in *StoredinfotypeRegex) DeepCopy() *StoredinfotypeRegex {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeRegex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeTable) DeepCopyInto(out *StoredinfotypeTable) {
	*out = *in
	if in.DatasetRef != nil {
		in, out := &in.DatasetRef, &out.DatasetRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.TableRef != nil {
		in, out := &in.TableRef, &out.TableRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeTable.
func (in *StoredinfotypeTable) DeepCopy() *StoredinfotypeTable {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoredinfotypeWordList) DeepCopyInto(out *StoredinfotypeWordList) {
	*out = *in
	if in.Words != nil {
		in, out := &in.Words, &out.Words
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoredinfotypeWordList.
func (in *StoredinfotypeWordList) DeepCopy() *StoredinfotypeWordList {
	if in == nil {
		return nil
	}
	out := new(StoredinfotypeWordList)
	in.DeepCopyInto(out)
	return out
}
