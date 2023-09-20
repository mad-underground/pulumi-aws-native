// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WafRegional
{
    public static class GetWebAcl
    {
        /// <summary>
        /// Resource Type definition for AWS::WAFRegional::WebACL
        /// </summary>
        public static Task<GetWebAclResult> InvokeAsync(GetWebAclArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAclResult>("aws-native:wafregional:getWebAcl", args ?? new GetWebAclArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WAFRegional::WebACL
        /// </summary>
        public static Output<GetWebAclResult> Invoke(GetWebAclInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAclResult>("aws-native:wafregional:getWebAcl", args ?? new GetWebAclInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAclArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWebAclArgs()
        {
        }
        public static new GetWebAclArgs Empty => new GetWebAclArgs();
    }

    public sealed class GetWebAclInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWebAclInvokeArgs()
        {
        }
        public static new GetWebAclInvokeArgs Empty => new GetWebAclInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAclResult
    {
        public readonly Outputs.WebAclAction? DefaultAction;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.WebAclRule> Rules;

        [OutputConstructor]
        private GetWebAclResult(
            Outputs.WebAclAction? defaultAction,

            string? id,

            ImmutableArray<Outputs.WebAclRule> rules)
        {
            DefaultAction = defaultAction;
            Id = id;
            Rules = rules;
        }
    }
}