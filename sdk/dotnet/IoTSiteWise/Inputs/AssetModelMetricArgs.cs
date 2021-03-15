// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html
    /// </summary>
    public sealed class AssetModelMetricArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html#cfn-iotsitewise-assetmodel-metric-expression
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("variables", required: true)]
        private InputList<Inputs.AssetModelExpressionVariableArgs>? _variables;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html#cfn-iotsitewise-assetmodel-metric-variables
        /// </summary>
        public InputList<Inputs.AssetModelExpressionVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.AssetModelExpressionVariableArgs>());
            set => _variables = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html#cfn-iotsitewise-assetmodel-metric-window
        /// </summary>
        [Input("window", required: true)]
        public Input<Inputs.AssetModelMetricWindowArgs> Window { get; set; } = null!;

        public AssetModelMetricArgs()
        {
        }
    }
}
