// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMdbClickhouseCluster(ctx *pulumi.Context, args *LookupMdbClickhouseClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbClickhouseClusterResult, error) {
	var rv LookupMdbClickhouseClusterResult
	err := ctx.Invoke("yandex:index/getMdbClickhouseCluster:getMdbClickhouseCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbClickhouseCluster.
type LookupMdbClickhouseClusterArgs struct {
	CloudStorage       *GetMdbClickhouseClusterCloudStorage `pulumi:"cloudStorage"`
	ClusterId          *string                              `pulumi:"clusterId"`
	DeletionProtection *bool                                `pulumi:"deletionProtection"`
	FolderId           *string                              `pulumi:"folderId"`
	Name               *string                              `pulumi:"name"`
	ServiceAccountId   *string                              `pulumi:"serviceAccountId"`
}

// A collection of values returned by getMdbClickhouseCluster.
type LookupMdbClickhouseClusterResult struct {
	Accesses           []GetMdbClickhouseClusterAccess            `pulumi:"accesses"`
	BackupWindowStarts []GetMdbClickhouseClusterBackupWindowStart `pulumi:"backupWindowStarts"`
	Clickhouses        []GetMdbClickhouseClusterClickhouse        `pulumi:"clickhouses"`
	CloudStorage       *GetMdbClickhouseClusterCloudStorage       `pulumi:"cloudStorage"`
	ClusterId          string                                     `pulumi:"clusterId"`
	CreatedAt          string                                     `pulumi:"createdAt"`
	Databases          []GetMdbClickhouseClusterDatabase          `pulumi:"databases"`
	DeletionProtection bool                                       `pulumi:"deletionProtection"`
	Description        string                                     `pulumi:"description"`
	EmbeddedKeeper     bool                                       `pulumi:"embeddedKeeper"`
	Environment        string                                     `pulumi:"environment"`
	FolderId           string                                     `pulumi:"folderId"`
	FormatSchemas      []GetMdbClickhouseClusterFormatSchema      `pulumi:"formatSchemas"`
	Health             string                                     `pulumi:"health"`
	Hosts              []GetMdbClickhouseClusterHost              `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string                                     `pulumi:"id"`
	Labels                map[string]string                          `pulumi:"labels"`
	MaintenanceWindows    []GetMdbClickhouseClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	MlModels              []GetMdbClickhouseClusterMlModel           `pulumi:"mlModels"`
	Name                  string                                     `pulumi:"name"`
	NetworkId             string                                     `pulumi:"networkId"`
	SecurityGroupIds      []string                                   `pulumi:"securityGroupIds"`
	ServiceAccountId      string                                     `pulumi:"serviceAccountId"`
	ShardGroups           []GetMdbClickhouseClusterShardGroup        `pulumi:"shardGroups"`
	SqlDatabaseManagement bool                                       `pulumi:"sqlDatabaseManagement"`
	SqlUserManagement     bool                                       `pulumi:"sqlUserManagement"`
	Status                string                                     `pulumi:"status"`
	Users                 []GetMdbClickhouseClusterUser              `pulumi:"users"`
	Version               string                                     `pulumi:"version"`
	Zookeepers            []GetMdbClickhouseClusterZookeeper         `pulumi:"zookeepers"`
}

func LookupMdbClickhouseClusterOutput(ctx *pulumi.Context, args LookupMdbClickhouseClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbClickhouseClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMdbClickhouseClusterResult, error) {
			args := v.(LookupMdbClickhouseClusterArgs)
			r, err := LookupMdbClickhouseCluster(ctx, &args, opts...)
			var s LookupMdbClickhouseClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMdbClickhouseClusterResultOutput)
}

// A collection of arguments for invoking getMdbClickhouseCluster.
type LookupMdbClickhouseClusterOutputArgs struct {
	CloudStorage       GetMdbClickhouseClusterCloudStoragePtrInput `pulumi:"cloudStorage"`
	ClusterId          pulumi.StringPtrInput                       `pulumi:"clusterId"`
	DeletionProtection pulumi.BoolPtrInput                         `pulumi:"deletionProtection"`
	FolderId           pulumi.StringPtrInput                       `pulumi:"folderId"`
	Name               pulumi.StringPtrInput                       `pulumi:"name"`
	ServiceAccountId   pulumi.StringPtrInput                       `pulumi:"serviceAccountId"`
}

func (LookupMdbClickhouseClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbClickhouseClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbClickhouseCluster.
type LookupMdbClickhouseClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbClickhouseClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbClickhouseClusterResult)(nil)).Elem()
}

func (o LookupMdbClickhouseClusterResultOutput) ToLookupMdbClickhouseClusterResultOutput() LookupMdbClickhouseClusterResultOutput {
	return o
}

func (o LookupMdbClickhouseClusterResultOutput) ToLookupMdbClickhouseClusterResultOutputWithContext(ctx context.Context) LookupMdbClickhouseClusterResultOutput {
	return o
}

func (o LookupMdbClickhouseClusterResultOutput) Accesses() GetMdbClickhouseClusterAccessArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterAccess { return v.Accesses }).(GetMdbClickhouseClusterAccessArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) BackupWindowStarts() GetMdbClickhouseClusterBackupWindowStartArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterBackupWindowStart {
		return v.BackupWindowStarts
	}).(GetMdbClickhouseClusterBackupWindowStartArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Clickhouses() GetMdbClickhouseClusterClickhouseArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterClickhouse { return v.Clickhouses }).(GetMdbClickhouseClusterClickhouseArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) CloudStorage() GetMdbClickhouseClusterCloudStoragePtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *GetMdbClickhouseClusterCloudStorage { return v.CloudStorage }).(GetMdbClickhouseClusterCloudStoragePtrOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Databases() GetMdbClickhouseClusterDatabaseArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterDatabase { return v.Databases }).(GetMdbClickhouseClusterDatabaseArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) EmbeddedKeeper() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.EmbeddedKeeper }).(pulumi.BoolOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) FormatSchemas() GetMdbClickhouseClusterFormatSchemaArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterFormatSchema { return v.FormatSchemas }).(GetMdbClickhouseClusterFormatSchemaArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Hosts() GetMdbClickhouseClusterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterHost { return v.Hosts }).(GetMdbClickhouseClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbClickhouseClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) MaintenanceWindows() GetMdbClickhouseClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterMaintenanceWindow {
		return v.MaintenanceWindows
	}).(GetMdbClickhouseClusterMaintenanceWindowArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) MlModels() GetMdbClickhouseClusterMlModelArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterMlModel { return v.MlModels }).(GetMdbClickhouseClusterMlModelArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ShardGroups() GetMdbClickhouseClusterShardGroupArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterShardGroup { return v.ShardGroups }).(GetMdbClickhouseClusterShardGroupArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) SqlDatabaseManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.SqlDatabaseManagement }).(pulumi.BoolOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) SqlUserManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.SqlUserManagement }).(pulumi.BoolOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Users() GetMdbClickhouseClusterUserArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterUser { return v.Users }).(GetMdbClickhouseClusterUserArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Zookeepers() GetMdbClickhouseClusterZookeeperArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterZookeeper { return v.Zookeepers }).(GetMdbClickhouseClusterZookeeperArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbClickhouseClusterResultOutput{})
}