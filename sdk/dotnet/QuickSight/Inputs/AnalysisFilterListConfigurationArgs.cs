// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisFilterListConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryValues")]
        private InputList<string>? _categoryValues;
        public InputList<string> CategoryValues
        {
            get => _categoryValues ?? (_categoryValues = new InputList<string>());
            set => _categoryValues = value;
        }

        [Input("matchOperator", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisCategoryFilterMatchOperator> MatchOperator { get; set; } = null!;

        [Input("nullOption")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisFilterNullOption>? NullOption { get; set; }

        [Input("selectAllOptions")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisCategoryFilterSelectAllOptions>? SelectAllOptions { get; set; }

        public AnalysisFilterListConfigurationArgs()
        {
        }
        public static new AnalysisFilterListConfigurationArgs Empty => new AnalysisFilterListConfigurationArgs();
    }
}
