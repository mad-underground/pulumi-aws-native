// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoT1Click.Outputs
{

    [OutputType]
    public sealed class ProjectProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-placementtemplate
        /// </summary>
        public readonly Outputs.ProjectPlacementTemplate PlacementTemplate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-projectname
        /// </summary>
        public readonly string? ProjectName;

        [OutputConstructor]
        private ProjectProperties(
            string? Description,

            Outputs.ProjectPlacementTemplate PlacementTemplate,

            string? ProjectName)
        {
            this.Description = Description;
            this.PlacementTemplate = PlacementTemplate;
            this.ProjectName = ProjectName;
        }
    }
}