// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbMongodbClusterClusterConfigMongosArgs : global::Pulumi.InvokeArgs
    {
        [Input("net")]
        public Inputs.GetMdbMongodbClusterClusterConfigMongosNetArgs? Net { get; set; }

        public GetMdbMongodbClusterClusterConfigMongosArgs()
        {
        }
        public static new GetMdbMongodbClusterClusterConfigMongosArgs Empty => new GetMdbMongodbClusterClusterConfigMongosArgs();
    }
}
