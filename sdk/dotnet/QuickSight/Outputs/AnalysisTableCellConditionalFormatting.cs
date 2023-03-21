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
    public sealed class AnalysisTableCellConditionalFormatting
    {
        public readonly string FieldId;
        public readonly Outputs.AnalysisTextConditionalFormat? TextFormat;

        [OutputConstructor]
        private AnalysisTableCellConditionalFormatting(
            string fieldId,

            Outputs.AnalysisTextConditionalFormat? textFormat)
        {
            FieldId = fieldId;
            TextFormat = textFormat;
        }
    }
}
