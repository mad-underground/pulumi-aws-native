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
    public sealed class DashboardKpiSparklineOptions
    {
        public readonly string? Color;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? TooltipVisibility;
        public readonly Pulumi.AwsNative.QuickSight.DashboardKpiSparklineType Type;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? Visibility;

        [OutputConstructor]
        private DashboardKpiSparklineOptions(
            string? color,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? tooltipVisibility,

            Pulumi.AwsNative.QuickSight.DashboardKpiSparklineType type,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? visibility)
        {
            Color = color;
            TooltipVisibility = tooltipVisibility;
            Type = type;
            Visibility = visibility;
        }
    }
}
