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
    public sealed class AnalysisFilterDateTimePickerControl
    {
        public readonly Outputs.AnalysisDateTimePickerControlDisplayOptions? DisplayOptions;
        public readonly string FilterControlId;
        public readonly string SourceFilterId;
        public readonly string Title;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisSheetControlDateTimePickerType? Type;

        [OutputConstructor]
        private AnalysisFilterDateTimePickerControl(
            Outputs.AnalysisDateTimePickerControlDisplayOptions? displayOptions,

            string filterControlId,

            string sourceFilterId,

            string title,

            Pulumi.AwsNative.QuickSight.AnalysisSheetControlDateTimePickerType? type)
        {
            DisplayOptions = displayOptions;
            FilterControlId = filterControlId;
            SourceFilterId = sourceFilterId;
            Title = title;
            Type = type;
        }
    }
}
