// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Kinesis::StreamConsumer
 *
 * @deprecated StreamConsumer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class StreamConsumer extends pulumi.CustomResource {
    /**
     * Get an existing StreamConsumer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StreamConsumer {
        pulumi.log.warn("StreamConsumer is deprecated: StreamConsumer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new StreamConsumer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:kinesis:StreamConsumer';

    /**
     * Returns true if the given object is an instance of StreamConsumer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamConsumer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamConsumer.__pulumiType;
    }

    public /*out*/ readonly consumerARN!: pulumi.Output<string>;
    public /*out*/ readonly consumerCreationTimestamp!: pulumi.Output<string>;
    public readonly consumerName!: pulumi.Output<string>;
    public /*out*/ readonly consumerStatus!: pulumi.Output<string>;
    public readonly streamARN!: pulumi.Output<string>;

    /**
     * Create a StreamConsumer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated StreamConsumer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: StreamConsumerArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("StreamConsumer is deprecated: StreamConsumer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.consumerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consumerName'");
            }
            if ((!args || args.streamARN === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamARN'");
            }
            resourceInputs["consumerName"] = args ? args.consumerName : undefined;
            resourceInputs["streamARN"] = args ? args.streamARN : undefined;
            resourceInputs["consumerARN"] = undefined /*out*/;
            resourceInputs["consumerCreationTimestamp"] = undefined /*out*/;
            resourceInputs["consumerStatus"] = undefined /*out*/;
        } else {
            resourceInputs["consumerARN"] = undefined /*out*/;
            resourceInputs["consumerCreationTimestamp"] = undefined /*out*/;
            resourceInputs["consumerName"] = undefined /*out*/;
            resourceInputs["consumerStatus"] = undefined /*out*/;
            resourceInputs["streamARN"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StreamConsumer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StreamConsumer resource.
 */
export interface StreamConsumerArgs {
    consumerName: pulumi.Input<string>;
    streamARN: pulumi.Input<string>;
}
