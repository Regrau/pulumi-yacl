// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DataprocClusterClusterConfigHadoopArgs : global::Pulumi.ResourceArgs
    {
        [Input("properties")]
        private InputMap<string>? _properties;
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("services")]
        private InputList<string>? _services;
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        [Input("sshPublicKeys")]
        private InputList<string>? _sshPublicKeys;
        public InputList<string> SshPublicKeys
        {
            get => _sshPublicKeys ?? (_sshPublicKeys = new InputList<string>());
            set => _sshPublicKeys = value;
        }

        public DataprocClusterClusterConfigHadoopArgs()
        {
        }
        public static new DataprocClusterClusterConfigHadoopArgs Empty => new DataprocClusterClusterConfigHadoopArgs();
    }
}