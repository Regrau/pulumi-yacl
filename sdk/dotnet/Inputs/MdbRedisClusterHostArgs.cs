// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbRedisClusterHostArgs : global::Pulumi.ResourceArgs
    {
        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        [Input("replicaPriority")]
        public Input<int>? ReplicaPriority { get; set; }

        [Input("shardName")]
        public Input<string>? ShardName { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public MdbRedisClusterHostArgs()
        {
        }
        public static new MdbRedisClusterHostArgs Empty => new MdbRedisClusterHostArgs();
    }
}