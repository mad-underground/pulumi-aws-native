// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2
{
    /// <summary>
    /// Resource Type definition for AWS::ElasticLoadBalancingV2::ListenerRule
    /// </summary>
    [AwsNativeResourceType("aws-native:elasticloadbalancingv2:ListenerRule")]
    public partial class ListenerRule : global::Pulumi.CustomResource
    {
        [Output("actions")]
        public Output<ImmutableArray<Outputs.ListenerRuleAction>> Actions { get; private set; } = null!;

        [Output("conditions")]
        public Output<ImmutableArray<Outputs.ListenerRuleRuleCondition>> Conditions { get; private set; } = null!;

        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        [Output("listenerArn")]
        public Output<string> ListenerArn { get; private set; } = null!;

        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        [Output("ruleArn")]
        public Output<string> RuleArn { get; private set; } = null!;


        /// <summary>
        /// Create a ListenerRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ListenerRule(string name, ListenerRuleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:elasticloadbalancingv2:ListenerRule", name, args ?? new ListenerRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ListenerRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:elasticloadbalancingv2:ListenerRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ListenerRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ListenerRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ListenerRule(name, id, options);
        }
    }

    public sealed class ListenerRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.ListenerRuleActionArgs>? _actions;
        public InputList<Inputs.ListenerRuleActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.ListenerRuleActionArgs>());
            set => _actions = value;
        }

        [Input("conditions", required: true)]
        private InputList<Inputs.ListenerRuleRuleConditionArgs>? _conditions;
        public InputList<Inputs.ListenerRuleRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ListenerRuleRuleConditionArgs>());
            set => _conditions = value;
        }

        [Input("listenerArn", required: true)]
        public Input<string> ListenerArn { get; set; } = null!;

        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        public ListenerRuleArgs()
        {
        }
        public static new ListenerRuleArgs Empty => new ListenerRuleArgs();
    }
}
