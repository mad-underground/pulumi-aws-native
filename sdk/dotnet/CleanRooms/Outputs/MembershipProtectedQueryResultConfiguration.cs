// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class MembershipProtectedQueryResultConfiguration
    {
        public readonly Outputs.MembershipProtectedQueryOutputConfiguration OutputConfiguration;
        public readonly string? RoleArn;

        [OutputConstructor]
        private MembershipProtectedQueryResultConfiguration(
            Outputs.MembershipProtectedQueryOutputConfiguration outputConfiguration,

            string? roleArn)
        {
            OutputConfiguration = outputConfiguration;
            RoleArn = roleArn;
        }
    }
}
