// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupHttpBackendTlsArgs : global::Pulumi.InvokeArgs
    {
        [Input("sni", required: true)]
        public string Sni { get; set; } = null!;

        [Input("validationContext", required: true)]
        public Inputs.GetAlbBackendGroupHttpBackendTlsValidationContextArgs ValidationContext { get; set; } = null!;

        public GetAlbBackendGroupHttpBackendTlsArgs()
        {
        }
        public static new GetAlbBackendGroupHttpBackendTlsArgs Empty => new GetAlbBackendGroupHttpBackendTlsArgs();
    }
}