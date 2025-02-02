// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type YandexYdbTopicIndex struct {
	pulumi.CustomResourceState

	Columns          pulumi.StringArrayOutput `pulumi:"columns"`
	ConnectionString pulumi.StringOutput      `pulumi:"connectionString"`
	Covers           pulumi.StringArrayOutput `pulumi:"covers"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	TableId          pulumi.StringOutput      `pulumi:"tableId"`
	TablePath        pulumi.StringOutput      `pulumi:"tablePath"`
	Type             pulumi.StringOutput      `pulumi:"type"`
}

// NewYandexYdbTopicIndex registers a new resource with the given unique name, arguments, and options.
func NewYandexYdbTopicIndex(ctx *pulumi.Context,
	name string, args *YandexYdbTopicIndexArgs, opts ...pulumi.ResourceOption) (*YandexYdbTopicIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Columns == nil {
		return nil, errors.New("invalid value for required argument 'Columns'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource YandexYdbTopicIndex
	err := ctx.RegisterResource("yandex:index/yandexYdbTopicIndex:yandexYdbTopicIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetYandexYdbTopicIndex gets an existing YandexYdbTopicIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetYandexYdbTopicIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *YandexYdbTopicIndexState, opts ...pulumi.ResourceOption) (*YandexYdbTopicIndex, error) {
	var resource YandexYdbTopicIndex
	err := ctx.ReadResource("yandex:index/yandexYdbTopicIndex:yandexYdbTopicIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering YandexYdbTopicIndex resources.
type yandexYdbTopicIndexState struct {
	Columns          []string `pulumi:"columns"`
	ConnectionString *string  `pulumi:"connectionString"`
	Covers           []string `pulumi:"covers"`
	Name             *string  `pulumi:"name"`
	TableId          *string  `pulumi:"tableId"`
	TablePath        *string  `pulumi:"tablePath"`
	Type             *string  `pulumi:"type"`
}

type YandexYdbTopicIndexState struct {
	Columns          pulumi.StringArrayInput
	ConnectionString pulumi.StringPtrInput
	Covers           pulumi.StringArrayInput
	Name             pulumi.StringPtrInput
	TableId          pulumi.StringPtrInput
	TablePath        pulumi.StringPtrInput
	Type             pulumi.StringPtrInput
}

func (YandexYdbTopicIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*yandexYdbTopicIndexState)(nil)).Elem()
}

type yandexYdbTopicIndexArgs struct {
	Columns          []string `pulumi:"columns"`
	ConnectionString *string  `pulumi:"connectionString"`
	Covers           []string `pulumi:"covers"`
	Name             *string  `pulumi:"name"`
	TableId          *string  `pulumi:"tableId"`
	TablePath        *string  `pulumi:"tablePath"`
	Type             string   `pulumi:"type"`
}

// The set of arguments for constructing a YandexYdbTopicIndex resource.
type YandexYdbTopicIndexArgs struct {
	Columns          pulumi.StringArrayInput
	ConnectionString pulumi.StringPtrInput
	Covers           pulumi.StringArrayInput
	Name             pulumi.StringPtrInput
	TableId          pulumi.StringPtrInput
	TablePath        pulumi.StringPtrInput
	Type             pulumi.StringInput
}

func (YandexYdbTopicIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*yandexYdbTopicIndexArgs)(nil)).Elem()
}

type YandexYdbTopicIndexInput interface {
	pulumi.Input

	ToYandexYdbTopicIndexOutput() YandexYdbTopicIndexOutput
	ToYandexYdbTopicIndexOutputWithContext(ctx context.Context) YandexYdbTopicIndexOutput
}

func (*YandexYdbTopicIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**YandexYdbTopicIndex)(nil)).Elem()
}

func (i *YandexYdbTopicIndex) ToYandexYdbTopicIndexOutput() YandexYdbTopicIndexOutput {
	return i.ToYandexYdbTopicIndexOutputWithContext(context.Background())
}

func (i *YandexYdbTopicIndex) ToYandexYdbTopicIndexOutputWithContext(ctx context.Context) YandexYdbTopicIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandexYdbTopicIndexOutput)
}

// YandexYdbTopicIndexArrayInput is an input type that accepts YandexYdbTopicIndexArray and YandexYdbTopicIndexArrayOutput values.
// You can construct a concrete instance of `YandexYdbTopicIndexArrayInput` via:
//
//	YandexYdbTopicIndexArray{ YandexYdbTopicIndexArgs{...} }
type YandexYdbTopicIndexArrayInput interface {
	pulumi.Input

	ToYandexYdbTopicIndexArrayOutput() YandexYdbTopicIndexArrayOutput
	ToYandexYdbTopicIndexArrayOutputWithContext(context.Context) YandexYdbTopicIndexArrayOutput
}

type YandexYdbTopicIndexArray []YandexYdbTopicIndexInput

func (YandexYdbTopicIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*YandexYdbTopicIndex)(nil)).Elem()
}

func (i YandexYdbTopicIndexArray) ToYandexYdbTopicIndexArrayOutput() YandexYdbTopicIndexArrayOutput {
	return i.ToYandexYdbTopicIndexArrayOutputWithContext(context.Background())
}

func (i YandexYdbTopicIndexArray) ToYandexYdbTopicIndexArrayOutputWithContext(ctx context.Context) YandexYdbTopicIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandexYdbTopicIndexArrayOutput)
}

// YandexYdbTopicIndexMapInput is an input type that accepts YandexYdbTopicIndexMap and YandexYdbTopicIndexMapOutput values.
// You can construct a concrete instance of `YandexYdbTopicIndexMapInput` via:
//
//	YandexYdbTopicIndexMap{ "key": YandexYdbTopicIndexArgs{...} }
type YandexYdbTopicIndexMapInput interface {
	pulumi.Input

	ToYandexYdbTopicIndexMapOutput() YandexYdbTopicIndexMapOutput
	ToYandexYdbTopicIndexMapOutputWithContext(context.Context) YandexYdbTopicIndexMapOutput
}

type YandexYdbTopicIndexMap map[string]YandexYdbTopicIndexInput

func (YandexYdbTopicIndexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*YandexYdbTopicIndex)(nil)).Elem()
}

func (i YandexYdbTopicIndexMap) ToYandexYdbTopicIndexMapOutput() YandexYdbTopicIndexMapOutput {
	return i.ToYandexYdbTopicIndexMapOutputWithContext(context.Background())
}

func (i YandexYdbTopicIndexMap) ToYandexYdbTopicIndexMapOutputWithContext(ctx context.Context) YandexYdbTopicIndexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandexYdbTopicIndexMapOutput)
}

type YandexYdbTopicIndexOutput struct{ *pulumi.OutputState }

func (YandexYdbTopicIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YandexYdbTopicIndex)(nil)).Elem()
}

func (o YandexYdbTopicIndexOutput) ToYandexYdbTopicIndexOutput() YandexYdbTopicIndexOutput {
	return o
}

func (o YandexYdbTopicIndexOutput) ToYandexYdbTopicIndexOutputWithContext(ctx context.Context) YandexYdbTopicIndexOutput {
	return o
}

func (o YandexYdbTopicIndexOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringArrayOutput { return v.Columns }).(pulumi.StringArrayOutput)
}

func (o YandexYdbTopicIndexOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o YandexYdbTopicIndexOutput) Covers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringArrayOutput { return v.Covers }).(pulumi.StringArrayOutput)
}

func (o YandexYdbTopicIndexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o YandexYdbTopicIndexOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringOutput { return v.TableId }).(pulumi.StringOutput)
}

func (o YandexYdbTopicIndexOutput) TablePath() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringOutput { return v.TablePath }).(pulumi.StringOutput)
}

func (o YandexYdbTopicIndexOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexYdbTopicIndex) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type YandexYdbTopicIndexArrayOutput struct{ *pulumi.OutputState }

func (YandexYdbTopicIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*YandexYdbTopicIndex)(nil)).Elem()
}

func (o YandexYdbTopicIndexArrayOutput) ToYandexYdbTopicIndexArrayOutput() YandexYdbTopicIndexArrayOutput {
	return o
}

func (o YandexYdbTopicIndexArrayOutput) ToYandexYdbTopicIndexArrayOutputWithContext(ctx context.Context) YandexYdbTopicIndexArrayOutput {
	return o
}

func (o YandexYdbTopicIndexArrayOutput) Index(i pulumi.IntInput) YandexYdbTopicIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *YandexYdbTopicIndex {
		return vs[0].([]*YandexYdbTopicIndex)[vs[1].(int)]
	}).(YandexYdbTopicIndexOutput)
}

type YandexYdbTopicIndexMapOutput struct{ *pulumi.OutputState }

func (YandexYdbTopicIndexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*YandexYdbTopicIndex)(nil)).Elem()
}

func (o YandexYdbTopicIndexMapOutput) ToYandexYdbTopicIndexMapOutput() YandexYdbTopicIndexMapOutput {
	return o
}

func (o YandexYdbTopicIndexMapOutput) ToYandexYdbTopicIndexMapOutputWithContext(ctx context.Context) YandexYdbTopicIndexMapOutput {
	return o
}

func (o YandexYdbTopicIndexMapOutput) MapIndex(k pulumi.StringInput) YandexYdbTopicIndexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *YandexYdbTopicIndex {
		return vs[0].(map[string]*YandexYdbTopicIndex)[vs[1].(string)]
	}).(YandexYdbTopicIndexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*YandexYdbTopicIndexInput)(nil)).Elem(), &YandexYdbTopicIndex{})
	pulumi.RegisterInputType(reflect.TypeOf((*YandexYdbTopicIndexArrayInput)(nil)).Elem(), YandexYdbTopicIndexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*YandexYdbTopicIndexMapInput)(nil)).Elem(), YandexYdbTopicIndexMap{})
	pulumi.RegisterOutputType(YandexYdbTopicIndexOutput{})
	pulumi.RegisterOutputType(YandexYdbTopicIndexArrayOutput{})
	pulumi.RegisterOutputType(YandexYdbTopicIndexMapOutput{})
}
