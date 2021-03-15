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
    public sealed class FlowTask
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-connectoroperator
        /// </summary>
        public readonly Outputs.FlowConnectorOperator? ConnectorOperator;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-destinationfield
        /// </summary>
        public readonly string? DestinationField;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-sourcefields
        /// </summary>
        public readonly ImmutableArray<string> SourceFields;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-taskproperties
        /// </summary>
        public readonly ImmutableArray<Outputs.FlowTaskPropertiesObject> TaskProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-tasktype
        /// </summary>
        public readonly string TaskType;

        [OutputConstructor]
        private FlowTask(
            Outputs.FlowConnectorOperator? connectorOperator,

            string? destinationField,

            ImmutableArray<string> sourceFields,

            ImmutableArray<Outputs.FlowTaskPropertiesObject> taskProperties,

            string taskType)
        {
            ConnectorOperator = connectorOperator;
            DestinationField = destinationField;
            SourceFields = sourceFields;
            TaskProperties = taskProperties;
            TaskType = taskType;
        }
    }
}
