// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowDestinationConnectorProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-eventbridge
        /// </summary>
        public readonly Outputs.FlowEventBridgeDestinationProperties? EventBridge;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-redshift
        /// </summary>
        public readonly Outputs.FlowRedshiftDestinationProperties? Redshift;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-s3
        /// </summary>
        public readonly Outputs.FlowS3DestinationProperties? S3;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-salesforce
        /// </summary>
        public readonly Outputs.FlowSalesforceDestinationProperties? Salesforce;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-snowflake
        /// </summary>
        public readonly Outputs.FlowSnowflakeDestinationProperties? Snowflake;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-upsolver
        /// </summary>
        public readonly Outputs.FlowUpsolverDestinationProperties? Upsolver;

        [OutputConstructor]
        private FlowDestinationConnectorProperties(
            Outputs.FlowEventBridgeDestinationProperties? EventBridge,

            Outputs.FlowRedshiftDestinationProperties? Redshift,

            Outputs.FlowS3DestinationProperties? S3,

            Outputs.FlowSalesforceDestinationProperties? Salesforce,

            Outputs.FlowSnowflakeDestinationProperties? Snowflake,

            Outputs.FlowUpsolverDestinationProperties? Upsolver)
        {
            this.EventBridge = EventBridge;
            this.Redshift = Redshift;
            this.S3 = S3;
            this.Salesforce = Salesforce;
            this.Snowflake = Snowflake;
            this.Upsolver = Upsolver;
        }
    }
}