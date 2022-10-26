// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMdbElasticSearchCluster(ctx *pulumi.Context, args *LookupMdbElasticSearchClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbElasticSearchClusterResult, error) {
	var rv LookupMdbElasticSearchClusterResult
	err := ctx.Invoke("yandex:index/getMdbElasticSearchCluster:getMdbElasticSearchCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbElasticSearchCluster.
type LookupMdbElasticSearchClusterArgs struct {
	ClusterId          *string           `pulumi:"clusterId"`
	DeletionProtection *bool             `pulumi:"deletionProtection"`
	Description        *string           `pulumi:"description"`
	FolderId           *string           `pulumi:"folderId"`
	Labels             map[string]string `pulumi:"labels"`
	Name               *string           `pulumi:"name"`
	SecurityGroupIds   []string          `pulumi:"securityGroupIds"`
	ServiceAccountId   *string           `pulumi:"serviceAccountId"`
}

// A collection of values returned by getMdbElasticSearchCluster.
type LookupMdbElasticSearchClusterResult struct {
	ClusterId          string                             `pulumi:"clusterId"`
	Configs            []GetMdbElasticSearchClusterConfig `pulumi:"configs"`
	CreatedAt          string                             `pulumi:"createdAt"`
	DeletionProtection bool                               `pulumi:"deletionProtection"`
	Description        string                             `pulumi:"description"`
	Environment        string                             `pulumi:"environment"`
	FolderId           string                             `pulumi:"folderId"`
	Health             string                             `pulumi:"health"`
	Hosts              []GetMdbElasticSearchClusterHost   `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                                        `pulumi:"id"`
	Labels             map[string]string                             `pulumi:"labels"`
	MaintenanceWindows []GetMdbElasticSearchClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	Name               string                                        `pulumi:"name"`
	NetworkId          string                                        `pulumi:"networkId"`
	SecurityGroupIds   []string                                      `pulumi:"securityGroupIds"`
	ServiceAccountId   string                                        `pulumi:"serviceAccountId"`
	Status             string                                        `pulumi:"status"`
}

func LookupMdbElasticSearchClusterOutput(ctx *pulumi.Context, args LookupMdbElasticSearchClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbElasticSearchClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMdbElasticSearchClusterResult, error) {
			args := v.(LookupMdbElasticSearchClusterArgs)
			r, err := LookupMdbElasticSearchCluster(ctx, &args, opts...)
			var s LookupMdbElasticSearchClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMdbElasticSearchClusterResultOutput)
}

// A collection of arguments for invoking getMdbElasticSearchCluster.
type LookupMdbElasticSearchClusterOutputArgs struct {
	ClusterId          pulumi.StringPtrInput   `pulumi:"clusterId"`
	DeletionProtection pulumi.BoolPtrInput     `pulumi:"deletionProtection"`
	Description        pulumi.StringPtrInput   `pulumi:"description"`
	FolderId           pulumi.StringPtrInput   `pulumi:"folderId"`
	Labels             pulumi.StringMapInput   `pulumi:"labels"`
	Name               pulumi.StringPtrInput   `pulumi:"name"`
	SecurityGroupIds   pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	ServiceAccountId   pulumi.StringPtrInput   `pulumi:"serviceAccountId"`
}

func (LookupMdbElasticSearchClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbElasticSearchClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbElasticSearchCluster.
type LookupMdbElasticSearchClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbElasticSearchClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbElasticSearchClusterResult)(nil)).Elem()
}

func (o LookupMdbElasticSearchClusterResultOutput) ToLookupMdbElasticSearchClusterResultOutput() LookupMdbElasticSearchClusterResultOutput {
	return o
}

func (o LookupMdbElasticSearchClusterResultOutput) ToLookupMdbElasticSearchClusterResultOutputWithContext(ctx context.Context) LookupMdbElasticSearchClusterResultOutput {
	return o
}

func (o LookupMdbElasticSearchClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Configs() GetMdbElasticSearchClusterConfigArrayOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) []GetMdbElasticSearchClusterConfig { return v.Configs }).(GetMdbElasticSearchClusterConfigArrayOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Hosts() GetMdbElasticSearchClusterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) []GetMdbElasticSearchClusterHost { return v.Hosts }).(GetMdbElasticSearchClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbElasticSearchClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) MaintenanceWindows() GetMdbElasticSearchClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) []GetMdbElasticSearchClusterMaintenanceWindow {
		return v.MaintenanceWindows
	}).(GetMdbElasticSearchClusterMaintenanceWindowArrayOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

func (o LookupMdbElasticSearchClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbElasticSearchClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbElasticSearchClusterResultOutput{})
}