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
    public sealed class AnalysisItemsLimitConfiguration
    {
        public readonly double? ItemsLimit;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisOtherCategories? OtherCategories;

        [OutputConstructor]
        private AnalysisItemsLimitConfiguration(
            double? itemsLimit,

            Pulumi.AwsNative.QuickSight.AnalysisOtherCategories? otherCategories)
        {
            ItemsLimit = itemsLimit;
            OtherCategories = otherCategories;
        }
    }
}
