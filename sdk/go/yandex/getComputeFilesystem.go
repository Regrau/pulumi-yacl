// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetComputeFilesystem(ctx *pulumi.Context, args *GetComputeFilesystemArgs, opts ...pulumi.InvokeOption) (*GetComputeFilesystemResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetComputeFilesystemResult
	err := ctx.Invoke("yandex:index/getComputeFilesystem:getComputeFilesystem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeFilesystem.
type GetComputeFilesystemArgs struct {
	FilesystemId *string `pulumi:"filesystemId"`
	FolderId     *string `pulumi:"folderId"`
	Name         *string `pulumi:"name"`
}

// A collection of values returned by getComputeFilesystem.
type GetComputeFilesystemResult struct {
	BlockSize    int    `pulumi:"blockSize"`
	CreatedAt    string `pulumi:"createdAt"`
	Description  string `pulumi:"description"`
	FilesystemId string `pulumi:"filesystemId"`
	FolderId     string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string            `pulumi:"id"`
	Labels map[string]string `pulumi:"labels"`
	Name   string            `pulumi:"name"`
	Size   int               `pulumi:"size"`
	Status string            `pulumi:"status"`
	Type   string            `pulumi:"type"`
	Zone   string            `pulumi:"zone"`
}

func GetComputeFilesystemOutput(ctx *pulumi.Context, args GetComputeFilesystemOutputArgs, opts ...pulumi.InvokeOption) GetComputeFilesystemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetComputeFilesystemResult, error) {
			args := v.(GetComputeFilesystemArgs)
			r, err := GetComputeFilesystem(ctx, &args, opts...)
			var s GetComputeFilesystemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetComputeFilesystemResultOutput)
}

// A collection of arguments for invoking getComputeFilesystem.
type GetComputeFilesystemOutputArgs struct {
	FilesystemId pulumi.StringPtrInput `pulumi:"filesystemId"`
	FolderId     pulumi.StringPtrInput `pulumi:"folderId"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (GetComputeFilesystemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeFilesystemArgs)(nil)).Elem()
}

// A collection of values returned by getComputeFilesystem.
type GetComputeFilesystemResultOutput struct{ *pulumi.OutputState }

func (GetComputeFilesystemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeFilesystemResult)(nil)).Elem()
}

func (o GetComputeFilesystemResultOutput) ToGetComputeFilesystemResultOutput() GetComputeFilesystemResultOutput {
	return o
}

func (o GetComputeFilesystemResultOutput) ToGetComputeFilesystemResultOutputWithContext(ctx context.Context) GetComputeFilesystemResultOutput {
	return o
}

func (o GetComputeFilesystemResultOutput) BlockSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) int { return v.BlockSize }).(pulumi.IntOutput)
}

func (o GetComputeFilesystemResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) FilesystemId() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.FilesystemId }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetComputeFilesystemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o GetComputeFilesystemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) int { return v.Size }).(pulumi.IntOutput)
}

func (o GetComputeFilesystemResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetComputeFilesystemResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeFilesystemResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetComputeFilesystemResultOutput{})
}
