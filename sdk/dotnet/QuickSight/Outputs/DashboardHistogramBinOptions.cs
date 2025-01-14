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
    public sealed class DashboardHistogramBinOptions
    {
        public readonly Outputs.DashboardBinCountOptions? BinCount;
        public readonly Outputs.DashboardBinWidthOptions? BinWidth;
        public readonly Pulumi.AwsNative.QuickSight.DashboardHistogramBinType? SelectedBinType;
        public readonly double? StartValue;

        [OutputConstructor]
        private DashboardHistogramBinOptions(
            Outputs.DashboardBinCountOptions? binCount,

            Outputs.DashboardBinWidthOptions? binWidth,

            Pulumi.AwsNative.QuickSight.DashboardHistogramBinType? selectedBinType,

            double? startValue)
        {
            BinCount = binCount;
            BinWidth = binWidth;
            SelectedBinType = selectedBinType;
            StartValue = startValue;
        }
    }
}
