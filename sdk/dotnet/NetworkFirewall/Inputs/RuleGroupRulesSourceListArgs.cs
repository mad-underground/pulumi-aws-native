// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class RuleGroupRulesSourceListArgs : global::Pulumi.ResourceArgs
    {
        [Input("generatedRulesType", required: true)]
        public Input<Pulumi.AwsNative.NetworkFirewall.RuleGroupGeneratedRulesType> GeneratedRulesType { get; set; } = null!;

        [Input("targetTypes", required: true)]
        private InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTargetType>? _targetTypes;
        public InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTargetType> TargetTypes
        {
            get => _targetTypes ?? (_targetTypes = new InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTargetType>());
            set => _targetTypes = value;
        }

        [Input("targets", required: true)]
        private InputList<string>? _targets;
        public InputList<string> Targets
        {
            get => _targets ?? (_targets = new InputList<string>());
            set => _targets = value;
        }

        public RuleGroupRulesSourceListArgs()
        {
        }
        public static new RuleGroupRulesSourceListArgs Empty => new RuleGroupRulesSourceListArgs();
    }
}
