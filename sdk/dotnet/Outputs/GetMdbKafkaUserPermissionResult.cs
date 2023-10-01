// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetMdbKafkaUserPermissionResult
    {
        public readonly ImmutableArray<string> AllowHosts;
        public readonly string Role;
        public readonly string TopicName;

        [OutputConstructor]
        private GetMdbKafkaUserPermissionResult(
            ImmutableArray<string> allowHosts,

            string role,

            string topicName)
        {
            AllowHosts = allowHosts;
            Role = role;
            TopicName = topicName;
        }
    }
}