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
    public sealed class DatatransferEndpointSettings
    {
        public readonly Outputs.DatatransferEndpointSettingsClickhouseSource? ClickhouseSource;
        public readonly Outputs.DatatransferEndpointSettingsClickhouseTarget? ClickhouseTarget;
        public readonly Outputs.DatatransferEndpointSettingsMongoSource? MongoSource;
        public readonly Outputs.DatatransferEndpointSettingsMongoTarget? MongoTarget;
        public readonly Outputs.DatatransferEndpointSettingsMysqlSource? MysqlSource;
        public readonly Outputs.DatatransferEndpointSettingsMysqlTarget? MysqlTarget;
        public readonly Outputs.DatatransferEndpointSettingsPostgresSource? PostgresSource;
        public readonly Outputs.DatatransferEndpointSettingsPostgresTarget? PostgresTarget;

        [OutputConstructor]
        private DatatransferEndpointSettings(
            Outputs.DatatransferEndpointSettingsClickhouseSource? clickhouseSource,

            Outputs.DatatransferEndpointSettingsClickhouseTarget? clickhouseTarget,

            Outputs.DatatransferEndpointSettingsMongoSource? mongoSource,

            Outputs.DatatransferEndpointSettingsMongoTarget? mongoTarget,

            Outputs.DatatransferEndpointSettingsMysqlSource? mysqlSource,

            Outputs.DatatransferEndpointSettingsMysqlTarget? mysqlTarget,

            Outputs.DatatransferEndpointSettingsPostgresSource? postgresSource,

            Outputs.DatatransferEndpointSettingsPostgresTarget? postgresTarget)
        {
            ClickhouseSource = clickhouseSource;
            ClickhouseTarget = clickhouseTarget;
            MongoSource = mongoSource;
            MongoTarget = mongoTarget;
            MysqlSource = mysqlSource;
            MysqlTarget = mysqlTarget;
            PostgresSource = postgresSource;
            PostgresTarget = postgresTarget;
        }
    }
}