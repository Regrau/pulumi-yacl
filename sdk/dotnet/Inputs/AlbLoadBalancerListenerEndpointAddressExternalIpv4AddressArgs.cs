// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressArgs : global::Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        public AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressArgs()
        {
        }
        public static new AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressArgs Empty => new AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressArgs();
    }
}