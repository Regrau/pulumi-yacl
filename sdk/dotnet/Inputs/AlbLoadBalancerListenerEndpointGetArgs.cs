// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerEndpointGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("addresses", required: true)]
        private InputList<Inputs.AlbLoadBalancerListenerEndpointAddressGetArgs>? _addresses;
        public InputList<Inputs.AlbLoadBalancerListenerEndpointAddressGetArgs> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<Inputs.AlbLoadBalancerListenerEndpointAddressGetArgs>());
            set => _addresses = value;
        }

        [Input("ports", required: true)]
        private InputList<int>? _ports;
        public InputList<int> Ports
        {
            get => _ports ?? (_ports = new InputList<int>());
            set => _ports = value;
        }

        public AlbLoadBalancerListenerEndpointGetArgs()
        {
        }
        public static new AlbLoadBalancerListenerEndpointGetArgs Empty => new AlbLoadBalancerListenerEndpointGetArgs();
    }
}