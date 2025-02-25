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
    public sealed class AnalysisDataBarsOptions
    {
        public readonly string FieldId;
        public readonly string? NegativeColor;
        public readonly string? PositiveColor;

        [OutputConstructor]
        private AnalysisDataBarsOptions(
            string fieldId,

            string? negativeColor,

            string? positiveColor)
        {
            FieldId = fieldId;
            NegativeColor = negativeColor;
            PositiveColor = positiveColor;
        }
    }
}
