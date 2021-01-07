// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Glue.Outputs
{

    [OutputType]
    public sealed class WorkflowProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-defaultrunproperties
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? DefaultRunProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-tags
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Tags;

        [OutputConstructor]
        private WorkflowProperties(
            Union<System.Text.Json.JsonElement, string>? DefaultRunProperties,

            string? Description,

            string? Name,

            Union<System.Text.Json.JsonElement, string>? Tags)
        {
            this.DefaultRunProperties = DefaultRunProperties;
            this.Description = Description;
            this.Name = Name;
            this.Tags = Tags;
        }
    }
}