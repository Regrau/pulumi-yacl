// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster")]
    public partial class MdbGreenplumCluster : global::Pulumi.CustomResource
    {
        [Output("access")]
        public Output<Outputs.MdbGreenplumClusterAccess> Access { get; private set; } = null!;

        [Output("assignPublicIp")]
        public Output<bool> AssignPublicIp { get; private set; } = null!;

        [Output("backupWindowStart")]
        public Output<Outputs.MdbGreenplumClusterBackupWindowStart> BackupWindowStart { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        [Output("greenplumConfig")]
        public Output<ImmutableDictionary<string, string>> GreenplumConfig { get; private set; } = null!;

        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("maintenanceWindow")]
        public Output<Outputs.MdbGreenplumClusterMaintenanceWindow> MaintenanceWindow { get; private set; } = null!;

        [Output("masterHostCount")]
        public Output<int> MasterHostCount { get; private set; } = null!;

        [Output("masterHosts")]
        public Output<ImmutableArray<Outputs.MdbGreenplumClusterMasterHost>> MasterHosts { get; private set; } = null!;

        [Output("masterSubcluster")]
        public Output<Outputs.MdbGreenplumClusterMasterSubcluster> MasterSubcluster { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        [Output("poolerConfig")]
        public Output<Outputs.MdbGreenplumClusterPoolerConfig?> PoolerConfig { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("segmentHostCount")]
        public Output<int> SegmentHostCount { get; private set; } = null!;

        [Output("segmentHosts")]
        public Output<ImmutableArray<Outputs.MdbGreenplumClusterSegmentHost>> SegmentHosts { get; private set; } = null!;

        [Output("segmentInHost")]
        public Output<int> SegmentInHost { get; private set; } = null!;

        [Output("segmentSubcluster")]
        public Output<Outputs.MdbGreenplumClusterSegmentSubcluster> SegmentSubcluster { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        [Output("userPassword")]
        public Output<string> UserPassword { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a MdbGreenplumCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbGreenplumCluster(string name, MdbGreenplumClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, args ?? new MdbGreenplumClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbGreenplumCluster(string name, Input<string> id, MdbGreenplumClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MdbGreenplumCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbGreenplumCluster Get(string name, Input<string> id, MdbGreenplumClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbGreenplumCluster(name, id, state, options);
        }
    }

    public sealed class MdbGreenplumClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("access")]
        public Input<Inputs.MdbGreenplumClusterAccessArgs>? Access { get; set; }

        [Input("assignPublicIp", required: true)]
        public Input<bool> AssignPublicIp { get; set; } = null!;

        [Input("backupWindowStart")]
        public Input<Inputs.MdbGreenplumClusterBackupWindowStartArgs>? BackupWindowStart { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("greenplumConfig")]
        private InputMap<string>? _greenplumConfig;
        public InputMap<string> GreenplumConfig
        {
            get => _greenplumConfig ?? (_greenplumConfig = new InputMap<string>());
            set => _greenplumConfig = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("maintenanceWindow")]
        public Input<Inputs.MdbGreenplumClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        [Input("masterHostCount", required: true)]
        public Input<int> MasterHostCount { get; set; } = null!;

        [Input("masterSubcluster", required: true)]
        public Input<Inputs.MdbGreenplumClusterMasterSubclusterArgs> MasterSubcluster { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        [Input("poolerConfig")]
        public Input<Inputs.MdbGreenplumClusterPoolerConfigArgs>? PoolerConfig { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("segmentHostCount", required: true)]
        public Input<int> SegmentHostCount { get; set; } = null!;

        [Input("segmentInHost", required: true)]
        public Input<int> SegmentInHost { get; set; } = null!;

        [Input("segmentSubcluster", required: true)]
        public Input<Inputs.MdbGreenplumClusterSegmentSubclusterArgs> SegmentSubcluster { get; set; } = null!;

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        [Input("userPassword", required: true)]
        public Input<string> UserPassword { get; set; } = null!;

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public MdbGreenplumClusterArgs()
        {
        }
        public static new MdbGreenplumClusterArgs Empty => new MdbGreenplumClusterArgs();
    }

    public sealed class MdbGreenplumClusterState : global::Pulumi.ResourceArgs
    {
        [Input("access")]
        public Input<Inputs.MdbGreenplumClusterAccessGetArgs>? Access { get; set; }

        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        [Input("backupWindowStart")]
        public Input<Inputs.MdbGreenplumClusterBackupWindowStartGetArgs>? BackupWindowStart { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environment")]
        public Input<string>? Environment { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("greenplumConfig")]
        private InputMap<string>? _greenplumConfig;
        public InputMap<string> GreenplumConfig
        {
            get => _greenplumConfig ?? (_greenplumConfig = new InputMap<string>());
            set => _greenplumConfig = value;
        }

        [Input("health")]
        public Input<string>? Health { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("maintenanceWindow")]
        public Input<Inputs.MdbGreenplumClusterMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        [Input("masterHostCount")]
        public Input<int>? MasterHostCount { get; set; }

        [Input("masterHosts")]
        private InputList<Inputs.MdbGreenplumClusterMasterHostGetArgs>? _masterHosts;
        public InputList<Inputs.MdbGreenplumClusterMasterHostGetArgs> MasterHosts
        {
            get => _masterHosts ?? (_masterHosts = new InputList<Inputs.MdbGreenplumClusterMasterHostGetArgs>());
            set => _masterHosts = value;
        }

        [Input("masterSubcluster")]
        public Input<Inputs.MdbGreenplumClusterMasterSubclusterGetArgs>? MasterSubcluster { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("poolerConfig")]
        public Input<Inputs.MdbGreenplumClusterPoolerConfigGetArgs>? PoolerConfig { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("segmentHostCount")]
        public Input<int>? SegmentHostCount { get; set; }

        [Input("segmentHosts")]
        private InputList<Inputs.MdbGreenplumClusterSegmentHostGetArgs>? _segmentHosts;
        public InputList<Inputs.MdbGreenplumClusterSegmentHostGetArgs> SegmentHosts
        {
            get => _segmentHosts ?? (_segmentHosts = new InputList<Inputs.MdbGreenplumClusterSegmentHostGetArgs>());
            set => _segmentHosts = value;
        }

        [Input("segmentInHost")]
        public Input<int>? SegmentInHost { get; set; }

        [Input("segmentSubcluster")]
        public Input<Inputs.MdbGreenplumClusterSegmentSubclusterGetArgs>? SegmentSubcluster { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        [Input("userPassword")]
        public Input<string>? UserPassword { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public MdbGreenplumClusterState()
        {
        }
        public static new MdbGreenplumClusterState Empty => new MdbGreenplumClusterState();
    }
}