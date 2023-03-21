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
    public sealed class AnalysisFieldTooltipItem
    {
        public readonly string FieldId;
        public readonly string? Label;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? Visibility;

        [OutputConstructor]
        private AnalysisFieldTooltipItem(
            string fieldId,

            string? label,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? visibility)
        {
            FieldId = fieldId;
            Label = label;
            Visibility = visibility;
        }
    }
}
