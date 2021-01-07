// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CodeBuild.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html
    /// </summary>
    public sealed class ProjectProjectTriggersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-buildtype
        /// </summary>
        [Input("BuildType")]
        public Input<string>? BuildType { get; set; }

        [Input("FilterGroups")]
        private InputList<Inputs.ProjectFilterGroupArgs>? _FilterGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-filtergroups
        /// </summary>
        public InputList<Inputs.ProjectFilterGroupArgs> FilterGroups
        {
            get => _FilterGroups ?? (_FilterGroups = new InputList<Inputs.ProjectFilterGroupArgs>());
            set => _FilterGroups = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-webhook
        /// </summary>
        [Input("Webhook")]
        public Input<bool>? Webhook { get; set; }

        public ProjectProjectTriggersArgs()
        {
        }
    }
}
