// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := yandex.NewKmsAsymetricEncryptionKey(ctx, "key-a", &yandex.KmsAsymetricEncryptionKeyArgs{
//				Description:         pulumi.String("description for key"),
//				EncryptionAlgorithm: pulumi.String("RSA_2048_ENC_OAEP_SHA_256"),
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
// A KMS asymmetric encryption key can be imported using the `id` of the resource, e.g.
//
// ```sh
//
//	$ pulumi import yandex:index/kmsAsymetricEncryptionKey:KmsAsymetricEncryptionKey top-secret kms_asymmetric_encryption_key_id
//
// ```
type KmsAsymetricEncryptionKey struct {
	pulumi.CustomResourceState

	// Creation timestamp of the key.
	CreatedAt          pulumi.StringOutput  `pulumi:"createdAt"`
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// An optional description of the key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Encryption algorithm to be used with a new key. The default value is `RSA_2048_ENC_OAEP_SHA_256`.
	EncryptionAlgorithm pulumi.StringPtrOutput `pulumi:"encryptionAlgorithm"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the key.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the key.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewKmsAsymetricEncryptionKey registers a new resource with the given unique name, arguments, and options.
func NewKmsAsymetricEncryptionKey(ctx *pulumi.Context,
	name string, args *KmsAsymetricEncryptionKeyArgs, opts ...pulumi.ResourceOption) (*KmsAsymetricEncryptionKey, error) {
	if args == nil {
		args = &KmsAsymetricEncryptionKeyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource KmsAsymetricEncryptionKey
	err := ctx.RegisterResource("yandex:index/kmsAsymetricEncryptionKey:KmsAsymetricEncryptionKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsAsymetricEncryptionKey gets an existing KmsAsymetricEncryptionKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsAsymetricEncryptionKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsAsymetricEncryptionKeyState, opts ...pulumi.ResourceOption) (*KmsAsymetricEncryptionKey, error) {
	var resource KmsAsymetricEncryptionKey
	err := ctx.ReadResource("yandex:index/kmsAsymetricEncryptionKey:KmsAsymetricEncryptionKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsAsymetricEncryptionKey resources.
type kmsAsymetricEncryptionKeyState struct {
	// Creation timestamp of the key.
	CreatedAt          *string `pulumi:"createdAt"`
	DeletionProtection *bool   `pulumi:"deletionProtection"`
	// An optional description of the key.
	Description *string `pulumi:"description"`
	// Encryption algorithm to be used with a new key. The default value is `RSA_2048_ENC_OAEP_SHA_256`.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the key.
	Labels map[string]string `pulumi:"labels"`
	// Name of the key.
	Name *string `pulumi:"name"`
	// The status of the key.
	Status *string `pulumi:"status"`
}

type KmsAsymetricEncryptionKeyState struct {
	// Creation timestamp of the key.
	CreatedAt          pulumi.StringPtrInput
	DeletionProtection pulumi.BoolPtrInput
	// An optional description of the key.
	Description pulumi.StringPtrInput
	// Encryption algorithm to be used with a new key. The default value is `RSA_2048_ENC_OAEP_SHA_256`.
	EncryptionAlgorithm pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the key.
	Labels pulumi.StringMapInput
	// Name of the key.
	Name pulumi.StringPtrInput
	// The status of the key.
	Status pulumi.StringPtrInput
}

func (KmsAsymetricEncryptionKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAsymetricEncryptionKeyState)(nil)).Elem()
}

type kmsAsymetricEncryptionKeyArgs struct {
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// An optional description of the key.
	Description *string `pulumi:"description"`
	// Encryption algorithm to be used with a new key. The default value is `RSA_2048_ENC_OAEP_SHA_256`.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the key.
	Labels map[string]string `pulumi:"labels"`
	// Name of the key.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a KmsAsymetricEncryptionKey resource.
type KmsAsymetricEncryptionKeyArgs struct {
	DeletionProtection pulumi.BoolPtrInput
	// An optional description of the key.
	Description pulumi.StringPtrInput
	// Encryption algorithm to be used with a new key. The default value is `RSA_2048_ENC_OAEP_SHA_256`.
	EncryptionAlgorithm pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the key.
	Labels pulumi.StringMapInput
	// Name of the key.
	Name pulumi.StringPtrInput
}

func (KmsAsymetricEncryptionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAsymetricEncryptionKeyArgs)(nil)).Elem()
}

type KmsAsymetricEncryptionKeyInput interface {
	pulumi.Input

	ToKmsAsymetricEncryptionKeyOutput() KmsAsymetricEncryptionKeyOutput
	ToKmsAsymetricEncryptionKeyOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyOutput
}

func (*KmsAsymetricEncryptionKey) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAsymetricEncryptionKey)(nil)).Elem()
}

func (i *KmsAsymetricEncryptionKey) ToKmsAsymetricEncryptionKeyOutput() KmsAsymetricEncryptionKeyOutput {
	return i.ToKmsAsymetricEncryptionKeyOutputWithContext(context.Background())
}

func (i *KmsAsymetricEncryptionKey) ToKmsAsymetricEncryptionKeyOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymetricEncryptionKeyOutput)
}

// KmsAsymetricEncryptionKeyArrayInput is an input type that accepts KmsAsymetricEncryptionKeyArray and KmsAsymetricEncryptionKeyArrayOutput values.
// You can construct a concrete instance of `KmsAsymetricEncryptionKeyArrayInput` via:
//
//	KmsAsymetricEncryptionKeyArray{ KmsAsymetricEncryptionKeyArgs{...} }
type KmsAsymetricEncryptionKeyArrayInput interface {
	pulumi.Input

	ToKmsAsymetricEncryptionKeyArrayOutput() KmsAsymetricEncryptionKeyArrayOutput
	ToKmsAsymetricEncryptionKeyArrayOutputWithContext(context.Context) KmsAsymetricEncryptionKeyArrayOutput
}

type KmsAsymetricEncryptionKeyArray []KmsAsymetricEncryptionKeyInput

func (KmsAsymetricEncryptionKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsAsymetricEncryptionKey)(nil)).Elem()
}

func (i KmsAsymetricEncryptionKeyArray) ToKmsAsymetricEncryptionKeyArrayOutput() KmsAsymetricEncryptionKeyArrayOutput {
	return i.ToKmsAsymetricEncryptionKeyArrayOutputWithContext(context.Background())
}

func (i KmsAsymetricEncryptionKeyArray) ToKmsAsymetricEncryptionKeyArrayOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymetricEncryptionKeyArrayOutput)
}

// KmsAsymetricEncryptionKeyMapInput is an input type that accepts KmsAsymetricEncryptionKeyMap and KmsAsymetricEncryptionKeyMapOutput values.
// You can construct a concrete instance of `KmsAsymetricEncryptionKeyMapInput` via:
//
//	KmsAsymetricEncryptionKeyMap{ "key": KmsAsymetricEncryptionKeyArgs{...} }
type KmsAsymetricEncryptionKeyMapInput interface {
	pulumi.Input

	ToKmsAsymetricEncryptionKeyMapOutput() KmsAsymetricEncryptionKeyMapOutput
	ToKmsAsymetricEncryptionKeyMapOutputWithContext(context.Context) KmsAsymetricEncryptionKeyMapOutput
}

type KmsAsymetricEncryptionKeyMap map[string]KmsAsymetricEncryptionKeyInput

func (KmsAsymetricEncryptionKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsAsymetricEncryptionKey)(nil)).Elem()
}

func (i KmsAsymetricEncryptionKeyMap) ToKmsAsymetricEncryptionKeyMapOutput() KmsAsymetricEncryptionKeyMapOutput {
	return i.ToKmsAsymetricEncryptionKeyMapOutputWithContext(context.Background())
}

func (i KmsAsymetricEncryptionKeyMap) ToKmsAsymetricEncryptionKeyMapOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymetricEncryptionKeyMapOutput)
}

type KmsAsymetricEncryptionKeyOutput struct{ *pulumi.OutputState }

func (KmsAsymetricEncryptionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAsymetricEncryptionKey)(nil)).Elem()
}

func (o KmsAsymetricEncryptionKeyOutput) ToKmsAsymetricEncryptionKeyOutput() KmsAsymetricEncryptionKeyOutput {
	return o
}

func (o KmsAsymetricEncryptionKeyOutput) ToKmsAsymetricEncryptionKeyOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyOutput {
	return o
}

// Creation timestamp of the key.
func (o KmsAsymetricEncryptionKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o KmsAsymetricEncryptionKeyOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// An optional description of the key.
func (o KmsAsymetricEncryptionKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Encryption algorithm to be used with a new key. The default value is `RSA_2048_ENC_OAEP_SHA_256`.
func (o KmsAsymetricEncryptionKeyOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringPtrOutput { return v.EncryptionAlgorithm }).(pulumi.StringPtrOutput)
}

// The ID of the folder that the resource belongs to. If it
// is not provided, the default provider folder is used.
func (o KmsAsymetricEncryptionKeyOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the key.
func (o KmsAsymetricEncryptionKeyOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the key.
func (o KmsAsymetricEncryptionKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the key.
func (o KmsAsymetricEncryptionKeyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymetricEncryptionKey) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type KmsAsymetricEncryptionKeyArrayOutput struct{ *pulumi.OutputState }

func (KmsAsymetricEncryptionKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsAsymetricEncryptionKey)(nil)).Elem()
}

func (o KmsAsymetricEncryptionKeyArrayOutput) ToKmsAsymetricEncryptionKeyArrayOutput() KmsAsymetricEncryptionKeyArrayOutput {
	return o
}

func (o KmsAsymetricEncryptionKeyArrayOutput) ToKmsAsymetricEncryptionKeyArrayOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyArrayOutput {
	return o
}

func (o KmsAsymetricEncryptionKeyArrayOutput) Index(i pulumi.IntInput) KmsAsymetricEncryptionKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KmsAsymetricEncryptionKey {
		return vs[0].([]*KmsAsymetricEncryptionKey)[vs[1].(int)]
	}).(KmsAsymetricEncryptionKeyOutput)
}

type KmsAsymetricEncryptionKeyMapOutput struct{ *pulumi.OutputState }

func (KmsAsymetricEncryptionKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsAsymetricEncryptionKey)(nil)).Elem()
}

func (o KmsAsymetricEncryptionKeyMapOutput) ToKmsAsymetricEncryptionKeyMapOutput() KmsAsymetricEncryptionKeyMapOutput {
	return o
}

func (o KmsAsymetricEncryptionKeyMapOutput) ToKmsAsymetricEncryptionKeyMapOutputWithContext(ctx context.Context) KmsAsymetricEncryptionKeyMapOutput {
	return o
}

func (o KmsAsymetricEncryptionKeyMapOutput) MapIndex(k pulumi.StringInput) KmsAsymetricEncryptionKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KmsAsymetricEncryptionKey {
		return vs[0].(map[string]*KmsAsymetricEncryptionKey)[vs[1].(string)]
	}).(KmsAsymetricEncryptionKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymetricEncryptionKeyInput)(nil)).Elem(), &KmsAsymetricEncryptionKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymetricEncryptionKeyArrayInput)(nil)).Elem(), KmsAsymetricEncryptionKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymetricEncryptionKeyMapInput)(nil)).Elem(), KmsAsymetricEncryptionKeyMap{})
	pulumi.RegisterOutputType(KmsAsymetricEncryptionKeyOutput{})
	pulumi.RegisterOutputType(KmsAsymetricEncryptionKeyArrayOutput{})
	pulumi.RegisterOutputType(KmsAsymetricEncryptionKeyMapOutput{})
}
