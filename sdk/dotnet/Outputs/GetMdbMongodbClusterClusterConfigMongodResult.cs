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
    public sealed class GetMdbMongodbClusterClusterConfigMongodResult
    {
        public readonly Outputs.GetMdbMongodbClusterClusterConfigMongodAuditLogResult AuditLog;
        public readonly Outputs.GetMdbMongodbClusterClusterConfigMongodSecurityResult Security;
        public readonly Outputs.GetMdbMongodbClusterClusterConfigMongodSetParameterResult SetParameter;

        [OutputConstructor]
        private GetMdbMongodbClusterClusterConfigMongodResult(
            Outputs.GetMdbMongodbClusterClusterConfigMongodAuditLogResult auditLog,

            Outputs.GetMdbMongodbClusterClusterConfigMongodSecurityResult security,

            Outputs.GetMdbMongodbClusterClusterConfigMongodSetParameterResult setParameter)
        {
            AuditLog = auditLog;
            Security = security;
            SetParameter = setParameter;
        }
    }
}