// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    [OutputType]
    public sealed class OrganizationConfigRuleOrganizationCustomRuleMetadata
    {
        public readonly string? Description;
        public readonly string? InputParameters;
        public readonly string LambdaFunctionArn;
        public readonly string? MaximumExecutionFrequency;
        public readonly ImmutableArray<string> OrganizationConfigRuleTriggerTypes;
        public readonly string? ResourceIdScope;
        public readonly ImmutableArray<string> ResourceTypesScope;
        public readonly string? TagKeyScope;
        public readonly string? TagValueScope;

        [OutputConstructor]
        private OrganizationConfigRuleOrganizationCustomRuleMetadata(
            string? description,

            string? inputParameters,

            string lambdaFunctionArn,

            string? maximumExecutionFrequency,

            ImmutableArray<string> organizationConfigRuleTriggerTypes,

            string? resourceIdScope,

            ImmutableArray<string> resourceTypesScope,

            string? tagKeyScope,

            string? tagValueScope)
        {
            Description = description;
            InputParameters = inputParameters;
            LambdaFunctionArn = lambdaFunctionArn;
            MaximumExecutionFrequency = maximumExecutionFrequency;
            OrganizationConfigRuleTriggerTypes = organizationConfigRuleTriggerTypes;
            ResourceIdScope = resourceIdScope;
            ResourceTypesScope = resourceTypesScope;
            TagKeyScope = tagKeyScope;
            TagValueScope = tagValueScope;
        }
    }
}
