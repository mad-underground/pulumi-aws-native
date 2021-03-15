// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowSourceFlowConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-connectorprofilename
        /// </summary>
        public readonly string? ConnectorProfileName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-connectortype
        /// </summary>
        public readonly string ConnectorType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-incrementalpullconfig
        /// </summary>
        public readonly Outputs.FlowIncrementalPullConfig? IncrementalPullConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-sourceconnectorproperties
        /// </summary>
        public readonly Outputs.FlowSourceConnectorProperties SourceConnectorProperties;

        [OutputConstructor]
        private FlowSourceFlowConfig(
            string? connectorProfileName,

            string connectorType,

            Outputs.FlowIncrementalPullConfig? incrementalPullConfig,

            Outputs.FlowSourceConnectorProperties sourceConnectorProperties)
        {
            ConnectorProfileName = connectorProfileName;
            ConnectorType = connectorType;
            IncrementalPullConfig = incrementalPullConfig;
            SourceConnectorProperties = sourceConnectorProperties;
        }
    }
}
