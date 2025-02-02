// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterUserQuotaInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of queries that threw exception.
        /// </summary>
        [Input("errors", required: true)]
        public Input<int> Errors { get; set; } = null!;

        /// <summary>
        /// The total query execution time, in milliseconds (wall time).
        /// </summary>
        [Input("executionTime", required: true)]
        public Input<int> ExecutionTime { get; set; } = null!;

        /// <summary>
        /// Duration of interval for quota in milliseconds.
        /// </summary>
        [Input("intervalDuration")]
        public Input<int>? IntervalDuration { get; set; }

        /// <summary>
        /// The total number of queries.
        /// </summary>
        [Input("queries", required: true)]
        public Input<int> Queries { get; set; } = null!;

        /// <summary>
        /// The total number of source rows read from tables for running the query, on all remote servers.
        /// </summary>
        [Input("readRows", required: true)]
        public Input<int> ReadRows { get; set; } = null!;

        /// <summary>
        /// The total number of rows given as the result.
        /// </summary>
        [Input("resultRows", required: true)]
        public Input<int> ResultRows { get; set; } = null!;

        public GetMdbClickhouseClusterUserQuotaInputArgs()
        {
        }
        public static new GetMdbClickhouseClusterUserQuotaInputArgs Empty => new GetMdbClickhouseClusterUserQuotaInputArgs();
    }
}
