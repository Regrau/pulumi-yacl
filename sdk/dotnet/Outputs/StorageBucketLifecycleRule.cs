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
    public sealed class StorageBucketLifecycleRule
    {
        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
        /// </summary>
        public readonly int? AbortIncompleteMultipartUploadDays;
        /// <summary>
        /// Specifies lifecycle rule status.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Specifies a period in the object's expire (documented below).
        /// </summary>
        public readonly Outputs.StorageBucketLifecycleRuleExpiration? Expiration;
        /// <summary>
        /// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Specifies when noncurrent object versions expire (documented below).
        /// </summary>
        public readonly Outputs.StorageBucketLifecycleRuleNoncurrentVersionExpiration? NoncurrentVersionExpiration;
        /// <summary>
        /// Specifies when noncurrent object versions transitions (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageBucketLifecycleRuleNoncurrentVersionTransition> NoncurrentVersionTransitions;
        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        public readonly string? Prefix;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies a period in the object's transitions (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageBucketLifecycleRuleTransition> Transitions;

        [OutputConstructor]
        private StorageBucketLifecycleRule(
            int? abortIncompleteMultipartUploadDays,

            bool enabled,

            Outputs.StorageBucketLifecycleRuleExpiration? expiration,

            string? id,

            Outputs.StorageBucketLifecycleRuleNoncurrentVersionExpiration? noncurrentVersionExpiration,

            ImmutableArray<Outputs.StorageBucketLifecycleRuleNoncurrentVersionTransition> noncurrentVersionTransitions,

            string? prefix,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.StorageBucketLifecycleRuleTransition> transitions)
        {
            AbortIncompleteMultipartUploadDays = abortIncompleteMultipartUploadDays;
            Enabled = enabled;
            Expiration = expiration;
            Id = id;
            NoncurrentVersionExpiration = noncurrentVersionExpiration;
            NoncurrentVersionTransitions = noncurrentVersionTransitions;
            Prefix = prefix;
            Tags = tags;
            Transitions = transitions;
        }
    }
}
