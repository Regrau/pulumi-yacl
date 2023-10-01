// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMongodbClusterClusterConfigMongosNetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of simultaneous connections that host will accept.
        /// For more information, see the [net.maxIncomingConnections](https://www.mongodb.com/docs/manual/reference/configuration-options/#mongodb-setting-net.maxIncomingConnections)
        /// description in the official documentation.
        /// </summary>
        [Input("maxIncomingConnections")]
        public Input<int>? MaxIncomingConnections { get; set; }

        public MdbMongodbClusterClusterConfigMongosNetGetArgs()
        {
        }
        public static new MdbMongodbClusterClusterConfigMongosNetGetArgs Empty => new MdbMongodbClusterClusterConfigMongosNetGetArgs();
    }
}