// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    public sealed class AssetModelMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mathematical expression that defines the metric aggregation function. You can specify up to 10 functions per expression.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("variables", required: true)]
        private InputList<Inputs.AssetModelExpressionVariableArgs>? _variables;

        /// <summary>
        /// The list of variables used in the expression.
        /// </summary>
        public InputList<Inputs.AssetModelExpressionVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.AssetModelExpressionVariableArgs>());
            set => _variables = value;
        }

        /// <summary>
        /// The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression
        /// </summary>
        [Input("window", required: true)]
        public Input<Inputs.AssetModelMetricWindowArgs> Window { get; set; } = null!;

        public AssetModelMetricArgs()
        {
        }
        public static new AssetModelMetricArgs Empty => new AssetModelMetricArgs();
    }
}
