// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Compute instance. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/compute/concepts/vm).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myInstance, err := yandex.LookupComputeInstance(ctx, &GetComputeInstanceArgs{
//				InstanceId: pulumi.StringRef("some_instance_id"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("instanceExternalIp", myInstance.NetworkInterfaces[0].NatIpAddress)
//			return nil
//		})
//	}
//
// ```
func LookupComputeInstance(ctx *pulumi.Context, args *LookupComputeInstanceArgs, opts ...pulumi.InvokeOption) (*LookupComputeInstanceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupComputeInstanceResult
	err := ctx.Invoke("yandex:index/getComputeInstance:getComputeInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeInstance.
type LookupComputeInstanceArgs struct {
	Filesystems []GetComputeInstanceFilesystem `pulumi:"filesystems"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// ID of GPU cluster if instance is part of it.
	GpuClusterId *string `pulumi:"gpuClusterId"`
	// The ID of a specific instance.
	InstanceId *string `pulumi:"instanceId"`
	// List of local disks that are attached to the instance. Structure is documented below.
	LocalDisks []GetComputeInstanceLocalDisk `pulumi:"localDisks"`
	// Options allow user to configure access to instance's metadata
	MetadataOptions *GetComputeInstanceMetadataOptions `pulumi:"metadataOptions"`
	// Name of the instance.
	Name *string `pulumi:"name"`
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy *GetComputeInstancePlacementPolicy `pulumi:"placementPolicy"`
}

// A collection of values returned by getComputeInstance.
type LookupComputeInstanceResult struct {
	// The boot disk for the instance. Structure is documented below.
	BootDisks []GetComputeInstanceBootDisk `pulumi:"bootDisks"`
	// Instance creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// Description of the boot disk.
	Description string                         `pulumi:"description"`
	Filesystems []GetComputeInstanceFilesystem `pulumi:"filesystems"`
	FolderId    string                         `pulumi:"folderId"`
	// DNS record FQDN.
	Fqdn string `pulumi:"fqdn"`
	// ID of GPU cluster if instance is part of it.
	GpuClusterId string `pulumi:"gpuClusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// A set of key/value label pairs assigned to the instance.
	Labels map[string]string `pulumi:"labels"`
	// List of local disks that are attached to the instance. Structure is documented below.
	LocalDisks []GetComputeInstanceLocalDisk `pulumi:"localDisks"`
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// Options allow user to configure access to instance's metadata
	MetadataOptions GetComputeInstanceMetadataOptions `pulumi:"metadataOptions"`
	// Name of the boot disk.
	Name string `pulumi:"name"`
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType string `pulumi:"networkAccelerationType"`
	// The networks attached to the instance. Structure is documented below.
	// * `network_interface.0.ip_address` - An internal IP address of the instance, either manually or dynamically assigned.
	// * `network_interface.0.nat_ip_address` - An assigned external IP address if the instance has NAT enabled.
	NetworkInterfaces []GetComputeInstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy *GetComputeInstancePlacementPolicy `pulumi:"placementPolicy"`
	// Type of virtual machine to create. Default is 'standard-v1'.
	PlatformId string                       `pulumi:"platformId"`
	Resources  []GetComputeInstanceResource `pulumi:"resources"`
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicies []GetComputeInstanceSchedulingPolicy `pulumi:"schedulingPolicies"`
	// List of secondary disks attached to the instance. Structure is documented below.
	SecondaryDisks []GetComputeInstanceSecondaryDisk `pulumi:"secondaryDisks"`
	// ID of the service account authorized for this instance.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// Status of the instance.
	// * `resources.0.memory` - Memory size allocated for the instance.
	// * `resources.0.cores` - Number of CPU cores allocated for the instance.
	// * `resources.0.core_fraction` - Baseline performance for a core, set as a percent.
	// * `resources.0.gpus` - Number of GPU cores allocated for the instance.
	Status string `pulumi:"status"`
	// Availability zone where the instance resides.
	Zone string `pulumi:"zone"`
}

func LookupComputeInstanceOutput(ctx *pulumi.Context, args LookupComputeInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupComputeInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputeInstanceResult, error) {
			args := v.(LookupComputeInstanceArgs)
			r, err := LookupComputeInstance(ctx, &args, opts...)
			var s LookupComputeInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputeInstanceResultOutput)
}

// A collection of arguments for invoking getComputeInstance.
type LookupComputeInstanceOutputArgs struct {
	Filesystems GetComputeInstanceFilesystemArrayInput `pulumi:"filesystems"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// ID of GPU cluster if instance is part of it.
	GpuClusterId pulumi.StringPtrInput `pulumi:"gpuClusterId"`
	// The ID of a specific instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// List of local disks that are attached to the instance. Structure is documented below.
	LocalDisks GetComputeInstanceLocalDiskArrayInput `pulumi:"localDisks"`
	// Options allow user to configure access to instance's metadata
	MetadataOptions GetComputeInstanceMetadataOptionsPtrInput `pulumi:"metadataOptions"`
	// Name of the instance.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy GetComputeInstancePlacementPolicyPtrInput `pulumi:"placementPolicy"`
}

func (LookupComputeInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getComputeInstance.
type LookupComputeInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupComputeInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeInstanceResult)(nil)).Elem()
}

func (o LookupComputeInstanceResultOutput) ToLookupComputeInstanceResultOutput() LookupComputeInstanceResultOutput {
	return o
}

func (o LookupComputeInstanceResultOutput) ToLookupComputeInstanceResultOutputWithContext(ctx context.Context) LookupComputeInstanceResultOutput {
	return o
}

// The boot disk for the instance. Structure is documented below.
func (o LookupComputeInstanceResultOutput) BootDisks() GetComputeInstanceBootDiskArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceBootDisk { return v.BootDisks }).(GetComputeInstanceBootDiskArrayOutput)
}

// Instance creation timestamp.
func (o LookupComputeInstanceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the boot disk.
func (o LookupComputeInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceResultOutput) Filesystems() GetComputeInstanceFilesystemArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceFilesystem { return v.Filesystems }).(GetComputeInstanceFilesystemArrayOutput)
}

func (o LookupComputeInstanceResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// DNS record FQDN.
func (o LookupComputeInstanceResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// ID of GPU cluster if instance is part of it.
func (o LookupComputeInstanceResultOutput) GpuClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.GpuClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupComputeInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// A set of key/value label pairs assigned to the instance.
func (o LookupComputeInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// List of local disks that are attached to the instance. Structure is documented below.
func (o LookupComputeInstanceResultOutput) LocalDisks() GetComputeInstanceLocalDiskArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceLocalDisk { return v.LocalDisks }).(GetComputeInstanceLocalDiskArrayOutput)
}

// Metadata key/value pairs to make available from
// within the instance.
func (o LookupComputeInstanceResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Options allow user to configure access to instance's metadata
func (o LookupComputeInstanceResultOutput) MetadataOptions() GetComputeInstanceMetadataOptionsOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) GetComputeInstanceMetadataOptions { return v.MetadataOptions }).(GetComputeInstanceMetadataOptionsOutput)
}

// Name of the boot disk.
func (o LookupComputeInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
func (o LookupComputeInstanceResultOutput) NetworkAccelerationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.NetworkAccelerationType }).(pulumi.StringOutput)
}

// The networks attached to the instance. Structure is documented below.
// * `network_interface.0.ip_address` - An internal IP address of the instance, either manually or dynamically assigned.
// * `network_interface.0.nat_ip_address` - An assigned external IP address if the instance has NAT enabled.
func (o LookupComputeInstanceResultOutput) NetworkInterfaces() GetComputeInstanceNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceNetworkInterface { return v.NetworkInterfaces }).(GetComputeInstanceNetworkInterfaceArrayOutput)
}

// The placement policy configuration. The structure is documented below.
func (o LookupComputeInstanceResultOutput) PlacementPolicy() GetComputeInstancePlacementPolicyPtrOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) *GetComputeInstancePlacementPolicy { return v.PlacementPolicy }).(GetComputeInstancePlacementPolicyPtrOutput)
}

// Type of virtual machine to create. Default is 'standard-v1'.
func (o LookupComputeInstanceResultOutput) PlatformId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.PlatformId }).(pulumi.StringOutput)
}

func (o LookupComputeInstanceResultOutput) Resources() GetComputeInstanceResourceArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceResource { return v.Resources }).(GetComputeInstanceResourceArrayOutput)
}

// Scheduling policy configuration. The structure is documented below.
func (o LookupComputeInstanceResultOutput) SchedulingPolicies() GetComputeInstanceSchedulingPolicyArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceSchedulingPolicy { return v.SchedulingPolicies }).(GetComputeInstanceSchedulingPolicyArrayOutput)
}

// List of secondary disks attached to the instance. Structure is documented below.
func (o LookupComputeInstanceResultOutput) SecondaryDisks() GetComputeInstanceSecondaryDiskArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []GetComputeInstanceSecondaryDisk { return v.SecondaryDisks }).(GetComputeInstanceSecondaryDiskArrayOutput)
}

// ID of the service account authorized for this instance.
func (o LookupComputeInstanceResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// Status of the instance.
// * `resources.0.memory` - Memory size allocated for the instance.
// * `resources.0.cores` - Number of CPU cores allocated for the instance.
// * `resources.0.core_fraction` - Baseline performance for a core, set as a percent.
// * `resources.0.gpus` - Number of GPU cores allocated for the instance.
func (o LookupComputeInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

// Availability zone where the instance resides.
func (o LookupComputeInstanceResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeInstanceResultOutput{})
}
