// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAF
{
    public static class GetWebACL
    {
        /// <summary>
        /// Resource Type definition for AWS::WAF::WebACL
        /// </summary>
        public static Task<GetWebACLResult> InvokeAsync(GetWebACLArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebACLResult>("aws-native:waf:getWebACL", args ?? new GetWebACLArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WAF::WebACL
        /// </summary>
        public static Output<GetWebACLResult> Invoke(GetWebACLInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebACLResult>("aws-native:waf:getWebACL", args ?? new GetWebACLInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebACLArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWebACLArgs()
        {
        }
        public static new GetWebACLArgs Empty => new GetWebACLArgs();
    }

    public sealed class GetWebACLInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWebACLInvokeArgs()
        {
        }
        public static new GetWebACLInvokeArgs Empty => new GetWebACLInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebACLResult
    {
        public readonly Outputs.WebACLWafAction? DefaultAction;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.WebACLActivatedRule> Rules;

        [OutputConstructor]
        private GetWebACLResult(
            Outputs.WebACLWafAction? defaultAction,

            string? id,

            ImmutableArray<Outputs.WebACLActivatedRule> rules)
        {
            DefaultAction = defaultAction;
            Id = id;
            Rules = rules;
        }
    }
}
