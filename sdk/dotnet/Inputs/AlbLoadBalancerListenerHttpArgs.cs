// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerHttpArgs : global::Pulumi.ResourceArgs
    {
        [Input("handler")]
        public Input<Inputs.AlbLoadBalancerListenerHttpHandlerArgs>? Handler { get; set; }

        [Input("redirects")]
        public Input<Inputs.AlbLoadBalancerListenerHttpRedirectsArgs>? Redirects { get; set; }

        public AlbLoadBalancerListenerHttpArgs()
        {
        }
        public static new AlbLoadBalancerListenerHttpArgs Empty => new AlbLoadBalancerListenerHttpArgs();
    }
}