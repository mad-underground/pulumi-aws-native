// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Lambda.Outputs
{

    [OutputType]
    public sealed class PermissionProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-action
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-eventsourcetoken
        /// </summary>
        public readonly string? EventSourceToken;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-functionname
        /// </summary>
        public readonly string FunctionName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-principal
        /// </summary>
        public readonly string Principal;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourceaccount
        /// </summary>
        public readonly string? SourceAccount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourcearn
        /// </summary>
        public readonly string? SourceArn;

        [OutputConstructor]
        private PermissionProperties(
            string Action,

            string? EventSourceToken,

            string FunctionName,

            string Principal,

            string? SourceAccount,

            string? SourceArn)
        {
            this.Action = Action;
            this.EventSourceToken = EventSourceToken;
            this.FunctionName = FunctionName;
            this.Principal = Principal;
            this.SourceAccount = SourceAccount;
            this.SourceArn = SourceArn;
        }
    }
}