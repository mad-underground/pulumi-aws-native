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
    public sealed class DashboardReferenceLineDataConfiguration
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardAxisBinding? AxisBinding;
        public readonly Outputs.DashboardReferenceLineDynamicDataConfiguration? DynamicConfiguration;
        public readonly Pulumi.AwsNative.QuickSight.DashboardReferenceLineSeriesType? SeriesType;
        public readonly Outputs.DashboardReferenceLineStaticDataConfiguration? StaticConfiguration;

        [OutputConstructor]
        private DashboardReferenceLineDataConfiguration(
            Pulumi.AwsNative.QuickSight.DashboardAxisBinding? axisBinding,

            Outputs.DashboardReferenceLineDynamicDataConfiguration? dynamicConfiguration,

            Pulumi.AwsNative.QuickSight.DashboardReferenceLineSeriesType? seriesType,

            Outputs.DashboardReferenceLineStaticDataConfiguration? staticConfiguration)
        {
            AxisBinding = axisBinding;
            DynamicConfiguration = dynamicConfiguration;
            SeriesType = seriesType;
            StaticConfiguration = staticConfiguration;
        }
    }
}
