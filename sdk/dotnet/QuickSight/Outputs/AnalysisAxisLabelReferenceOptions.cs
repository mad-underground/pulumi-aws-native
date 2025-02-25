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
    public sealed class AnalysisAxisLabelReferenceOptions
    {
        public readonly Outputs.AnalysisColumnIdentifier Column;
        public readonly string FieldId;

        [OutputConstructor]
        private AnalysisAxisLabelReferenceOptions(
            Outputs.AnalysisColumnIdentifier column,

            string fieldId)
        {
            Column = column;
            FieldId = fieldId;
        }
    }
}
