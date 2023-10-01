// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsYdbSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database to transfer.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("paths")]
        private InputList<string>? _paths;
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        [Input("saKeyContent")]
        private Input<string>? _saKeyContent;
        public Input<string>? SaKeyContent
        {
            get => _saKeyContent;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _saKeyContent = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of security groups that the transfer associated with this endpoint should use.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        /// <summary>
        /// Identifier of the Yandex Cloud VPC subnetwork to user for accessing the database. If omitted, the server has to be accessible via Internet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public DatatransferEndpointSettingsYdbSourceArgs()
        {
        }
        public static new DatatransferEndpointSettingsYdbSourceArgs Empty => new DatatransferEndpointSettingsYdbSourceArgs();
    }
}