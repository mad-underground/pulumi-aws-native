// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    public static class GetOrganizationConfigRule
    {
        /// <summary>
        /// Resource Type definition for AWS::Config::OrganizationConfigRule
        /// </summary>
        public static Task<GetOrganizationConfigRuleResult> InvokeAsync(GetOrganizationConfigRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationConfigRuleResult>("aws-native:configuration:getOrganizationConfigRule", args ?? new GetOrganizationConfigRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Config::OrganizationConfigRule
        /// </summary>
        public static Output<GetOrganizationConfigRuleResult> Invoke(GetOrganizationConfigRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationConfigRuleResult>("aws-native:configuration:getOrganizationConfigRule", args ?? new GetOrganizationConfigRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationConfigRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetOrganizationConfigRuleArgs()
        {
        }
        public static new GetOrganizationConfigRuleArgs Empty => new GetOrganizationConfigRuleArgs();
    }

    public sealed class GetOrganizationConfigRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetOrganizationConfigRuleInvokeArgs()
        {
        }
        public static new GetOrganizationConfigRuleInvokeArgs Empty => new GetOrganizationConfigRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationConfigRuleResult
    {
        public readonly ImmutableArray<string> ExcludedAccounts;
        public readonly string? Id;
        public readonly Outputs.OrganizationConfigRuleOrganizationCustomRuleMetadata? OrganizationCustomRuleMetadata;
        public readonly Outputs.OrganizationConfigRuleOrganizationManagedRuleMetadata? OrganizationManagedRuleMetadata;

        [OutputConstructor]
        private GetOrganizationConfigRuleResult(
            ImmutableArray<string> excludedAccounts,

            string? id,

            Outputs.OrganizationConfigRuleOrganizationCustomRuleMetadata? organizationCustomRuleMetadata,

            Outputs.OrganizationConfigRuleOrganizationManagedRuleMetadata? organizationManagedRuleMetadata)
        {
            ExcludedAccounts = excludedAccounts;
            Id = id;
            OrganizationCustomRuleMetadata = organizationCustomRuleMetadata;
            OrganizationManagedRuleMetadata = organizationManagedRuleMetadata;
        }
    }
}
