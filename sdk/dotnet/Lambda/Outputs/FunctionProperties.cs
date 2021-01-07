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
    public sealed class FunctionProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-code
        /// </summary>
        public readonly Outputs.FunctionCode Code;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-codesigningconfigarn
        /// </summary>
        public readonly string? CodeSigningConfigArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-deadletterconfig
        /// </summary>
        public readonly Outputs.FunctionDeadLetterConfig? DeadLetterConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-environment
        /// </summary>
        public readonly Outputs.FunctionEnvironment? Environment;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-filesystemconfigs
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionFileSystemConfig> FileSystemConfigs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-functionname
        /// </summary>
        public readonly string? FunctionName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-handler
        /// </summary>
        public readonly string? Handler;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-imageconfig
        /// </summary>
        public readonly Outputs.FunctionImageConfig? ImageConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-kmskeyarn
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-layers
        /// </summary>
        public readonly ImmutableArray<string> Layers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-memorysize
        /// </summary>
        public readonly int? MemorySize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-packagetype
        /// </summary>
        public readonly string? PackageType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-reservedconcurrentexecutions
        /// </summary>
        public readonly int? ReservedConcurrentExecutions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-role
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-runtime
        /// </summary>
        public readonly string? Runtime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-timeout
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tracingconfig
        /// </summary>
        public readonly Outputs.FunctionTracingConfig? TracingConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-vpcconfig
        /// </summary>
        public readonly Outputs.FunctionVpcConfig? VpcConfig;

        [OutputConstructor]
        private FunctionProperties(
            Outputs.FunctionCode Code,

            string? CodeSigningConfigArn,

            Outputs.FunctionDeadLetterConfig? DeadLetterConfig,

            string? Description,

            Outputs.FunctionEnvironment? Environment,

            ImmutableArray<Outputs.FunctionFileSystemConfig> FileSystemConfigs,

            string? FunctionName,

            string? Handler,

            Outputs.FunctionImageConfig? ImageConfig,

            string? KmsKeyArn,

            ImmutableArray<string> Layers,

            int? MemorySize,

            string? PackageType,

            int? ReservedConcurrentExecutions,

            string Role,

            string? Runtime,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            int? Timeout,

            Outputs.FunctionTracingConfig? TracingConfig,

            Outputs.FunctionVpcConfig? VpcConfig)
        {
            this.Code = Code;
            this.CodeSigningConfigArn = CodeSigningConfigArn;
            this.DeadLetterConfig = DeadLetterConfig;
            this.Description = Description;
            this.Environment = Environment;
            this.FileSystemConfigs = FileSystemConfigs;
            this.FunctionName = FunctionName;
            this.Handler = Handler;
            this.ImageConfig = ImageConfig;
            this.KmsKeyArn = KmsKeyArn;
            this.Layers = Layers;
            this.MemorySize = MemorySize;
            this.PackageType = PackageType;
            this.ReservedConcurrentExecutions = ReservedConcurrentExecutions;
            this.Role = Role;
            this.Runtime = Runtime;
            this.Tags = Tags;
            this.Timeout = Timeout;
            this.TracingConfig = TracingConfig;
            this.VpcConfig = VpcConfig;
        }
    }
}