// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    /// <summary>
    /// The ``AccessLogSetting`` property type specifies settings for logging access in this stage.
    ///   ``AccessLogSetting`` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.
    /// </summary>
    [OutputType]
    public sealed class StageAccessLogSetting
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``. This parameter is required to enable access logging.
        /// </summary>
        public readonly string? DestinationArn;
        /// <summary>
        /// A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least ``$context.requestId``. This parameter is required to enable access logging.
        /// </summary>
        public readonly string? Format;

        [OutputConstructor]
        private StageAccessLogSetting(
            string? destinationArn,

            string? format)
        {
            DestinationArn = destinationArn;
            Format = format;
        }
    }
}
