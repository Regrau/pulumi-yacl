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
    public sealed class ComputeInstanceGroupDeployPolicy
    {
        public readonly int? MaxCreating;
        public readonly int? MaxDeleting;
        public readonly int MaxExpansion;
        public readonly int MaxUnavailable;
        public readonly int? StartupDuration;
        public readonly string? Strategy;

        [OutputConstructor]
        private ComputeInstanceGroupDeployPolicy(
            int? maxCreating,

            int? maxDeleting,

            int maxExpansion,

            int maxUnavailable,

            int? startupDuration,

            string? strategy)
        {
            MaxCreating = maxCreating;
            MaxDeleting = maxDeleting;
            MaxExpansion = maxExpansion;
            MaxUnavailable = maxUnavailable;
            StartupDuration = startupDuration;
            Strategy = strategy;
        }
    }
}