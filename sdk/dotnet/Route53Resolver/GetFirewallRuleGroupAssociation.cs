// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver
{
    public static class GetFirewallRuleGroupAssociation
    {
        /// <summary>
        /// Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.
        /// </summary>
        public static Task<GetFirewallRuleGroupAssociationResult> InvokeAsync(GetFirewallRuleGroupAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFirewallRuleGroupAssociationResult>("aws-native:route53resolver:getFirewallRuleGroupAssociation", args ?? new GetFirewallRuleGroupAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.
        /// </summary>
        public static Output<GetFirewallRuleGroupAssociationResult> Invoke(GetFirewallRuleGroupAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFirewallRuleGroupAssociationResult>("aws-native:route53resolver:getFirewallRuleGroupAssociation", args ?? new GetFirewallRuleGroupAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallRuleGroupAssociationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFirewallRuleGroupAssociationArgs()
        {
        }
    }

    public sealed class GetFirewallRuleGroupAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFirewallRuleGroupAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFirewallRuleGroupAssociationResult
    {
        /// <summary>
        /// Arn
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Rfc3339TimeString
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The id of the creator request.
        /// </summary>
        public readonly string? CreatorRequestId;
        /// <summary>
        /// Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// ServicePrincipal
        /// </summary>
        public readonly string? ManagedOwnerName;
        /// <summary>
        /// Rfc3339TimeString
        /// </summary>
        public readonly string? ModificationTime;
        /// <summary>
        /// MutationProtectionStatus
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupAssociationMutationProtection? MutationProtection;
        /// <summary>
        /// FirewallRuleGroupAssociationName
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Priority
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupAssociationStatus? Status;
        /// <summary>
        /// FirewallDomainListAssociationStatus
        /// </summary>
        public readonly string? StatusMessage;
        /// <summary>
        /// Tags
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallRuleGroupAssociationTag> Tags;

        [OutputConstructor]
        private GetFirewallRuleGroupAssociationResult(
            string? arn,

            string? creationTime,

            string? creatorRequestId,

            string? id,

            string? managedOwnerName,

            string? modificationTime,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupAssociationMutationProtection? mutationProtection,

            string? name,

            int? priority,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupAssociationStatus? status,

            string? statusMessage,

            ImmutableArray<Outputs.FirewallRuleGroupAssociationTag> tags)
        {
            Arn = arn;
            CreationTime = creationTime;
            CreatorRequestId = creatorRequestId;
            Id = id;
            ManagedOwnerName = managedOwnerName;
            ModificationTime = modificationTime;
            MutationProtection = mutationProtection;
            Name = name;
            Priority = priority;
            Status = status;
            StatusMessage = statusMessage;
            Tags = tags;
        }
    }
}