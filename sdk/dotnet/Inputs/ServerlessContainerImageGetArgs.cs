// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ServerlessContainerImageGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("digest")]
        public Input<string>? Digest { get; set; }

        [Input("environment")]
        private InputMap<string>? _environment;
        public InputMap<string> Environment
        {
            get => _environment ?? (_environment = new InputMap<string>());
            set => _environment = value;
        }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("workDir")]
        public Input<string>? WorkDir { get; set; }

        public ServerlessContainerImageGetArgs()
        {
        }
        public static new ServerlessContainerImageGetArgs Empty => new ServerlessContainerImageGetArgs();
    }
}