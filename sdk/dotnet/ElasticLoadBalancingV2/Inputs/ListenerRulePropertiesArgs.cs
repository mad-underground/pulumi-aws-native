// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElasticLoadBalancingV2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
    /// </summary>
    public sealed class ListenerRulePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("Actions", required: true)]
        private InputList<Inputs.ListenerRuleActionArgs>? _Actions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-actions
        /// </summary>
        public InputList<Inputs.ListenerRuleActionArgs> Actions
        {
            get => _Actions ?? (_Actions = new InputList<Inputs.ListenerRuleActionArgs>());
            set => _Actions = value;
        }

        [Input("Conditions", required: true)]
        private InputList<Inputs.ListenerRuleRuleConditionArgs>? _Conditions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-conditions
        /// </summary>
        public InputList<Inputs.ListenerRuleRuleConditionArgs> Conditions
        {
            get => _Conditions ?? (_Conditions = new InputList<Inputs.ListenerRuleRuleConditionArgs>());
            set => _Conditions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-listenerarn
        /// </summary>
        [Input("ListenerArn", required: true)]
        public Input<string> ListenerArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-priority
        /// </summary>
        [Input("Priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        public ListenerRulePropertiesArgs()
        {
        }
    }
}