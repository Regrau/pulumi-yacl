// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creating the bind, which connect the cloud to the billing account.
// For more information, see [Cloud binding](https://cloud.yandex.ru/docs/billing/operations/pin-cloud).
//
// **Note**: Currently resource deletion do not unbind cloud from billing account. Instead it does no-operations.
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
//			_, err := yandex.NewyandexBillingCloudBinding(ctx, "foo", &yandex.yandexBillingCloudBindingArgs{
//				BillingAccountId: pulumi.String("foo-ba-id"),
//				CloudId:          pulumi.String("foo-cloud-id"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Cloud binding can be imported by ID
//
// ```sh
//
//	$ pulumi import yandex:index/yandexBillingCloudBinding:yandexBillingCloudBinding foo cloud-binding-id
//
// ```
//
//	**Note**`cloud-binding-id` has the following structure `{billing_account_id}/cloud/{cloud_id}`,
//
// where `{billing_account_id}` refers to the billing account id (`foo-ba-id` in example above)
//
// and `{cloud_id}` refers to the cloud id (`foo-cloud-id` in example above). This way `cloud-binding-id` must be equals to `foo-ba-id/cloud/foo-cloud-id`.
type YandexBillingCloudBinding struct {
	pulumi.CustomResourceState

	// ID of billing account to bind cloud to.
	BillingAccountId pulumi.StringOutput `pulumi:"billingAccountId"`
	// ID of cloud to bind.
	CloudId pulumi.StringOutput `pulumi:"cloudId"`
}

// NewYandexBillingCloudBinding registers a new resource with the given unique name, arguments, and options.
func NewYandexBillingCloudBinding(ctx *pulumi.Context,
	name string, args *YandexBillingCloudBindingArgs, opts ...pulumi.ResourceOption) (*YandexBillingCloudBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.CloudId == nil {
		return nil, errors.New("invalid value for required argument 'CloudId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource YandexBillingCloudBinding
	err := ctx.RegisterResource("yandex:index/yandexBillingCloudBinding:yandexBillingCloudBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetYandexBillingCloudBinding gets an existing YandexBillingCloudBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetYandexBillingCloudBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *YandexBillingCloudBindingState, opts ...pulumi.ResourceOption) (*YandexBillingCloudBinding, error) {
	var resource YandexBillingCloudBinding
	err := ctx.ReadResource("yandex:index/yandexBillingCloudBinding:yandexBillingCloudBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering YandexBillingCloudBinding resources.
type yandexBillingCloudBindingState struct {
	// ID of billing account to bind cloud to.
	BillingAccountId *string `pulumi:"billingAccountId"`
	// ID of cloud to bind.
	CloudId *string `pulumi:"cloudId"`
}

type YandexBillingCloudBindingState struct {
	// ID of billing account to bind cloud to.
	BillingAccountId pulumi.StringPtrInput
	// ID of cloud to bind.
	CloudId pulumi.StringPtrInput
}

func (YandexBillingCloudBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*yandexBillingCloudBindingState)(nil)).Elem()
}

type yandexBillingCloudBindingArgs struct {
	// ID of billing account to bind cloud to.
	BillingAccountId string `pulumi:"billingAccountId"`
	// ID of cloud to bind.
	CloudId string `pulumi:"cloudId"`
}

// The set of arguments for constructing a YandexBillingCloudBinding resource.
type YandexBillingCloudBindingArgs struct {
	// ID of billing account to bind cloud to.
	BillingAccountId pulumi.StringInput
	// ID of cloud to bind.
	CloudId pulumi.StringInput
}

func (YandexBillingCloudBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*yandexBillingCloudBindingArgs)(nil)).Elem()
}

type YandexBillingCloudBindingInput interface {
	pulumi.Input

	ToYandexBillingCloudBindingOutput() YandexBillingCloudBindingOutput
	ToYandexBillingCloudBindingOutputWithContext(ctx context.Context) YandexBillingCloudBindingOutput
}

func (*YandexBillingCloudBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**YandexBillingCloudBinding)(nil)).Elem()
}

func (i *YandexBillingCloudBinding) ToYandexBillingCloudBindingOutput() YandexBillingCloudBindingOutput {
	return i.ToYandexBillingCloudBindingOutputWithContext(context.Background())
}

func (i *YandexBillingCloudBinding) ToYandexBillingCloudBindingOutputWithContext(ctx context.Context) YandexBillingCloudBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandexBillingCloudBindingOutput)
}

// YandexBillingCloudBindingArrayInput is an input type that accepts YandexBillingCloudBindingArray and YandexBillingCloudBindingArrayOutput values.
// You can construct a concrete instance of `YandexBillingCloudBindingArrayInput` via:
//
//	YandexBillingCloudBindingArray{ YandexBillingCloudBindingArgs{...} }
type YandexBillingCloudBindingArrayInput interface {
	pulumi.Input

	ToYandexBillingCloudBindingArrayOutput() YandexBillingCloudBindingArrayOutput
	ToYandexBillingCloudBindingArrayOutputWithContext(context.Context) YandexBillingCloudBindingArrayOutput
}

type YandexBillingCloudBindingArray []YandexBillingCloudBindingInput

func (YandexBillingCloudBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*YandexBillingCloudBinding)(nil)).Elem()
}

func (i YandexBillingCloudBindingArray) ToYandexBillingCloudBindingArrayOutput() YandexBillingCloudBindingArrayOutput {
	return i.ToYandexBillingCloudBindingArrayOutputWithContext(context.Background())
}

func (i YandexBillingCloudBindingArray) ToYandexBillingCloudBindingArrayOutputWithContext(ctx context.Context) YandexBillingCloudBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandexBillingCloudBindingArrayOutput)
}

// YandexBillingCloudBindingMapInput is an input type that accepts YandexBillingCloudBindingMap and YandexBillingCloudBindingMapOutput values.
// You can construct a concrete instance of `YandexBillingCloudBindingMapInput` via:
//
//	YandexBillingCloudBindingMap{ "key": YandexBillingCloudBindingArgs{...} }
type YandexBillingCloudBindingMapInput interface {
	pulumi.Input

	ToYandexBillingCloudBindingMapOutput() YandexBillingCloudBindingMapOutput
	ToYandexBillingCloudBindingMapOutputWithContext(context.Context) YandexBillingCloudBindingMapOutput
}

type YandexBillingCloudBindingMap map[string]YandexBillingCloudBindingInput

func (YandexBillingCloudBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*YandexBillingCloudBinding)(nil)).Elem()
}

func (i YandexBillingCloudBindingMap) ToYandexBillingCloudBindingMapOutput() YandexBillingCloudBindingMapOutput {
	return i.ToYandexBillingCloudBindingMapOutputWithContext(context.Background())
}

func (i YandexBillingCloudBindingMap) ToYandexBillingCloudBindingMapOutputWithContext(ctx context.Context) YandexBillingCloudBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandexBillingCloudBindingMapOutput)
}

type YandexBillingCloudBindingOutput struct{ *pulumi.OutputState }

func (YandexBillingCloudBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YandexBillingCloudBinding)(nil)).Elem()
}

func (o YandexBillingCloudBindingOutput) ToYandexBillingCloudBindingOutput() YandexBillingCloudBindingOutput {
	return o
}

func (o YandexBillingCloudBindingOutput) ToYandexBillingCloudBindingOutputWithContext(ctx context.Context) YandexBillingCloudBindingOutput {
	return o
}

// ID of billing account to bind cloud to.
func (o YandexBillingCloudBindingOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexBillingCloudBinding) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

// ID of cloud to bind.
func (o YandexBillingCloudBindingOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *YandexBillingCloudBinding) pulumi.StringOutput { return v.CloudId }).(pulumi.StringOutput)
}

type YandexBillingCloudBindingArrayOutput struct{ *pulumi.OutputState }

func (YandexBillingCloudBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*YandexBillingCloudBinding)(nil)).Elem()
}

func (o YandexBillingCloudBindingArrayOutput) ToYandexBillingCloudBindingArrayOutput() YandexBillingCloudBindingArrayOutput {
	return o
}

func (o YandexBillingCloudBindingArrayOutput) ToYandexBillingCloudBindingArrayOutputWithContext(ctx context.Context) YandexBillingCloudBindingArrayOutput {
	return o
}

func (o YandexBillingCloudBindingArrayOutput) Index(i pulumi.IntInput) YandexBillingCloudBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *YandexBillingCloudBinding {
		return vs[0].([]*YandexBillingCloudBinding)[vs[1].(int)]
	}).(YandexBillingCloudBindingOutput)
}

type YandexBillingCloudBindingMapOutput struct{ *pulumi.OutputState }

func (YandexBillingCloudBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*YandexBillingCloudBinding)(nil)).Elem()
}

func (o YandexBillingCloudBindingMapOutput) ToYandexBillingCloudBindingMapOutput() YandexBillingCloudBindingMapOutput {
	return o
}

func (o YandexBillingCloudBindingMapOutput) ToYandexBillingCloudBindingMapOutputWithContext(ctx context.Context) YandexBillingCloudBindingMapOutput {
	return o
}

func (o YandexBillingCloudBindingMapOutput) MapIndex(k pulumi.StringInput) YandexBillingCloudBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *YandexBillingCloudBinding {
		return vs[0].(map[string]*YandexBillingCloudBinding)[vs[1].(string)]
	}).(YandexBillingCloudBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*YandexBillingCloudBindingInput)(nil)).Elem(), &YandexBillingCloudBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*YandexBillingCloudBindingArrayInput)(nil)).Elem(), YandexBillingCloudBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*YandexBillingCloudBindingMapInput)(nil)).Elem(), YandexBillingCloudBindingMap{})
	pulumi.RegisterOutputType(YandexBillingCloudBindingOutput{})
	pulumi.RegisterOutputType(YandexBillingCloudBindingArrayOutput{})
	pulumi.RegisterOutputType(YandexBillingCloudBindingMapOutput{})
}