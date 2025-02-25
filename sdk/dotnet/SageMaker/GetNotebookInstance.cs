// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetNotebookInstance
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::NotebookInstance
        /// </summary>
        public static Task<GetNotebookInstanceResult> InvokeAsync(GetNotebookInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotebookInstanceResult>("aws-native:sagemaker:getNotebookInstance", args ?? new GetNotebookInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::NotebookInstance
        /// </summary>
        public static Output<GetNotebookInstanceResult> Invoke(GetNotebookInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotebookInstanceResult>("aws-native:sagemaker:getNotebookInstance", args ?? new GetNotebookInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotebookInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNotebookInstanceArgs()
        {
        }
        public static new GetNotebookInstanceArgs Empty => new GetNotebookInstanceArgs();
    }

    public sealed class GetNotebookInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNotebookInstanceInvokeArgs()
        {
        }
        public static new GetNotebookInstanceInvokeArgs Empty => new GetNotebookInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotebookInstanceResult
    {
        public readonly ImmutableArray<string> AcceleratorTypes;
        public readonly ImmutableArray<string> AdditionalCodeRepositories;
        public readonly string? DefaultCodeRepository;
        public readonly string? Id;
        public readonly Outputs.NotebookInstanceInstanceMetadataServiceConfiguration? InstanceMetadataServiceConfiguration;
        public readonly string? InstanceType;
        public readonly string? LifecycleConfigName;
        public readonly string? RoleArn;
        public readonly string? RootAccess;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly int? VolumeSizeInGb;

        [OutputConstructor]
        private GetNotebookInstanceResult(
            ImmutableArray<string> acceleratorTypes,

            ImmutableArray<string> additionalCodeRepositories,

            string? defaultCodeRepository,

            string? id,

            Outputs.NotebookInstanceInstanceMetadataServiceConfiguration? instanceMetadataServiceConfiguration,

            string? instanceType,

            string? lifecycleConfigName,

            string? roleArn,

            string? rootAccess,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            int? volumeSizeInGb)
        {
            AcceleratorTypes = acceleratorTypes;
            AdditionalCodeRepositories = additionalCodeRepositories;
            DefaultCodeRepository = defaultCodeRepository;
            Id = id;
            InstanceMetadataServiceConfiguration = instanceMetadataServiceConfiguration;
            InstanceType = instanceType;
            LifecycleConfigName = lifecycleConfigName;
            RoleArn = roleArn;
            RootAccess = rootAccess;
            Tags = tags;
            VolumeSizeInGb = volumeSizeInGb;
        }
    }
}
