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
    public sealed class DashboardFilterSliderControl
    {
        public readonly Outputs.DashboardSliderControlDisplayOptions? DisplayOptions;
        public readonly string FilterControlId;
        public readonly double MaximumValue;
        public readonly double MinimumValue;
        public readonly string SourceFilterId;
        public readonly double StepSize;
        public readonly string Title;
        public readonly Pulumi.AwsNative.QuickSight.DashboardSheetControlSliderType? Type;

        [OutputConstructor]
        private DashboardFilterSliderControl(
            Outputs.DashboardSliderControlDisplayOptions? displayOptions,

            string filterControlId,

            double maximumValue,

            double minimumValue,

            string sourceFilterId,

            double stepSize,

            string title,

            Pulumi.AwsNative.QuickSight.DashboardSheetControlSliderType? type)
        {
            DisplayOptions = displayOptions;
            FilterControlId = filterControlId;
            MaximumValue = maximumValue;
            MinimumValue = minimumValue;
            SourceFilterId = sourceFilterId;
            StepSize = stepSize;
            Title = title;
            Type = type;
        }
    }
}
