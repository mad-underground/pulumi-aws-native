// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFormation.Outputs
{

    [OutputType]
    public sealed class MacroProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-functionname
        /// </summary>
        public readonly string FunctionName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-loggroupname
        /// </summary>
        public readonly string? LogGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-logrolearn
        /// </summary>
        public readonly string? LogRoleARN;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private MacroProperties(
            string? Description,

            string FunctionName,

            string? LogGroupName,

            string? LogRoleARN,

            string Name)
        {
            this.Description = Description;
            this.FunctionName = FunctionName;
            this.LogGroupName = LogGroupName;
            this.LogRoleARN = LogRoleARN;
            this.Name = Name;
        }
    }
}