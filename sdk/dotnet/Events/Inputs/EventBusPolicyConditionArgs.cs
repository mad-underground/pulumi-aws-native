// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Events.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html
    /// </summary>
    public sealed class EventBusPolicyConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-key
        /// </summary>
        [Input("Key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-type
        /// </summary>
        [Input("Type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-value
        /// </summary>
        [Input("Value")]
        public Input<string>? Value { get; set; }

        public EventBusPolicyConditionArgs()
        {
        }
    }
}