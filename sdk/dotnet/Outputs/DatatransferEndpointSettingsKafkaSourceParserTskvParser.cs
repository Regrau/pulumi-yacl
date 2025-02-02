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
    public sealed class DatatransferEndpointSettingsKafkaSourceParserTskvParser
    {
        /// <summary>
        /// Add fields, that are not in the schema, into the _rest column.
        /// </summary>
        public readonly bool? AddRestColumn;
        /// <summary>
        /// Data parsing scheme.The structure is documented below.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsKafkaSourceParserTskvParserDataSchema? DataSchema;
        /// <summary>
        /// Allow null keys. If `false` - null keys will be putted to unparsed data
        /// </summary>
        public readonly bool? NullKeysAllowed;

        [OutputConstructor]
        private DatatransferEndpointSettingsKafkaSourceParserTskvParser(
            bool? addRestColumn,

            Outputs.DatatransferEndpointSettingsKafkaSourceParserTskvParserDataSchema? dataSchema,

            bool? nullKeysAllowed)
        {
            AddRestColumn = addRestColumn;
            DataSchema = dataSchema;
            NullKeysAllowed = nullKeysAllowed;
        }
    }
}
