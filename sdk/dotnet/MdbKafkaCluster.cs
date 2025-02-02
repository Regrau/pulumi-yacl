// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// Manages a Kafka cluster within the Yandex.Cloud. For more information, see
    /// [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).
    /// 
    /// ## Example Usage
    /// 
    /// Example of creating a Single Node Kafka.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fooVpcNetwork = new Yandex.VpcNetwork("fooVpcNetwork");
    /// 
    ///     var fooVpcSubnet = new Yandex.VpcSubnet("fooVpcSubnet", new()
    ///     {
    ///         NetworkId = fooVpcNetwork.Id,
    ///         V4CidrBlocks = new[]
    ///         {
    ///             "10.5.0.0/24",
    ///         },
    ///         Zone = "ru-central1-a",
    ///     });
    /// 
    ///     var fooMdbKafkaCluster = new Yandex.MdbKafkaCluster("fooMdbKafkaCluster", new()
    ///     {
    ///         Config = new Yandex.Inputs.MdbKafkaClusterConfigArgs
    ///         {
    ///             AssignPublicIp = false,
    ///             BrokersCount = 1,
    ///             Kafka = new Yandex.Inputs.MdbKafkaClusterConfigKafkaArgs
    ///             {
    ///                 KafkaConfig = new Yandex.Inputs.MdbKafkaClusterConfigKafkaKafkaConfigArgs
    ///                 {
    ///                     CompressionType = "COMPRESSION_TYPE_ZSTD",
    ///                     DefaultReplicationFactor = "1",
    ///                     LogFlushIntervalMessages = "1024",
    ///                     LogFlushIntervalMs = "1000",
    ///                     LogFlushSchedulerIntervalMs = "1000",
    ///                     LogPreallocate = true,
    ///                     LogRetentionBytes = "1073741824",
    ///                     LogRetentionHours = "168",
    ///                     LogRetentionMinutes = "10080",
    ///                     LogRetentionMs = "86400000",
    ///                     LogSegmentBytes = "134217728",
    ///                     MessageMaxBytes = "1048588",
    ///                     NumPartitions = "10",
    ///                     OffsetsRetentionMinutes = "10080",
    ///                     ReplicaFetchMaxBytes = "1048576",
    ///                     SaslEnabledMechanisms = new[]
    ///                     {
    ///                         "SASL_MECHANISM_SCRAM_SHA_256",
    ///                         "SASL_MECHANISM_SCRAM_SHA_512",
    ///                     },
    ///                     SslCipherSuites = new[]
    ///                     {
    ///                         "TLS_DHE_RSA_WITH_AES_128_CBC_SHA",
    ///                         "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
    ///                     },
    ///                 },
    ///                 Resources = new Yandex.Inputs.MdbKafkaClusterConfigKafkaResourcesArgs
    ///                 {
    ///                     DiskSize = 32,
    ///                     DiskTypeId = "network-ssd",
    ///                     ResourcePresetId = "s2.micro",
    ///                 },
    ///             },
    ///             SchemaRegistry = false,
    ///             UnmanagedTopics = false,
    ///             Version = "2.8",
    ///             Zones = new[]
    ///             {
    ///                 "ru-central1-a",
    ///             },
    ///         },
    ///         Environment = "PRESTABLE",
    ///         NetworkId = fooVpcNetwork.Id,
    ///         SubnetIds = new[]
    ///         {
    ///             fooVpcSubnet.Id,
    ///         },
    ///         Users = new[]
    ///         {
    ///             new Yandex.Inputs.MdbKafkaClusterUserArgs
    ///             {
    ///                 Name = "producer-application",
    ///                 Password = "password",
    ///                 Permissions = new[]
    ///                 {
    ///                     new Yandex.Inputs.MdbKafkaClusterUserPermissionArgs
    ///                     {
    ///                         AllowHosts = new[]
    ///                         {
    ///                             "host1.db.yandex.net",
    ///                             "host2.db.yandex.net",
    ///                         },
    ///                         Role = "ACCESS_ROLE_PRODUCER",
    ///                         TopicName = "input",
    ///                     },
    ///                 },
    ///             },
    ///             new Yandex.Inputs.MdbKafkaClusterUserArgs
    ///             {
    ///                 Name = "worker",
    ///                 Password = "password",
    ///                 Permissions = new[]
    ///                 {
    ///                     new Yandex.Inputs.MdbKafkaClusterUserPermissionArgs
    ///                     {
    ///                         Role = "ACCESS_ROLE_CONSUMER",
    ///                         TopicName = "input",
    ///                     },
    ///                     new Yandex.Inputs.MdbKafkaClusterUserPermissionArgs
    ///                     {
    ///                         Role = "ACCESS_ROLE_PRODUCER",
    ///                         TopicName = "output",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Example of creating a HA Kafka Cluster with two brokers per AZ (6 brokers + 3 zk)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fooVpcNetwork = new Yandex.VpcNetwork("fooVpcNetwork");
    /// 
    ///     var fooVpcSubnet = new Yandex.VpcSubnet("fooVpcSubnet", new()
    ///     {
    ///         NetworkId = fooVpcNetwork.Id,
    ///         V4CidrBlocks = new[]
    ///         {
    ///             "10.1.0.0/24",
    ///         },
    ///         Zone = "ru-central1-a",
    ///     });
    /// 
    ///     var bar = new Yandex.VpcSubnet("bar", new()
    ///     {
    ///         NetworkId = fooVpcNetwork.Id,
    ///         V4CidrBlocks = new[]
    ///         {
    ///             "10.2.0.0/24",
    ///         },
    ///         Zone = "ru-central1-b",
    ///     });
    /// 
    ///     var baz = new Yandex.VpcSubnet("baz", new()
    ///     {
    ///         NetworkId = fooVpcNetwork.Id,
    ///         V4CidrBlocks = new[]
    ///         {
    ///             "10.3.0.0/24",
    ///         },
    ///         Zone = "ru-central1-c",
    ///     });
    /// 
    ///     var fooMdbKafkaCluster = new Yandex.MdbKafkaCluster("fooMdbKafkaCluster", new()
    ///     {
    ///         Config = new Yandex.Inputs.MdbKafkaClusterConfigArgs
    ///         {
    ///             AssignPublicIp = true,
    ///             BrokersCount = 2,
    ///             Kafka = new Yandex.Inputs.MdbKafkaClusterConfigKafkaArgs
    ///             {
    ///                 KafkaConfig = new Yandex.Inputs.MdbKafkaClusterConfigKafkaKafkaConfigArgs
    ///                 {
    ///                     CompressionType = "COMPRESSION_TYPE_ZSTD",
    ///                     DefaultReplicationFactor = "6",
    ///                     LogFlushIntervalMessages = "1024",
    ///                     LogFlushIntervalMs = "1000",
    ///                     LogFlushSchedulerIntervalMs = "1000",
    ///                     LogPreallocate = true,
    ///                     LogRetentionBytes = "1073741824",
    ///                     LogRetentionHours = "168",
    ///                     LogRetentionMinutes = "10080",
    ///                     LogRetentionMs = "86400000",
    ///                     LogSegmentBytes = "134217728",
    ///                     MessageMaxBytes = "1048588",
    ///                     NumPartitions = "10",
    ///                     OffsetsRetentionMinutes = "10080",
    ///                     ReplicaFetchMaxBytes = "1048576",
    ///                     SaslEnabledMechanisms = new[]
    ///                     {
    ///                         "SASL_MECHANISM_SCRAM_SHA_256",
    ///                         "SASL_MECHANISM_SCRAM_SHA_512",
    ///                     },
    ///                     SslCipherSuites = new[]
    ///                     {
    ///                         "TLS_DHE_RSA_WITH_AES_128_CBC_SHA",
    ///                         "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
    ///                     },
    ///                 },
    ///                 Resources = new Yandex.Inputs.MdbKafkaClusterConfigKafkaResourcesArgs
    ///                 {
    ///                     DiskSize = 128,
    ///                     DiskTypeId = "network-ssd",
    ///                     ResourcePresetId = "s2.medium",
    ///                 },
    ///             },
    ///             SchemaRegistry = false,
    ///             UnmanagedTopics = false,
    ///             Version = "2.8",
    ///             Zones = new[]
    ///             {
    ///                 "ru-central1-a",
    ///                 "ru-central1-b",
    ///                 "ru-central1-c",
    ///             },
    ///             Zookeeper = new Yandex.Inputs.MdbKafkaClusterConfigZookeeperArgs
    ///             {
    ///                 Resources = new Yandex.Inputs.MdbKafkaClusterConfigZookeeperResourcesArgs
    ///                 {
    ///                     DiskSize = 20,
    ///                     DiskTypeId = "network-ssd",
    ///                     ResourcePresetId = "s2.micro",
    ///                 },
    ///             },
    ///         },
    ///         Environment = "PRESTABLE",
    ///         NetworkId = fooVpcNetwork.Id,
    ///         SubnetIds = new[]
    ///         {
    ///             fooVpcSubnet.Id,
    ///             bar.Id,
    ///             baz.Id,
    ///         },
    ///         Users = new[]
    ///         {
    ///             new Yandex.Inputs.MdbKafkaClusterUserArgs
    ///             {
    ///                 Name = "producer-application",
    ///                 Password = "password",
    ///                 Permissions = new[]
    ///                 {
    ///                     new Yandex.Inputs.MdbKafkaClusterUserPermissionArgs
    ///                     {
    ///                         AllowHosts = new[]
    ///                         {
    ///                             "host1.db.yandex.net",
    ///                             "host2.db.yandex.net",
    ///                         },
    ///                         Role = "ACCESS_ROLE_PRODUCER",
    ///                         TopicName = "input",
    ///                     },
    ///                 },
    ///             },
    ///             new Yandex.Inputs.MdbKafkaClusterUserArgs
    ///             {
    ///                 Name = "worker",
    ///                 Password = "password",
    ///                 Permissions = new[]
    ///                 {
    ///                     new Yandex.Inputs.MdbKafkaClusterUserPermissionArgs
    ///                     {
    ///                         Role = "ACCESS_ROLE_CONSUMER",
    ///                         TopicName = "input",
    ///                     },
    ///                     new Yandex.Inputs.MdbKafkaClusterUserPermissionArgs
    ///                     {
    ///                         Role = "ACCESS_ROLE_PRODUCER",
    ///                         TopicName = "output",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A cluster can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import yandex:index/mdbKafkaCluster:MdbKafkaCluster foo cluster_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/mdbKafkaCluster:MdbKafkaCluster")]
    public partial class MdbKafkaCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Output("config")]
        public Output<Outputs.MdbKafkaClusterConfig> Config { get; private set; } = null!;

        /// <summary>
        /// Timestamp of cluster creation.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Inhibits deletion of the cluster.  Can be either `true` or `false`.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Description of the Kafka cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Deployment environment of the Kafka cluster. Can be either `PRESTABLE` or `PRODUCTION`. 
        /// The default is `PRODUCTION`.
        /// </summary>
        [Output("environment")]
        public Output<string?> Environment { get; private set; } = null!;

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Health of the host.
        /// </summary>
        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        /// <summary>
        /// A list of IDs of the host groups to place VMs of the cluster on.
        /// </summary>
        [Output("hostGroupIds")]
        public Output<ImmutableArray<string>> HostGroupIds { get; private set; } = null!;

        /// <summary>
        /// A host of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<Outputs.MdbKafkaClusterHost>> Hosts { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the Kafka cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Maintenance policy of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.MdbKafkaClusterMaintenanceWindow> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the network, to which the Kafka cluster belongs.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Security group ids, to which the Kafka cluster belongs.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
        /// For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-kafka/api-ref/Cluster/).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// IDs of the subnets, to which the Kafka cluster belongs.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// To manage topics, please switch to using a separate resource type `yandex.MdbKafkaTopic`.
        /// </summary>
        [Output("topics")]
        public Output<ImmutableArray<Outputs.MdbKafkaClusterTopic>> Topics { get; private set; } = null!;

        /// <summary>
        /// To manage users, please switch to using a separate resource type `yandex.mdbKafkaUser`.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.MdbKafkaClusterUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a MdbKafkaCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbKafkaCluster(string name, MdbKafkaClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbKafkaCluster:MdbKafkaCluster", name, args ?? new MdbKafkaClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbKafkaCluster(string name, Input<string> id, MdbKafkaClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbKafkaCluster:MdbKafkaCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/regrau/pulumi-yandex/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MdbKafkaCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbKafkaCluster Get(string name, Input<string> id, MdbKafkaClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbKafkaCluster(name, id, state, options);
        }
    }

    public sealed class MdbKafkaClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.MdbKafkaClusterConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// Inhibits deletion of the cluster.  Can be either `true` or `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the Kafka cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deployment environment of the Kafka cluster. Can be either `PRESTABLE` or `PRODUCTION`. 
        /// The default is `PRODUCTION`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("hostGroupIds")]
        private InputList<string>? _hostGroupIds;

        /// <summary>
        /// A list of IDs of the host groups to place VMs of the cluster on.
        /// </summary>
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Kafka cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maintenance policy of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbKafkaClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network, to which the Kafka cluster belongs.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group ids, to which the Kafka cluster belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// IDs of the subnets, to which the Kafka cluster belongs.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("topics")]
        private InputList<Inputs.MdbKafkaClusterTopicArgs>? _topics;

        /// <summary>
        /// To manage topics, please switch to using a separate resource type `yandex.MdbKafkaTopic`.
        /// </summary>
        [Obsolete(@"to manage topics, please switch to using a separate resource type yandex_mdb_kafka_topic")]
        public InputList<Inputs.MdbKafkaClusterTopicArgs> Topics
        {
            get => _topics ?? (_topics = new InputList<Inputs.MdbKafkaClusterTopicArgs>());
            set => _topics = value;
        }

        [Input("users")]
        private InputList<Inputs.MdbKafkaClusterUserArgs>? _users;

        /// <summary>
        /// To manage users, please switch to using a separate resource type `yandex.mdbKafkaUser`.
        /// </summary>
        [Obsolete(@"to manage users, please switch to using a separate resource type yandex_mdb_kafka_user")]
        public InputList<Inputs.MdbKafkaClusterUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.MdbKafkaClusterUserArgs>());
            set => _users = value;
        }

        public MdbKafkaClusterArgs()
        {
        }
        public static new MdbKafkaClusterArgs Empty => new MdbKafkaClusterArgs();
    }

    public sealed class MdbKafkaClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Input("config")]
        public Input<Inputs.MdbKafkaClusterConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// Timestamp of cluster creation.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Inhibits deletion of the cluster.  Can be either `true` or `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the Kafka cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deployment environment of the Kafka cluster. Can be either `PRESTABLE` or `PRODUCTION`. 
        /// The default is `PRODUCTION`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Health of the host.
        /// </summary>
        [Input("health")]
        public Input<string>? Health { get; set; }

        [Input("hostGroupIds")]
        private InputList<string>? _hostGroupIds;

        /// <summary>
        /// A list of IDs of the host groups to place VMs of the cluster on.
        /// </summary>
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("hosts")]
        private InputList<Inputs.MdbKafkaClusterHostGetArgs>? _hosts;

        /// <summary>
        /// A host of the Kafka cluster. The structure is documented below.
        /// </summary>
        public InputList<Inputs.MdbKafkaClusterHostGetArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbKafkaClusterHostGetArgs>());
            set => _hosts = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Kafka cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maintenance policy of the Kafka cluster. The structure is documented below.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbKafkaClusterMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network, to which the Kafka cluster belongs.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group ids, to which the Kafka cluster belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
        /// For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-kafka/api-ref/Cluster/).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// IDs of the subnets, to which the Kafka cluster belongs.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("topics")]
        private InputList<Inputs.MdbKafkaClusterTopicGetArgs>? _topics;

        /// <summary>
        /// To manage topics, please switch to using a separate resource type `yandex.MdbKafkaTopic`.
        /// </summary>
        [Obsolete(@"to manage topics, please switch to using a separate resource type yandex_mdb_kafka_topic")]
        public InputList<Inputs.MdbKafkaClusterTopicGetArgs> Topics
        {
            get => _topics ?? (_topics = new InputList<Inputs.MdbKafkaClusterTopicGetArgs>());
            set => _topics = value;
        }

        [Input("users")]
        private InputList<Inputs.MdbKafkaClusterUserGetArgs>? _users;

        /// <summary>
        /// To manage users, please switch to using a separate resource type `yandex.mdbKafkaUser`.
        /// </summary>
        [Obsolete(@"to manage users, please switch to using a separate resource type yandex_mdb_kafka_user")]
        public InputList<Inputs.MdbKafkaClusterUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.MdbKafkaClusterUserGetArgs>());
            set => _users = value;
        }

        public MdbKafkaClusterState()
        {
        }
        public static new MdbKafkaClusterState Empty => new MdbKafkaClusterState();
    }
}
