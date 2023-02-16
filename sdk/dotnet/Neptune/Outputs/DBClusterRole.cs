// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Neptune.Outputs
{

    [OutputType]
    public sealed class DBClusterRole
    {
        public readonly string? FeatureName;
        public readonly string RoleArn;

        [OutputConstructor]
        private DBClusterRole(
            string? featureName,

            string roleArn)
        {
            FeatureName = featureName;
            RoleArn = roleArn;
        }
    }
}
