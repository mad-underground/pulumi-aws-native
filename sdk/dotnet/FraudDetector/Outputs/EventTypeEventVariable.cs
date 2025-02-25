// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Outputs
{

    [OutputType]
    public sealed class EventTypeEventVariable
    {
        public readonly string? Arn;
        /// <summary>
        /// The time when the event type was created.
        /// </summary>
        public readonly string? CreatedTime;
        public readonly Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataSource? DataSource;
        public readonly Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataType? DataType;
        public readonly string? DefaultValue;
        /// <summary>
        /// The description.
        /// </summary>
        public readonly string? Description;
        public readonly bool? Inline;
        /// <summary>
        /// The time when the event type was last updated.
        /// </summary>
        public readonly string? LastUpdatedTime;
        public readonly string? Name;
        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventTypeTag> Tags;
        public readonly Pulumi.AwsNative.FraudDetector.EventTypeEventVariableVariableType? VariableType;

        [OutputConstructor]
        private EventTypeEventVariable(
            string? arn,

            string? createdTime,

            Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataSource? dataSource,

            Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataType? dataType,

            string? defaultValue,

            string? description,

            bool? inline,

            string? lastUpdatedTime,

            string? name,

            ImmutableArray<Outputs.EventTypeTag> tags,

            Pulumi.AwsNative.FraudDetector.EventTypeEventVariableVariableType? variableType)
        {
            Arn = arn;
            CreatedTime = createdTime;
            DataSource = dataSource;
            DataType = dataType;
            DefaultValue = defaultValue;
            Description = description;
            Inline = inline;
            LastUpdatedTime = lastUpdatedTime;
            Name = name;
            Tags = tags;
            VariableType = variableType;
        }
    }
}
