// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupGrpcBackendInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("healthcheck", required: true)]
        public Input<Inputs.GetAlbBackendGroupGrpcBackendHealthcheckInputArgs> Healthcheck { get; set; } = null!;

        [Input("loadBalancingConfig", required: true)]
        public Input<Inputs.GetAlbBackendGroupGrpcBackendLoadBalancingConfigInputArgs> LoadBalancingConfig { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("targetGroupIds", required: true)]
        private InputList<string>? _targetGroupIds;
        public InputList<string> TargetGroupIds
        {
            get => _targetGroupIds ?? (_targetGroupIds = new InputList<string>());
            set => _targetGroupIds = value;
        }

        [Input("tls", required: true)]
        public Input<Inputs.GetAlbBackendGroupGrpcBackendTlsInputArgs> Tls { get; set; } = null!;

        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public GetAlbBackendGroupGrpcBackendInputArgs()
        {
        }
        public static new GetAlbBackendGroupGrpcBackendInputArgs Empty => new GetAlbBackendGroupGrpcBackendInputArgs();
    }
}