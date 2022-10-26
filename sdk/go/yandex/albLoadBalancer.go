// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlbLoadBalancer struct {
	pulumi.CustomResourceState

	AllocationPolicy AlbLoadBalancerAllocationPolicyOutput `pulumi:"allocationPolicy"`
	CreatedAt        pulumi.StringOutput                   `pulumi:"createdAt"`
	Description      pulumi.StringPtrOutput                `pulumi:"description"`
	FolderId         pulumi.StringOutput                   `pulumi:"folderId"`
	Labels           pulumi.StringMapOutput                `pulumi:"labels"`
	Listeners        AlbLoadBalancerListenerArrayOutput    `pulumi:"listeners"`
	LogGroupId       pulumi.StringOutput                   `pulumi:"logGroupId"`
	Name             pulumi.StringOutput                   `pulumi:"name"`
	NetworkId        pulumi.StringOutput                   `pulumi:"networkId"`
	RegionId         pulumi.StringPtrOutput                `pulumi:"regionId"`
	SecurityGroupIds pulumi.StringArrayOutput              `pulumi:"securityGroupIds"`
	Status           pulumi.StringOutput                   `pulumi:"status"`
}

// NewAlbLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewAlbLoadBalancer(ctx *pulumi.Context,
	name string, args *AlbLoadBalancerArgs, opts ...pulumi.ResourceOption) (*AlbLoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllocationPolicy == nil {
		return nil, errors.New("invalid value for required argument 'AllocationPolicy'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	var resource AlbLoadBalancer
	err := ctx.RegisterResource("yandex:index/albLoadBalancer:AlbLoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlbLoadBalancer gets an existing AlbLoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlbLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlbLoadBalancerState, opts ...pulumi.ResourceOption) (*AlbLoadBalancer, error) {
	var resource AlbLoadBalancer
	err := ctx.ReadResource("yandex:index/albLoadBalancer:AlbLoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlbLoadBalancer resources.
type albLoadBalancerState struct {
	AllocationPolicy *AlbLoadBalancerAllocationPolicy `pulumi:"allocationPolicy"`
	CreatedAt        *string                          `pulumi:"createdAt"`
	Description      *string                          `pulumi:"description"`
	FolderId         *string                          `pulumi:"folderId"`
	Labels           map[string]string                `pulumi:"labels"`
	Listeners        []AlbLoadBalancerListener        `pulumi:"listeners"`
	LogGroupId       *string                          `pulumi:"logGroupId"`
	Name             *string                          `pulumi:"name"`
	NetworkId        *string                          `pulumi:"networkId"`
	RegionId         *string                          `pulumi:"regionId"`
	SecurityGroupIds []string                         `pulumi:"securityGroupIds"`
	Status           *string                          `pulumi:"status"`
}

type AlbLoadBalancerState struct {
	AllocationPolicy AlbLoadBalancerAllocationPolicyPtrInput
	CreatedAt        pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	FolderId         pulumi.StringPtrInput
	Labels           pulumi.StringMapInput
	Listeners        AlbLoadBalancerListenerArrayInput
	LogGroupId       pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	NetworkId        pulumi.StringPtrInput
	RegionId         pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
	Status           pulumi.StringPtrInput
}

func (AlbLoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*albLoadBalancerState)(nil)).Elem()
}

type albLoadBalancerArgs struct {
	AllocationPolicy AlbLoadBalancerAllocationPolicy `pulumi:"allocationPolicy"`
	Description      *string                         `pulumi:"description"`
	FolderId         *string                         `pulumi:"folderId"`
	Labels           map[string]string               `pulumi:"labels"`
	Listeners        []AlbLoadBalancerListener       `pulumi:"listeners"`
	Name             *string                         `pulumi:"name"`
	NetworkId        string                          `pulumi:"networkId"`
	RegionId         *string                         `pulumi:"regionId"`
	SecurityGroupIds []string                        `pulumi:"securityGroupIds"`
}

// The set of arguments for constructing a AlbLoadBalancer resource.
type AlbLoadBalancerArgs struct {
	AllocationPolicy AlbLoadBalancerAllocationPolicyInput
	Description      pulumi.StringPtrInput
	FolderId         pulumi.StringPtrInput
	Labels           pulumi.StringMapInput
	Listeners        AlbLoadBalancerListenerArrayInput
	Name             pulumi.StringPtrInput
	NetworkId        pulumi.StringInput
	RegionId         pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
}

func (AlbLoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*albLoadBalancerArgs)(nil)).Elem()
}

type AlbLoadBalancerInput interface {
	pulumi.Input

	ToAlbLoadBalancerOutput() AlbLoadBalancerOutput
	ToAlbLoadBalancerOutputWithContext(ctx context.Context) AlbLoadBalancerOutput
}

func (*AlbLoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**AlbLoadBalancer)(nil)).Elem()
}

func (i *AlbLoadBalancer) ToAlbLoadBalancerOutput() AlbLoadBalancerOutput {
	return i.ToAlbLoadBalancerOutputWithContext(context.Background())
}

func (i *AlbLoadBalancer) ToAlbLoadBalancerOutputWithContext(ctx context.Context) AlbLoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlbLoadBalancerOutput)
}

// AlbLoadBalancerArrayInput is an input type that accepts AlbLoadBalancerArray and AlbLoadBalancerArrayOutput values.
// You can construct a concrete instance of `AlbLoadBalancerArrayInput` via:
//
//	AlbLoadBalancerArray{ AlbLoadBalancerArgs{...} }
type AlbLoadBalancerArrayInput interface {
	pulumi.Input

	ToAlbLoadBalancerArrayOutput() AlbLoadBalancerArrayOutput
	ToAlbLoadBalancerArrayOutputWithContext(context.Context) AlbLoadBalancerArrayOutput
}

type AlbLoadBalancerArray []AlbLoadBalancerInput

func (AlbLoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlbLoadBalancer)(nil)).Elem()
}

func (i AlbLoadBalancerArray) ToAlbLoadBalancerArrayOutput() AlbLoadBalancerArrayOutput {
	return i.ToAlbLoadBalancerArrayOutputWithContext(context.Background())
}

func (i AlbLoadBalancerArray) ToAlbLoadBalancerArrayOutputWithContext(ctx context.Context) AlbLoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlbLoadBalancerArrayOutput)
}

// AlbLoadBalancerMapInput is an input type that accepts AlbLoadBalancerMap and AlbLoadBalancerMapOutput values.
// You can construct a concrete instance of `AlbLoadBalancerMapInput` via:
//
//	AlbLoadBalancerMap{ "key": AlbLoadBalancerArgs{...} }
type AlbLoadBalancerMapInput interface {
	pulumi.Input

	ToAlbLoadBalancerMapOutput() AlbLoadBalancerMapOutput
	ToAlbLoadBalancerMapOutputWithContext(context.Context) AlbLoadBalancerMapOutput
}

type AlbLoadBalancerMap map[string]AlbLoadBalancerInput

func (AlbLoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlbLoadBalancer)(nil)).Elem()
}

func (i AlbLoadBalancerMap) ToAlbLoadBalancerMapOutput() AlbLoadBalancerMapOutput {
	return i.ToAlbLoadBalancerMapOutputWithContext(context.Background())
}

func (i AlbLoadBalancerMap) ToAlbLoadBalancerMapOutputWithContext(ctx context.Context) AlbLoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlbLoadBalancerMapOutput)
}

type AlbLoadBalancerOutput struct{ *pulumi.OutputState }

func (AlbLoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlbLoadBalancer)(nil)).Elem()
}

func (o AlbLoadBalancerOutput) ToAlbLoadBalancerOutput() AlbLoadBalancerOutput {
	return o
}

func (o AlbLoadBalancerOutput) ToAlbLoadBalancerOutputWithContext(ctx context.Context) AlbLoadBalancerOutput {
	return o
}

func (o AlbLoadBalancerOutput) AllocationPolicy() AlbLoadBalancerAllocationPolicyOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) AlbLoadBalancerAllocationPolicyOutput { return v.AllocationPolicy }).(AlbLoadBalancerAllocationPolicyOutput)
}

func (o AlbLoadBalancerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o AlbLoadBalancerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AlbLoadBalancerOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

func (o AlbLoadBalancerOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o AlbLoadBalancerOutput) Listeners() AlbLoadBalancerListenerArrayOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) AlbLoadBalancerListenerArrayOutput { return v.Listeners }).(AlbLoadBalancerListenerArrayOutput)
}

func (o AlbLoadBalancerOutput) LogGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringOutput { return v.LogGroupId }).(pulumi.StringOutput)
}

func (o AlbLoadBalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AlbLoadBalancerOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

func (o AlbLoadBalancerOutput) RegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringPtrOutput { return v.RegionId }).(pulumi.StringPtrOutput)
}

func (o AlbLoadBalancerOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o AlbLoadBalancerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AlbLoadBalancer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AlbLoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (AlbLoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlbLoadBalancer)(nil)).Elem()
}

func (o AlbLoadBalancerArrayOutput) ToAlbLoadBalancerArrayOutput() AlbLoadBalancerArrayOutput {
	return o
}

func (o AlbLoadBalancerArrayOutput) ToAlbLoadBalancerArrayOutputWithContext(ctx context.Context) AlbLoadBalancerArrayOutput {
	return o
}

func (o AlbLoadBalancerArrayOutput) Index(i pulumi.IntInput) AlbLoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlbLoadBalancer {
		return vs[0].([]*AlbLoadBalancer)[vs[1].(int)]
	}).(AlbLoadBalancerOutput)
}

type AlbLoadBalancerMapOutput struct{ *pulumi.OutputState }

func (AlbLoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlbLoadBalancer)(nil)).Elem()
}

func (o AlbLoadBalancerMapOutput) ToAlbLoadBalancerMapOutput() AlbLoadBalancerMapOutput {
	return o
}

func (o AlbLoadBalancerMapOutput) ToAlbLoadBalancerMapOutputWithContext(ctx context.Context) AlbLoadBalancerMapOutput {
	return o
}

func (o AlbLoadBalancerMapOutput) MapIndex(k pulumi.StringInput) AlbLoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlbLoadBalancer {
		return vs[0].(map[string]*AlbLoadBalancer)[vs[1].(string)]
	}).(AlbLoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlbLoadBalancerInput)(nil)).Elem(), &AlbLoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlbLoadBalancerArrayInput)(nil)).Elem(), AlbLoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlbLoadBalancerMapInput)(nil)).Elem(), AlbLoadBalancerMap{})
	pulumi.RegisterOutputType(AlbLoadBalancerOutput{})
	pulumi.RegisterOutputType(AlbLoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(AlbLoadBalancerMapOutput{})
}