// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Greengrass.Outputs
{

    [OutputType]
    public sealed class FunctionDefinitionVersionEnvironment
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-accesssysfs
        /// </summary>
        public readonly bool? AccessSysfs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-execution
        /// </summary>
        public readonly Outputs.FunctionDefinitionVersionExecution? Execution;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-resourceaccesspolicies
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionDefinitionVersionResourceAccessPolicy> ResourceAccessPolicies;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-variables
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Variables;

        [OutputConstructor]
        private FunctionDefinitionVersionEnvironment(
            bool? AccessSysfs,

            Outputs.FunctionDefinitionVersionExecution? Execution,

            ImmutableArray<Outputs.FunctionDefinitionVersionResourceAccessPolicy> ResourceAccessPolicies,

            Union<System.Text.Json.JsonElement, string>? Variables)
        {
            this.AccessSysfs = AccessSysfs;
            this.Execution = Execution;
            this.ResourceAccessPolicies = ResourceAccessPolicies;
            this.Variables = Variables;
        }
    }
}