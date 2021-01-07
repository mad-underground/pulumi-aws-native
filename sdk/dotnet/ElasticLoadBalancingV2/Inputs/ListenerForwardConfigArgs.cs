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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html
    /// </summary>
    public sealed class ListenerForwardConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html#cfn-elasticloadbalancingv2-listener-forwardconfig-targetgroupstickinessconfig
        /// </summary>
        [Input("TargetGroupStickinessConfig")]
        public Input<Inputs.ListenerTargetGroupStickinessConfigArgs>? TargetGroupStickinessConfig { get; set; }

        [Input("TargetGroups")]
        private InputList<Inputs.ListenerTargetGroupTupleArgs>? _TargetGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html#cfn-elasticloadbalancingv2-listener-forwardconfig-targetgroups
        /// </summary>
        public InputList<Inputs.ListenerTargetGroupTupleArgs> TargetGroups
        {
            get => _TargetGroups ?? (_TargetGroups = new InputList<Inputs.ListenerTargetGroupTupleArgs>());
            set => _TargetGroups = value;
        }

        public ListenerForwardConfigArgs()
        {
        }
    }
}