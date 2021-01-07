// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Events.Outputs
{

    [OutputType]
    public sealed class RuleProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventbusname
        /// </summary>
        public readonly string? EventBusName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventpattern
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? EventPattern;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-rolearn
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-scheduleexpression
        /// </summary>
        public readonly string? ScheduleExpression;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-state
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-targets
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleTarget> Targets;

        [OutputConstructor]
        private RuleProperties(
            string? Description,

            string? EventBusName,

            Union<System.Text.Json.JsonElement, string>? EventPattern,

            string? Name,

            string? RoleArn,

            string? ScheduleExpression,

            string? State,

            ImmutableArray<Outputs.RuleTarget> Targets)
        {
            this.Description = Description;
            this.EventBusName = EventBusName;
            this.EventPattern = EventPattern;
            this.Name = Name;
            this.RoleArn = RoleArn;
            this.ScheduleExpression = ScheduleExpression;
            this.State = State;
            this.Targets = Targets;
        }
    }
}