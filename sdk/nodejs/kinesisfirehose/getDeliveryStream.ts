// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::KinesisFirehose::DeliveryStream
 */
export function getDeliveryStream(args: GetDeliveryStreamArgs, opts?: pulumi.InvokeOptions): Promise<GetDeliveryStreamResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:kinesisfirehose:getDeliveryStream", {
        "deliveryStreamName": args.deliveryStreamName,
    }, opts);
}

export interface GetDeliveryStreamArgs {
    deliveryStreamName: string;
}

export interface GetDeliveryStreamResult {
    readonly amazonopensearchserviceDestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamAmazonopensearchserviceDestinationConfiguration;
    readonly arn?: string;
    readonly deliveryStreamEncryptionConfigurationInput?: outputs.kinesisfirehose.DeliveryStreamEncryptionConfigurationInput;
    readonly elasticsearchDestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamElasticsearchDestinationConfiguration;
    readonly extendedS3DestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamExtendedS3DestinationConfiguration;
    readonly httpEndpointDestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamHttpEndpointDestinationConfiguration;
    readonly redshiftDestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamRedshiftDestinationConfiguration;
    readonly s3DestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamS3DestinationConfiguration;
    readonly splunkDestinationConfiguration?: outputs.kinesisfirehose.DeliveryStreamSplunkDestinationConfiguration;
    readonly tags?: outputs.kinesisfirehose.DeliveryStreamTag[];
}

export function getDeliveryStreamOutput(args: GetDeliveryStreamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeliveryStreamResult> {
    return pulumi.output(args).apply(a => getDeliveryStream(a, opts))
}

export interface GetDeliveryStreamOutputArgs {
    deliveryStreamName: pulumi.Input<string>;
}
