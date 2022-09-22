// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryControl
{
    public static class GetSafetyRule
    {
        /// <summary>
        /// Resource schema for AWS Route53 Recovery Control basic constructs and validation rules.
        /// </summary>
        public static Task<GetSafetyRuleResult> InvokeAsync(GetSafetyRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSafetyRuleResult>("aws-native:route53recoverycontrol:getSafetyRule", args ?? new GetSafetyRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS Route53 Recovery Control basic constructs and validation rules.
        /// </summary>
        public static Output<GetSafetyRuleResult> Invoke(GetSafetyRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSafetyRuleResult>("aws-native:route53recoverycontrol:getSafetyRule", args ?? new GetSafetyRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSafetyRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the safety rule.
        /// </summary>
        [Input("safetyRuleArn", required: true)]
        public string SafetyRuleArn { get; set; } = null!;

        public GetSafetyRuleArgs()
        {
        }
        public static new GetSafetyRuleArgs Empty => new GetSafetyRuleArgs();
    }

    public sealed class GetSafetyRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the safety rule.
        /// </summary>
        [Input("safetyRuleArn", required: true)]
        public Input<string> SafetyRuleArn { get; set; } = null!;

        public GetSafetyRuleInvokeArgs()
        {
        }
        public static new GetSafetyRuleInvokeArgs Empty => new GetSafetyRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetSafetyRuleResult
    {
        public readonly Outputs.SafetyRuleAssertionRule? AssertionRule;
        public readonly Outputs.SafetyRuleGatingRule? GatingRule;
        public readonly string? Name;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the safety rule.
        /// </summary>
        public readonly string? SafetyRuleArn;
        /// <summary>
        /// The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53RecoveryControl.SafetyRuleStatus? Status;

        [OutputConstructor]
        private GetSafetyRuleResult(
            Outputs.SafetyRuleAssertionRule? assertionRule,

            Outputs.SafetyRuleGatingRule? gatingRule,

            string? name,

            string? safetyRuleArn,

            Pulumi.AwsNative.Route53RecoveryControl.SafetyRuleStatus? status)
        {
            AssertionRule = assertionRule;
            GatingRule = gatingRule;
            Name = name;
            SafetyRuleArn = safetyRuleArn;
            Status = status;
        }
    }
}
