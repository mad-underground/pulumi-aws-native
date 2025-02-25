// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Inputs
{

    public sealed class EventTypeEventVariableArgs : global::Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The time when the event type was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        [Input("dataSource")]
        public Input<Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataSource>? DataSource { get; set; }

        [Input("dataType")]
        public Input<Pulumi.AwsNative.FraudDetector.EventTypeEventVariableDataType>? DataType { get; set; }

        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// The description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inline")]
        public Input<bool>? Inline { get; set; }

        /// <summary>
        /// The time when the event type was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.EventTypeTagArgs>? _tags;

        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        public InputList<Inputs.EventTypeTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.EventTypeTagArgs>());
            set => _tags = value;
        }

        [Input("variableType")]
        public Input<Pulumi.AwsNative.FraudDetector.EventTypeEventVariableVariableType>? VariableType { get; set; }

        public EventTypeEventVariableArgs()
        {
        }
        public static new EventTypeEventVariableArgs Empty => new EventTypeEventVariableArgs();
    }
}
