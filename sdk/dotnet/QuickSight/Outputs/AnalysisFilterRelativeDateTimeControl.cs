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
    public sealed class AnalysisFilterRelativeDateTimeControl
    {
        public readonly Outputs.AnalysisRelativeDateTimeControlDisplayOptions? DisplayOptions;
        public readonly string FilterControlId;
        public readonly string SourceFilterId;
        public readonly string Title;

        [OutputConstructor]
        private AnalysisFilterRelativeDateTimeControl(
            Outputs.AnalysisRelativeDateTimeControlDisplayOptions? displayOptions,

            string filterControlId,

            string sourceFilterId,

            string title)
        {
            DisplayOptions = displayOptions;
            FilterControlId = filterControlId;
            SourceFilterId = sourceFilterId;
            Title = title;
        }
    }
}