// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html
    /// </summary>
    [OutputType]
    public sealed class ComponentVersionLambdaFunctionRecipeSource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentdependencies
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ComponentVersionComponentDependencyRequirement>? ComponentDependencies;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentlambdaparameters
        /// </summary>
        public readonly Outputs.ComponentVersionLambdaExecutionParameters? ComponentLambdaParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentname
        /// </summary>
        public readonly string? ComponentName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentplatforms
        /// </summary>
        public readonly ImmutableArray<Outputs.ComponentVersionComponentPlatform> ComponentPlatforms;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentversion
        /// </summary>
        public readonly string? ComponentVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-lambdaarn
        /// </summary>
        public readonly string? LambdaArn;

        [OutputConstructor]
        private ComponentVersionLambdaFunctionRecipeSource(
            ImmutableDictionary<string, Outputs.ComponentVersionComponentDependencyRequirement>? componentDependencies,

            Outputs.ComponentVersionLambdaExecutionParameters? componentLambdaParameters,

            string? componentName,

            ImmutableArray<Outputs.ComponentVersionComponentPlatform> componentPlatforms,

            string? componentVersion,

            string? lambdaArn)
        {
            ComponentDependencies = componentDependencies;
            ComponentLambdaParameters = componentLambdaParameters;
            ComponentName = componentName;
            ComponentPlatforms = componentPlatforms;
            ComponentVersion = componentVersion;
            LambdaArn = lambdaArn;
        }
    }
}
