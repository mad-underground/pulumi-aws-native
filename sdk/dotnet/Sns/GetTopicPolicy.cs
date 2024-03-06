// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sns
{
    public static class GetTopicPolicy
    {
        /// <summary>
        /// The ``AWS::SNS::TopicPolicy`` resource associates SNS topics with a policy. For an example snippet, see [Declaring an policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sns-policy) in the *User Guide*.
        /// </summary>
        public static Task<GetTopicPolicyResult> InvokeAsync(GetTopicPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTopicPolicyResult>("aws-native:sns:getTopicPolicy", args ?? new GetTopicPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SNS::TopicPolicy`` resource associates SNS topics with a policy. For an example snippet, see [Declaring an policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sns-policy) in the *User Guide*.
        /// </summary>
        public static Output<GetTopicPolicyResult> Invoke(GetTopicPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicPolicyResult>("aws-native:sns:getTopicPolicy", args ?? new GetTopicPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTopicPolicyArgs()
        {
        }
        public static new GetTopicPolicyArgs Empty => new GetTopicPolicyArgs();
    }

    public sealed class GetTopicPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTopicPolicyInvokeArgs()
        {
        }
        public static new GetTopicPolicyInvokeArgs Empty => new GetTopicPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetTopicPolicyResult
    {
        public readonly string? Id;
        /// <summary>
        /// A policy document that contains permissions to add to the specified SNS topics.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::TopicPolicy` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? PolicyDocument;
        /// <summary>
        /// The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the ``Ref`` function to specify an ``AWS::SNS::Topic`` resource.
        /// </summary>
        public readonly ImmutableArray<string> Topics;

        [OutputConstructor]
        private GetTopicPolicyResult(
            string? id,

            object? policyDocument,

            ImmutableArray<string> topics)
        {
            Id = id;
            PolicyDocument = policyDocument;
            Topics = topics;
        }
    }
}
