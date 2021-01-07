// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Outputs
{

    [OutputType]
    public sealed class DeviceProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-device
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Device;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-tags
        /// </summary>
        public readonly Pulumi.Cloudformation.Outputs.Tag? Tags;

        [OutputConstructor]
        private DeviceProperties(
            Union<System.Text.Json.JsonElement, string>? Device,

            Pulumi.Cloudformation.Outputs.Tag? Tags)
        {
            this.Device = Device;
            this.Tags = Tags;
        }
    }
}