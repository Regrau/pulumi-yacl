// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetMdbElasticSearchClusterConfigResult
    {
        public readonly string AdminPassword;
        public readonly ImmutableArray<Outputs.GetMdbElasticSearchClusterConfigDataNodeResult> DataNodes;
        public readonly string Edition;
        public readonly Outputs.GetMdbElasticSearchClusterConfigMasterNodeResult MasterNode;
        public readonly ImmutableArray<string> Plugins;
        public readonly string Version;

        [OutputConstructor]
        private GetMdbElasticSearchClusterConfigResult(
            string adminPassword,

            ImmutableArray<Outputs.GetMdbElasticSearchClusterConfigDataNodeResult> dataNodes,

            string edition,

            Outputs.GetMdbElasticSearchClusterConfigMasterNodeResult masterNode,

            ImmutableArray<string> plugins,

            string version)
        {
            AdminPassword = adminPassword;
            DataNodes = dataNodes;
            Edition = edition;
            MasterNode = masterNode;
            Plugins = plugins;
            Version = version;
        }
    }
}