// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetComputeInstanceFilesystemInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the device.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        /// <summary>
        /// Access to the Disk resource. By default, a disk is attached in `READ_WRITE` mode.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        public GetComputeInstanceFilesystemInputArgs()
        {
        }
        public static new GetComputeInstanceFilesystemInputArgs Empty => new GetComputeInstanceFilesystemInputArgs();
    }
}