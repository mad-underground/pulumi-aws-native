// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardValidationStrategy
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardValidationStrategyMode Mode;

        [OutputConstructor]
        private DashboardValidationStrategy(Pulumi.AwsNative.QuickSight.DashboardValidationStrategyMode mode)
        {
            Mode = mode;
        }
    }
}
