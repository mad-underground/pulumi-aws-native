// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.DLM.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html
    /// </summary>
    public sealed class LifecyclePolicyPolicyDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("Actions")]
        private InputList<Inputs.LifecyclePolicyActionArgs>? _Actions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-actions
        /// </summary>
        public InputList<Inputs.LifecyclePolicyActionArgs> Actions
        {
            get => _Actions ?? (_Actions = new InputList<Inputs.LifecyclePolicyActionArgs>());
            set => _Actions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-eventsource
        /// </summary>
        [Input("EventSource")]
        public Input<Inputs.LifecyclePolicyEventSourceArgs>? EventSource { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-parameters
        /// </summary>
        [Input("Parameters")]
        public Input<Inputs.LifecyclePolicyParametersArgs>? Parameters { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-policytype
        /// </summary>
        [Input("PolicyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("ResourceTypes")]
        private InputList<string>? _ResourceTypes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcetypes
        /// </summary>
        public InputList<string> ResourceTypes
        {
            get => _ResourceTypes ?? (_ResourceTypes = new InputList<string>());
            set => _ResourceTypes = value;
        }

        [Input("Schedules")]
        private InputList<Inputs.LifecyclePolicyScheduleArgs>? _Schedules;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-schedules
        /// </summary>
        public InputList<Inputs.LifecyclePolicyScheduleArgs> Schedules
        {
            get => _Schedules ?? (_Schedules = new InputList<Inputs.LifecyclePolicyScheduleArgs>());
            set => _Schedules = value;
        }

        [Input("TargetTags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _TargetTags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-targettags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> TargetTags
        {
            get => _TargetTags ?? (_TargetTags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _TargetTags = value;
        }

        public LifecyclePolicyPolicyDetailsArgs()
        {
        }
    }
}
