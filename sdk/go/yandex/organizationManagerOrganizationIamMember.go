// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of a single member for a single binding within
// the IAM policy for an existing Yandex Organization Manager organization.
//
// > **Note:** Roles controlled by `OrganizationManagerOrganizationIamBinding`
//
//	should not be assigned using `OrganizationManagerOrganizationIamMember`.
//
// > **Note:** When you delete `OrganizationManagerOrganizationIamBinding` resource,
//
//	the roles can be deleted from other users within the organization as well. Be careful!
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
//			_, err := yandex.NewOrganizationManagerOrganizationIamMember(ctx, "editor", &yandex.OrganizationManagerOrganizationIamMemberArgs{
//				Member:         pulumi.String("userAccount:user_id"),
//				OrganizationId: pulumi.String("some_organization_id"),
//				Role:           pulumi.String("editor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the `organization id`, role, and account, e.g.
//
// ```sh
//
//	$ pulumi import yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember my_project "organization_id viewer foo@example.com"
//
// ```
type OrganizationManagerOrganizationIamMember struct {
	pulumi.CustomResourceState

	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
	// * **group:{group_id}**: A unique group ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Member pulumi.StringOutput `pulumi:"member"`
	// ID of the organization to attach a policy to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The role that should be assigned.
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewOrganizationManagerOrganizationIamMember registers a new resource with the given unique name, arguments, and options.
func NewOrganizationManagerOrganizationIamMember(ctx *pulumi.Context,
	name string, args *OrganizationManagerOrganizationIamMemberArgs, opts ...pulumi.ResourceOption) (*OrganizationManagerOrganizationIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource OrganizationManagerOrganizationIamMember
	err := ctx.RegisterResource("yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationManagerOrganizationIamMember gets an existing OrganizationManagerOrganizationIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationManagerOrganizationIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationManagerOrganizationIamMemberState, opts ...pulumi.ResourceOption) (*OrganizationManagerOrganizationIamMember, error) {
	var resource OrganizationManagerOrganizationIamMember
	err := ctx.ReadResource("yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationManagerOrganizationIamMember resources.
type organizationManagerOrganizationIamMemberState struct {
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
	// * **group:{group_id}**: A unique group ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Member *string `pulumi:"member"`
	// ID of the organization to attach a policy to.
	OrganizationId *string `pulumi:"organizationId"`
	// The role that should be assigned.
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type OrganizationManagerOrganizationIamMemberState struct {
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
	// * **group:{group_id}**: A unique group ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Member pulumi.StringPtrInput
	// ID of the organization to attach a policy to.
	OrganizationId pulumi.StringPtrInput
	// The role that should be assigned.
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (OrganizationManagerOrganizationIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationManagerOrganizationIamMemberState)(nil)).Elem()
}

type organizationManagerOrganizationIamMemberArgs struct {
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
	// * **group:{group_id}**: A unique group ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Member string `pulumi:"member"`
	// ID of the organization to attach a policy to.
	OrganizationId string `pulumi:"organizationId"`
	// The role that should be assigned.
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a OrganizationManagerOrganizationIamMember resource.
type OrganizationManagerOrganizationIamMemberArgs struct {
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
	// * **group:{group_id}**: A unique group ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Member pulumi.StringInput
	// ID of the organization to attach a policy to.
	OrganizationId pulumi.StringInput
	// The role that should be assigned.
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (OrganizationManagerOrganizationIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationManagerOrganizationIamMemberArgs)(nil)).Elem()
}

type OrganizationManagerOrganizationIamMemberInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamMemberOutput() OrganizationManagerOrganizationIamMemberOutput
	ToOrganizationManagerOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberOutput
}

func (*OrganizationManagerOrganizationIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationManagerOrganizationIamMember)(nil)).Elem()
}

func (i *OrganizationManagerOrganizationIamMember) ToOrganizationManagerOrganizationIamMemberOutput() OrganizationManagerOrganizationIamMemberOutput {
	return i.ToOrganizationManagerOrganizationIamMemberOutputWithContext(context.Background())
}

func (i *OrganizationManagerOrganizationIamMember) ToOrganizationManagerOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamMemberOutput)
}

// OrganizationManagerOrganizationIamMemberArrayInput is an input type that accepts OrganizationManagerOrganizationIamMemberArray and OrganizationManagerOrganizationIamMemberArrayOutput values.
// You can construct a concrete instance of `OrganizationManagerOrganizationIamMemberArrayInput` via:
//
//	OrganizationManagerOrganizationIamMemberArray{ OrganizationManagerOrganizationIamMemberArgs{...} }
type OrganizationManagerOrganizationIamMemberArrayInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamMemberArrayOutput() OrganizationManagerOrganizationIamMemberArrayOutput
	ToOrganizationManagerOrganizationIamMemberArrayOutputWithContext(context.Context) OrganizationManagerOrganizationIamMemberArrayOutput
}

type OrganizationManagerOrganizationIamMemberArray []OrganizationManagerOrganizationIamMemberInput

func (OrganizationManagerOrganizationIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationManagerOrganizationIamMember)(nil)).Elem()
}

func (i OrganizationManagerOrganizationIamMemberArray) ToOrganizationManagerOrganizationIamMemberArrayOutput() OrganizationManagerOrganizationIamMemberArrayOutput {
	return i.ToOrganizationManagerOrganizationIamMemberArrayOutputWithContext(context.Background())
}

func (i OrganizationManagerOrganizationIamMemberArray) ToOrganizationManagerOrganizationIamMemberArrayOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamMemberArrayOutput)
}

// OrganizationManagerOrganizationIamMemberMapInput is an input type that accepts OrganizationManagerOrganizationIamMemberMap and OrganizationManagerOrganizationIamMemberMapOutput values.
// You can construct a concrete instance of `OrganizationManagerOrganizationIamMemberMapInput` via:
//
//	OrganizationManagerOrganizationIamMemberMap{ "key": OrganizationManagerOrganizationIamMemberArgs{...} }
type OrganizationManagerOrganizationIamMemberMapInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamMemberMapOutput() OrganizationManagerOrganizationIamMemberMapOutput
	ToOrganizationManagerOrganizationIamMemberMapOutputWithContext(context.Context) OrganizationManagerOrganizationIamMemberMapOutput
}

type OrganizationManagerOrganizationIamMemberMap map[string]OrganizationManagerOrganizationIamMemberInput

func (OrganizationManagerOrganizationIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationManagerOrganizationIamMember)(nil)).Elem()
}

func (i OrganizationManagerOrganizationIamMemberMap) ToOrganizationManagerOrganizationIamMemberMapOutput() OrganizationManagerOrganizationIamMemberMapOutput {
	return i.ToOrganizationManagerOrganizationIamMemberMapOutputWithContext(context.Background())
}

func (i OrganizationManagerOrganizationIamMemberMap) ToOrganizationManagerOrganizationIamMemberMapOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamMemberMapOutput)
}

type OrganizationManagerOrganizationIamMemberOutput struct{ *pulumi.OutputState }

func (OrganizationManagerOrganizationIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationManagerOrganizationIamMember)(nil)).Elem()
}

func (o OrganizationManagerOrganizationIamMemberOutput) ToOrganizationManagerOrganizationIamMemberOutput() OrganizationManagerOrganizationIamMemberOutput {
	return o
}

func (o OrganizationManagerOrganizationIamMemberOutput) ToOrganizationManagerOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberOutput {
	return o
}

// The identity that will be granted the privilege that is specified in the `role` field.
// This field can have one of the following values:
// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
// * **serviceAccount:{service_account_id}**: A unique service account ID.
// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
// * **group:{group_id}**: A unique group ID.
// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
func (o OrganizationManagerOrganizationIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationManagerOrganizationIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// ID of the organization to attach a policy to.
func (o OrganizationManagerOrganizationIamMemberOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationManagerOrganizationIamMember) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The role that should be assigned.
func (o OrganizationManagerOrganizationIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationManagerOrganizationIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o OrganizationManagerOrganizationIamMemberOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OrganizationManagerOrganizationIamMember) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

type OrganizationManagerOrganizationIamMemberArrayOutput struct{ *pulumi.OutputState }

func (OrganizationManagerOrganizationIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationManagerOrganizationIamMember)(nil)).Elem()
}

func (o OrganizationManagerOrganizationIamMemberArrayOutput) ToOrganizationManagerOrganizationIamMemberArrayOutput() OrganizationManagerOrganizationIamMemberArrayOutput {
	return o
}

func (o OrganizationManagerOrganizationIamMemberArrayOutput) ToOrganizationManagerOrganizationIamMemberArrayOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberArrayOutput {
	return o
}

func (o OrganizationManagerOrganizationIamMemberArrayOutput) Index(i pulumi.IntInput) OrganizationManagerOrganizationIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationManagerOrganizationIamMember {
		return vs[0].([]*OrganizationManagerOrganizationIamMember)[vs[1].(int)]
	}).(OrganizationManagerOrganizationIamMemberOutput)
}

type OrganizationManagerOrganizationIamMemberMapOutput struct{ *pulumi.OutputState }

func (OrganizationManagerOrganizationIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationManagerOrganizationIamMember)(nil)).Elem()
}

func (o OrganizationManagerOrganizationIamMemberMapOutput) ToOrganizationManagerOrganizationIamMemberMapOutput() OrganizationManagerOrganizationIamMemberMapOutput {
	return o
}

func (o OrganizationManagerOrganizationIamMemberMapOutput) ToOrganizationManagerOrganizationIamMemberMapOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamMemberMapOutput {
	return o
}

func (o OrganizationManagerOrganizationIamMemberMapOutput) MapIndex(k pulumi.StringInput) OrganizationManagerOrganizationIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationManagerOrganizationIamMember {
		return vs[0].(map[string]*OrganizationManagerOrganizationIamMember)[vs[1].(string)]
	}).(OrganizationManagerOrganizationIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationManagerOrganizationIamMemberInput)(nil)).Elem(), &OrganizationManagerOrganizationIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationManagerOrganizationIamMemberArrayInput)(nil)).Elem(), OrganizationManagerOrganizationIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationManagerOrganizationIamMemberMapInput)(nil)).Elem(), OrganizationManagerOrganizationIamMemberMap{})
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamMemberOutput{})
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamMemberArrayOutput{})
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamMemberMapOutput{})
}
