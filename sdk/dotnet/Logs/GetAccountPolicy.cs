// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    public static class GetAccountPolicy
    {
        /// <summary>
        /// The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.
        /// </summary>
        public static Task<GetAccountPolicyResult> InvokeAsync(GetAccountPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountPolicyResult>("aws-native:logs:getAccountPolicy", args ?? new GetAccountPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.
        /// </summary>
        public static Output<GetAccountPolicyResult> Invoke(GetAccountPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountPolicyResult>("aws-native:logs:getAccountPolicy", args ?? new GetAccountPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// User account id
        /// </summary>
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        /// <summary>
        /// The name of the account policy
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        /// <summary>
        /// Type of the policy.
        /// </summary>
        [Input("policyType", required: true)]
        public Pulumi.AwsNative.Logs.AccountPolicyPolicyType PolicyType { get; set; }

        public GetAccountPolicyArgs()
        {
        }
        public static new GetAccountPolicyArgs Empty => new GetAccountPolicyArgs();
    }

    public sealed class GetAccountPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// User account id
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The name of the account policy
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// Type of the policy.
        /// </summary>
        [Input("policyType", required: true)]
        public Input<Pulumi.AwsNative.Logs.AccountPolicyPolicyType> PolicyType { get; set; } = null!;

        public GetAccountPolicyInvokeArgs()
        {
        }
        public static new GetAccountPolicyInvokeArgs Empty => new GetAccountPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountPolicyResult
    {
        /// <summary>
        /// User account id
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// The body of the policy document you want to use for this topic.
        /// 
        /// You can only add one policy per PolicyType.
        /// 
        /// The policy must be in JSON string format.
        /// 
        /// Length Constraints: Maximum length of 30720
        /// </summary>
        public readonly string? PolicyDocument;
        /// <summary>
        /// Scope for policy application
        /// </summary>
        public readonly Pulumi.AwsNative.Logs.AccountPolicyScope? Scope;

        [OutputConstructor]
        private GetAccountPolicyResult(
            string? accountId,

            string? policyDocument,

            Pulumi.AwsNative.Logs.AccountPolicyScope? scope)
        {
            AccountId = accountId;
            PolicyDocument = policyDocument;
            Scope = scope;
        }
    }
}
