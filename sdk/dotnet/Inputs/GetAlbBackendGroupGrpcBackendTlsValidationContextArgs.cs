// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupGrpcBackendTlsValidationContextInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("trustedCaBytes", required: true)]
        public Input<string> TrustedCaBytes { get; set; } = null!;

        [Input("trustedCaId", required: true)]
        public Input<string> TrustedCaId { get; set; } = null!;

        public GetAlbBackendGroupGrpcBackendTlsValidationContextInputArgs()
        {
        }
        public static new GetAlbBackendGroupGrpcBackendTlsValidationContextInputArgs Empty => new GetAlbBackendGroupGrpcBackendTlsValidationContextInputArgs();
    }
}