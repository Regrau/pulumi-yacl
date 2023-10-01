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
    public sealed class YandexYdbTableColumn
    {
        public readonly string? Family;
        public readonly string Name;
        public readonly bool? NotNull;
        public readonly string Type;

        [OutputConstructor]
        private YandexYdbTableColumn(
            string? family,

            string name,

            bool? notNull,

            string type)
        {
            Family = family;
            Name = name;
            NotNull = notNull;
            Type = type;
        }
    }
}