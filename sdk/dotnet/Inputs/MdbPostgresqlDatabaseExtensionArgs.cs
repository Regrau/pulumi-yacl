// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbPostgresqlDatabaseExtensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database extension. For more information on available extensions see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-extensions).
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Version of the extension.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public MdbPostgresqlDatabaseExtensionArgs()
        {
        }
        public static new MdbPostgresqlDatabaseExtensionArgs Empty => new MdbPostgresqlDatabaseExtensionArgs();
    }
}