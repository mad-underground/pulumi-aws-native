// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MWAA.Outputs
{

    [OutputType]
    public sealed class EnvironmentNetworkConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-networkconfiguration.html#cfn-mwaa-environment-networkconfiguration-securitygroupids
        /// </summary>
        public readonly Outputs.EnvironmentSecurityGroupList? SecurityGroupIds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-networkconfiguration.html#cfn-mwaa-environment-networkconfiguration-subnetids
        /// </summary>
        public readonly Outputs.EnvironmentSubnetList? SubnetIds;

        [OutputConstructor]
        private EnvironmentNetworkConfiguration(
            Outputs.EnvironmentSecurityGroupList? securityGroupIds,

            Outputs.EnvironmentSubnetList? subnetIds)
        {
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
        }
    }
}
