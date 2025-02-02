// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class StorageBucketLifecycleRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
        /// </summary>
        [Input("abortIncompleteMultipartUploadDays")]
        public Input<int>? AbortIncompleteMultipartUploadDays { get; set; }

        /// <summary>
        /// Specifies lifecycle rule status.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Specifies a period in the object's expire (documented below).
        /// </summary>
        [Input("expiration")]
        public Input<Inputs.StorageBucketLifecycleRuleExpirationArgs>? Expiration { get; set; }

        /// <summary>
        /// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies when noncurrent object versions expire (documented below).
        /// </summary>
        [Input("noncurrentVersionExpiration")]
        public Input<Inputs.StorageBucketLifecycleRuleNoncurrentVersionExpirationArgs>? NoncurrentVersionExpiration { get; set; }

        [Input("noncurrentVersionTransitions")]
        private InputList<Inputs.StorageBucketLifecycleRuleNoncurrentVersionTransitionArgs>? _noncurrentVersionTransitions;

        /// <summary>
        /// Specifies when noncurrent object versions transitions (documented below).
        /// </summary>
        public InputList<Inputs.StorageBucketLifecycleRuleNoncurrentVersionTransitionArgs> NoncurrentVersionTransitions
        {
            get => _noncurrentVersionTransitions ?? (_noncurrentVersionTransitions = new InputList<Inputs.StorageBucketLifecycleRuleNoncurrentVersionTransitionArgs>());
            set => _noncurrentVersionTransitions = value;
        }

        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("transitions")]
        private InputList<Inputs.StorageBucketLifecycleRuleTransitionArgs>? _transitions;

        /// <summary>
        /// Specifies a period in the object's transitions (documented below).
        /// </summary>
        public InputList<Inputs.StorageBucketLifecycleRuleTransitionArgs> Transitions
        {
            get => _transitions ?? (_transitions = new InputList<Inputs.StorageBucketLifecycleRuleTransitionArgs>());
            set => _transitions = value;
        }

        public StorageBucketLifecycleRuleArgs()
        {
        }
        public static new StorageBucketLifecycleRuleArgs Empty => new StorageBucketLifecycleRuleArgs();
    }
}
