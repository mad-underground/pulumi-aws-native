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
    public sealed class AnalysisDefaultInteractiveLayoutConfiguration
    {
        public readonly Outputs.AnalysisDefaultFreeFormLayoutConfiguration? FreeForm;
        public readonly Outputs.AnalysisDefaultGridLayoutConfiguration? Grid;

        [OutputConstructor]
        private AnalysisDefaultInteractiveLayoutConfiguration(
            Outputs.AnalysisDefaultFreeFormLayoutConfiguration? freeForm,

            Outputs.AnalysisDefaultGridLayoutConfiguration? grid)
        {
            FreeForm = freeForm;
            Grid = grid;
        }
    }
}
