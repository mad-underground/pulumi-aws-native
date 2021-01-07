// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-aggregationconfig.html
    /// </summary>
    public sealed class FlowAggregationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-aggregationconfig.html#cfn-appflow-flow-aggregationconfig-aggregationtype
        /// </summary>
        [Input("AggregationType")]
        public Input<string>? AggregationType { get; set; }

        public FlowAggregationConfigArgs()
        {
        }
    }
}