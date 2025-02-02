// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class YandexYdbTopicConsumerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Topic name. Type: string, required. Default value: "".
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("startingMessageTimestampMs")]
        public Input<int>? StartingMessageTimestampMs { get; set; }

        [Input("supportedCodecs")]
        private InputList<string>? _supportedCodecs;

        /// <summary>
        /// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
        /// </summary>
        public InputList<string> SupportedCodecs
        {
            get => _supportedCodecs ?? (_supportedCodecs = new InputList<string>());
            set => _supportedCodecs = value;
        }

        public YandexYdbTopicConsumerGetArgs()
        {
        }
        public static new YandexYdbTopicConsumerGetArgs Empty => new YandexYdbTopicConsumerGetArgs();
    }
}
