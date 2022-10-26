// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesNodeGroupInstanceTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("bootDisk")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateBootDiskArgs>? BootDisk { get; set; }

        [Input("containerRuntime")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateContainerRuntimeArgs>? ContainerRuntime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("metadata")]
        private InputMap<string>? _metadata;
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nat")]
        public Input<bool>? Nat { get; set; }

        [Input("networkAccelerationType")]
        public Input<string>? NetworkAccelerationType { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceArgs>? _networkInterfaces;
        public InputList<Inputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("placementPolicy")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplatePlacementPolicyArgs>? PlacementPolicy { get; set; }

        [Input("platformId")]
        public Input<string>? PlatformId { get; set; }

        [Input("resources")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateResourcesArgs>? Resources { get; set; }

        [Input("schedulingPolicy")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateSchedulingPolicyArgs>? SchedulingPolicy { get; set; }

        public KubernetesNodeGroupInstanceTemplateArgs()
        {
        }
        public static new KubernetesNodeGroupInstanceTemplateArgs Empty => new KubernetesNodeGroupInstanceTemplateArgs();
    }
}