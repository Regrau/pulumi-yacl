// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class YandexYdbTableFamilyArgs : global::Pulumi.ResourceArgs
    {
        [Input("compression", required: true)]
        public Input<string> Compression { get; set; } = null!;

        [Input("data", required: true)]
        public Input<string> Data { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public YandexYdbTableFamilyArgs()
        {
        }
        public static new YandexYdbTableFamilyArgs Empty => new YandexYdbTableFamilyArgs();
    }
}