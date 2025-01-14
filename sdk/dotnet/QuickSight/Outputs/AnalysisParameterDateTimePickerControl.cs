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
    public sealed class AnalysisParameterDateTimePickerControl
    {
        public readonly Outputs.AnalysisDateTimePickerControlDisplayOptions? DisplayOptions;
        public readonly string ParameterControlId;
        public readonly string SourceParameterName;
        public readonly string Title;

        [OutputConstructor]
        private AnalysisParameterDateTimePickerControl(
            Outputs.AnalysisDateTimePickerControlDisplayOptions? displayOptions,

            string parameterControlId,

            string sourceParameterName,

            string title)
        {
            DisplayOptions = displayOptions;
            ParameterControlId = parameterControlId;
            SourceParameterName = sourceParameterName;
            Title = title;
        }
    }
}
