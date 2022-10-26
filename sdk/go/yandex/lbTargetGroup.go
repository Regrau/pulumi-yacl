// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LbTargetGroup struct {
	pulumi.CustomResourceState

	CreatedAt   pulumi.StringOutput            `pulumi:"createdAt"`
	Description pulumi.StringPtrOutput         `pulumi:"description"`
	FolderId    pulumi.StringOutput            `pulumi:"folderId"`
	Labels      pulumi.StringMapOutput         `pulumi:"labels"`
	Name        pulumi.StringOutput            `pulumi:"name"`
	RegionId    pulumi.StringPtrOutput         `pulumi:"regionId"`
	Targets     LbTargetGroupTargetArrayOutput `pulumi:"targets"`
}

// NewLbTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewLbTargetGroup(ctx *pulumi.Context,
	name string, args *LbTargetGroupArgs, opts ...pulumi.ResourceOption) (*LbTargetGroup, error) {
	if args == nil {
		args = &LbTargetGroupArgs{}
	}

	var resource LbTargetGroup
	err := ctx.RegisterResource("yandex:index/lbTargetGroup:LbTargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbTargetGroup gets an existing LbTargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbTargetGroupState, opts ...pulumi.ResourceOption) (*LbTargetGroup, error) {
	var resource LbTargetGroup
	err := ctx.ReadResource("yandex:index/lbTargetGroup:LbTargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbTargetGroup resources.
type lbTargetGroupState struct {
	CreatedAt   *string               `pulumi:"createdAt"`
	Description *string               `pulumi:"description"`
	FolderId    *string               `pulumi:"folderId"`
	Labels      map[string]string     `pulumi:"labels"`
	Name        *string               `pulumi:"name"`
	RegionId    *string               `pulumi:"regionId"`
	Targets     []LbTargetGroupTarget `pulumi:"targets"`
}

type LbTargetGroupState struct {
	CreatedAt   pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	FolderId    pulumi.StringPtrInput
	Labels      pulumi.StringMapInput
	Name        pulumi.StringPtrInput
	RegionId    pulumi.StringPtrInput
	Targets     LbTargetGroupTargetArrayInput
}

func (LbTargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbTargetGroupState)(nil)).Elem()
}

type lbTargetGroupArgs struct {
	Description *string               `pulumi:"description"`
	FolderId    *string               `pulumi:"folderId"`
	Labels      map[string]string     `pulumi:"labels"`
	Name        *string               `pulumi:"name"`
	RegionId    *string               `pulumi:"regionId"`
	Targets     []LbTargetGroupTarget `pulumi:"targets"`
}

// The set of arguments for constructing a LbTargetGroup resource.
type LbTargetGroupArgs struct {
	Description pulumi.StringPtrInput
	FolderId    pulumi.StringPtrInput
	Labels      pulumi.StringMapInput
	Name        pulumi.StringPtrInput
	RegionId    pulumi.StringPtrInput
	Targets     LbTargetGroupTargetArrayInput
}

func (LbTargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbTargetGroupArgs)(nil)).Elem()
}

type LbTargetGroupInput interface {
	pulumi.Input

	ToLbTargetGroupOutput() LbTargetGroupOutput
	ToLbTargetGroupOutputWithContext(ctx context.Context) LbTargetGroupOutput
}

func (*LbTargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**LbTargetGroup)(nil)).Elem()
}

func (i *LbTargetGroup) ToLbTargetGroupOutput() LbTargetGroupOutput {
	return i.ToLbTargetGroupOutputWithContext(context.Background())
}

func (i *LbTargetGroup) ToLbTargetGroupOutputWithContext(ctx context.Context) LbTargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbTargetGroupOutput)
}

// LbTargetGroupArrayInput is an input type that accepts LbTargetGroupArray and LbTargetGroupArrayOutput values.
// You can construct a concrete instance of `LbTargetGroupArrayInput` via:
//
//	LbTargetGroupArray{ LbTargetGroupArgs{...} }
type LbTargetGroupArrayInput interface {
	pulumi.Input

	ToLbTargetGroupArrayOutput() LbTargetGroupArrayOutput
	ToLbTargetGroupArrayOutputWithContext(context.Context) LbTargetGroupArrayOutput
}

type LbTargetGroupArray []LbTargetGroupInput

func (LbTargetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbTargetGroup)(nil)).Elem()
}

func (i LbTargetGroupArray) ToLbTargetGroupArrayOutput() LbTargetGroupArrayOutput {
	return i.ToLbTargetGroupArrayOutputWithContext(context.Background())
}

func (i LbTargetGroupArray) ToLbTargetGroupArrayOutputWithContext(ctx context.Context) LbTargetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbTargetGroupArrayOutput)
}

// LbTargetGroupMapInput is an input type that accepts LbTargetGroupMap and LbTargetGroupMapOutput values.
// You can construct a concrete instance of `LbTargetGroupMapInput` via:
//
//	LbTargetGroupMap{ "key": LbTargetGroupArgs{...} }
type LbTargetGroupMapInput interface {
	pulumi.Input

	ToLbTargetGroupMapOutput() LbTargetGroupMapOutput
	ToLbTargetGroupMapOutputWithContext(context.Context) LbTargetGroupMapOutput
}

type LbTargetGroupMap map[string]LbTargetGroupInput

func (LbTargetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbTargetGroup)(nil)).Elem()
}

func (i LbTargetGroupMap) ToLbTargetGroupMapOutput() LbTargetGroupMapOutput {
	return i.ToLbTargetGroupMapOutputWithContext(context.Background())
}

func (i LbTargetGroupMap) ToLbTargetGroupMapOutputWithContext(ctx context.Context) LbTargetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbTargetGroupMapOutput)
}

type LbTargetGroupOutput struct{ *pulumi.OutputState }

func (LbTargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbTargetGroup)(nil)).Elem()
}

func (o LbTargetGroupOutput) ToLbTargetGroupOutput() LbTargetGroupOutput {
	return o
}

func (o LbTargetGroupOutput) ToLbTargetGroupOutputWithContext(ctx context.Context) LbTargetGroupOutput {
	return o
}

func (o LbTargetGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LbTargetGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LbTargetGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbTargetGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LbTargetGroupOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbTargetGroup) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

func (o LbTargetGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LbTargetGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LbTargetGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbTargetGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LbTargetGroupOutput) RegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbTargetGroup) pulumi.StringPtrOutput { return v.RegionId }).(pulumi.StringPtrOutput)
}

func (o LbTargetGroupOutput) Targets() LbTargetGroupTargetArrayOutput {
	return o.ApplyT(func(v *LbTargetGroup) LbTargetGroupTargetArrayOutput { return v.Targets }).(LbTargetGroupTargetArrayOutput)
}

type LbTargetGroupArrayOutput struct{ *pulumi.OutputState }

func (LbTargetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbTargetGroup)(nil)).Elem()
}

func (o LbTargetGroupArrayOutput) ToLbTargetGroupArrayOutput() LbTargetGroupArrayOutput {
	return o
}

func (o LbTargetGroupArrayOutput) ToLbTargetGroupArrayOutputWithContext(ctx context.Context) LbTargetGroupArrayOutput {
	return o
}

func (o LbTargetGroupArrayOutput) Index(i pulumi.IntInput) LbTargetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbTargetGroup {
		return vs[0].([]*LbTargetGroup)[vs[1].(int)]
	}).(LbTargetGroupOutput)
}

type LbTargetGroupMapOutput struct{ *pulumi.OutputState }

func (LbTargetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbTargetGroup)(nil)).Elem()
}

func (o LbTargetGroupMapOutput) ToLbTargetGroupMapOutput() LbTargetGroupMapOutput {
	return o
}

func (o LbTargetGroupMapOutput) ToLbTargetGroupMapOutputWithContext(ctx context.Context) LbTargetGroupMapOutput {
	return o
}

func (o LbTargetGroupMapOutput) MapIndex(k pulumi.StringInput) LbTargetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbTargetGroup {
		return vs[0].(map[string]*LbTargetGroup)[vs[1].(string)]
	}).(LbTargetGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbTargetGroupInput)(nil)).Elem(), &LbTargetGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbTargetGroupArrayInput)(nil)).Elem(), LbTargetGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbTargetGroupMapInput)(nil)).Elem(), LbTargetGroupMap{})
	pulumi.RegisterOutputType(LbTargetGroupOutput{})
	pulumi.RegisterOutputType(LbTargetGroupArrayOutput{})
	pulumi.RegisterOutputType(LbTargetGroupMapOutput{})
}