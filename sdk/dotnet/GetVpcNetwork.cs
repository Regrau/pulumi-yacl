// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetVpcNetwork
    {
        public static Task<GetVpcNetworkResult> InvokeAsync(GetVpcNetworkArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcNetworkResult>("yandex:index/getVpcNetwork:getVpcNetwork", args ?? new GetVpcNetworkArgs(), options.WithDefaults());

        public static Output<GetVpcNetworkResult> Invoke(GetVpcNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpcNetworkResult>("yandex:index/getVpcNetwork:getVpcNetwork", args ?? new GetVpcNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcNetworkArgs : global::Pulumi.InvokeArgs
    {
        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("networkId")]
        public string? NetworkId { get; set; }

        public GetVpcNetworkArgs()
        {
        }
        public static new GetVpcNetworkArgs Empty => new GetVpcNetworkArgs();
    }

    public sealed class GetVpcNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        public GetVpcNetworkInvokeArgs()
        {
        }
        public static new GetVpcNetworkInvokeArgs Empty => new GetVpcNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcNetworkResult
    {
        public readonly string CreatedAt;
        public readonly string DefaultSecurityGroupId;
        public readonly string Description;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly string NetworkId;
        public readonly ImmutableArray<string> SubnetIds;

        [OutputConstructor]
        private GetVpcNetworkResult(
            string createdAt,

            string defaultSecurityGroupId,

            string description,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string networkId,

            ImmutableArray<string> subnetIds)
        {
            CreatedAt = createdAt;
            DefaultSecurityGroupId = defaultSecurityGroupId;
            Description = description;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            NetworkId = networkId;
            SubnetIds = subnetIds;
        }
    }
}