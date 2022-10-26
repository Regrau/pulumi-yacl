// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerTlsSniHandlerArgs : global::Pulumi.ResourceArgs
    {
        [Input("handler", required: true)]
        public Input<Inputs.AlbLoadBalancerListenerTlsSniHandlerHandlerArgs> Handler { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("serverNames", required: true)]
        private InputList<string>? _serverNames;
        public InputList<string> ServerNames
        {
            get => _serverNames ?? (_serverNames = new InputList<string>());
            set => _serverNames = value;
        }

        public AlbLoadBalancerListenerTlsSniHandlerArgs()
        {
        }
        public static new AlbLoadBalancerListenerTlsSniHandlerArgs Empty => new AlbLoadBalancerListenerTlsSniHandlerArgs();
    }
}