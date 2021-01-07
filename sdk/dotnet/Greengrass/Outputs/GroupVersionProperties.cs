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
    public sealed class GroupVersionProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-connectordefinitionversionarn
        /// </summary>
        public readonly string? ConnectorDefinitionVersionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-coredefinitionversionarn
        /// </summary>
        public readonly string? CoreDefinitionVersionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-devicedefinitionversionarn
        /// </summary>
        public readonly string? DeviceDefinitionVersionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-functiondefinitionversionarn
        /// </summary>
        public readonly string? FunctionDefinitionVersionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-groupid
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-loggerdefinitionversionarn
        /// </summary>
        public readonly string? LoggerDefinitionVersionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-resourcedefinitionversionarn
        /// </summary>
        public readonly string? ResourceDefinitionVersionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-subscriptiondefinitionversionarn
        /// </summary>
        public readonly string? SubscriptionDefinitionVersionArn;

        [OutputConstructor]
        private GroupVersionProperties(
            string? ConnectorDefinitionVersionArn,

            string? CoreDefinitionVersionArn,

            string? DeviceDefinitionVersionArn,

            string? FunctionDefinitionVersionArn,

            string GroupId,

            string? LoggerDefinitionVersionArn,

            string? ResourceDefinitionVersionArn,

            string? SubscriptionDefinitionVersionArn)
        {
            this.ConnectorDefinitionVersionArn = ConnectorDefinitionVersionArn;
            this.CoreDefinitionVersionArn = CoreDefinitionVersionArn;
            this.DeviceDefinitionVersionArn = DeviceDefinitionVersionArn;
            this.FunctionDefinitionVersionArn = FunctionDefinitionVersionArn;
            this.GroupId = GroupId;
            this.LoggerDefinitionVersionArn = LoggerDefinitionVersionArn;
            this.ResourceDefinitionVersionArn = ResourceDefinitionVersionArn;
            this.SubscriptionDefinitionVersionArn = SubscriptionDefinitionVersionArn;
        }
    }
}