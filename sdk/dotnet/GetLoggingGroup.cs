// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetLoggingGroup
    {
        public static Task<GetLoggingGroupResult> InvokeAsync(GetLoggingGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoggingGroupResult>("yandex:index/getLoggingGroup:getLoggingGroup", args ?? new GetLoggingGroupArgs(), options.WithDefaults());

        public static Output<GetLoggingGroupResult> Invoke(GetLoggingGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLoggingGroupResult>("yandex:index/getLoggingGroup:getLoggingGroup", args ?? new GetLoggingGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoggingGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("groupId")]
        public string? GroupId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetLoggingGroupArgs()
        {
        }
        public static new GetLoggingGroupArgs Empty => new GetLoggingGroupArgs();
    }

    public sealed class GetLoggingGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetLoggingGroupInvokeArgs()
        {
        }
        public static new GetLoggingGroupInvokeArgs Empty => new GetLoggingGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoggingGroupResult
    {
        public readonly string CloudId;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string FolderId;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly string RetentionPeriod;
        public readonly string Status;

        [OutputConstructor]
        private GetLoggingGroupResult(
            string cloudId,

            string createdAt,

            string description,

            string folderId,

            string groupId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string retentionPeriod,

            string status)
        {
            CloudId = cloudId;
            CreatedAt = createdAt;
            Description = description;
            FolderId = folderId;
            GroupId = groupId;
            Id = id;
            Labels = labels;
            Name = name;
            RetentionPeriod = retentionPeriod;
            Status = status;
        }
    }
}