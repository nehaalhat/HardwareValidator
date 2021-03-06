// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpectedCPU) DeepCopyInto(out *ExpectedCPU) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpectedCPU.
func (in *ExpectedCPU) DeepCopy() *ExpectedCPU {
	if in == nil {
		return nil
	}
	out := new(ExpectedCPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpectedDisk) DeepCopyInto(out *ExpectedDisk) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpectedDisk.
func (in *ExpectedDisk) DeepCopy() *ExpectedDisk {
	if in == nil {
		return nil
	}
	out := new(ExpectedDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpectedHardwareConfiguration) DeepCopyInto(out *ExpectedHardwareConfiguration) {
	*out = *in
	out.ExpectedCPU = in.ExpectedCPU
	out.ExpectedDisk = in.ExpectedDisk
	out.ExpectedNICS = in.ExpectedNICS
	out.ExpectedSystemVendor = in.ExpectedSystemVendor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpectedHardwareConfiguration.
func (in *ExpectedHardwareConfiguration) DeepCopy() *ExpectedHardwareConfiguration {
	if in == nil {
		return nil
	}
	out := new(ExpectedHardwareConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpectedNICS) DeepCopyInto(out *ExpectedNICS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpectedNICS.
func (in *ExpectedNICS) DeepCopy() *ExpectedNICS {
	if in == nil {
		return nil
	}
	out := new(ExpectedNICS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpectedSystemVendor) DeepCopyInto(out *ExpectedSystemVendor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpectedSystemVendor.
func (in *ExpectedSystemVendor) DeepCopy() *ExpectedSystemVendor {
	if in == nil {
		return nil
	}
	out := new(ExpectedSystemVendor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareClassificationController) DeepCopyInto(out *HardwareClassificationController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareClassificationController.
func (in *HardwareClassificationController) DeepCopy() *HardwareClassificationController {
	if in == nil {
		return nil
	}
	out := new(HardwareClassificationController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HardwareClassificationController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareClassificationControllerList) DeepCopyInto(out *HardwareClassificationControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HardwareClassificationController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareClassificationControllerList.
func (in *HardwareClassificationControllerList) DeepCopy() *HardwareClassificationControllerList {
	if in == nil {
		return nil
	}
	out := new(HardwareClassificationControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HardwareClassificationControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareClassificationControllerSpec) DeepCopyInto(out *HardwareClassificationControllerSpec) {
	*out = *in
	if in.ExpectedHardwareConfiguration != nil {
		in, out := &in.ExpectedHardwareConfiguration, &out.ExpectedHardwareConfiguration
		*out = make([]ExpectedHardwareConfiguration, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareClassificationControllerSpec.
func (in *HardwareClassificationControllerSpec) DeepCopy() *HardwareClassificationControllerSpec {
	if in == nil {
		return nil
	}
	out := new(HardwareClassificationControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareClassificationControllerStatus) DeepCopyInto(out *HardwareClassificationControllerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareClassificationControllerStatus.
func (in *HardwareClassificationControllerStatus) DeepCopy() *HardwareClassificationControllerStatus {
	if in == nil {
		return nil
	}
	out := new(HardwareClassificationControllerStatus)
	in.DeepCopyInto(out)
	return out
}
