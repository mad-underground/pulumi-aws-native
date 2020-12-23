// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.GreengrassV2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html
    /// </summary>
    public sealed class ComponentVersionLambdaFunctionRecipeSourceArgs : Pulumi.ResourceArgs
    {
        [Input("ComponentDependencies")]
        private InputMap<Inputs.ComponentVersionComponentDependencyRequirementArgs>? _ComponentDependencies;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentdependencies
        /// </summary>
        public InputMap<Inputs.ComponentVersionComponentDependencyRequirementArgs> ComponentDependencies
        {
            get => _ComponentDependencies ?? (_ComponentDependencies = new InputMap<Inputs.ComponentVersionComponentDependencyRequirementArgs>());
            set => _ComponentDependencies = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentlambdaparameters
        /// </summary>
        [Input("ComponentLambdaParameters")]
        public Input<Inputs.ComponentVersionLambdaExecutionParametersArgs>? ComponentLambdaParameters { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentname
        /// </summary>
        [Input("ComponentName")]
        public Input<string>? ComponentName { get; set; }

        [Input("ComponentPlatforms")]
        private InputList<Inputs.ComponentVersionComponentPlatformArgs>? _ComponentPlatforms;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentplatforms
        /// </summary>
        public InputList<Inputs.ComponentVersionComponentPlatformArgs> ComponentPlatforms
        {
            get => _ComponentPlatforms ?? (_ComponentPlatforms = new InputList<Inputs.ComponentVersionComponentPlatformArgs>());
            set => _ComponentPlatforms = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentversion
        /// </summary>
        [Input("ComponentVersion")]
        public Input<string>? ComponentVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-lambdaarn
        /// </summary>
        [Input("LambdaArn")]
        public Input<string>? LambdaArn { get; set; }

        public ComponentVersionLambdaFunctionRecipeSourceArgs()
        {
        }
    }
}
