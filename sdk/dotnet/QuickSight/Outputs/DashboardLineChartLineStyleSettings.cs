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
    public sealed class DashboardLineChartLineStyleSettings
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardLineInterpolation? LineInterpolation;
        public readonly Pulumi.AwsNative.QuickSight.DashboardLineChartLineStyle? LineStyle;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? LineVisibility;
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? LineWidth;

        [OutputConstructor]
        private DashboardLineChartLineStyleSettings(
            Pulumi.AwsNative.QuickSight.DashboardLineInterpolation? lineInterpolation,

            Pulumi.AwsNative.QuickSight.DashboardLineChartLineStyle? lineStyle,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? lineVisibility,

            string? lineWidth)
        {
            LineInterpolation = lineInterpolation;
            LineStyle = lineStyle;
            LineVisibility = lineVisibility;
            LineWidth = lineWidth;
        }
    }
}
