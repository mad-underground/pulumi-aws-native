// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::EventStream
 *
 * @deprecated EventStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class EventStream extends pulumi.CustomResource {
    /**
     * Get an existing EventStream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventStream {
        pulumi.log.warn("EventStream is deprecated: EventStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new EventStream(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:EventStream';

    /**
     * Returns true if the given object is an instance of EventStream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventStream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventStream.__pulumiType;
    }

    public readonly applicationId!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly destinationStreamArn!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string>;

    /**
     * Create a EventStream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated EventStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: EventStreamArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EventStream is deprecated: EventStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.destinationStreamArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationStreamArn'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["destinationStreamArn"] = args ? args.destinationStreamArn : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["destinationStreamArn"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EventStream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventStream resource.
 */
export interface EventStreamArgs {
    applicationId: pulumi.Input<string>;
    destinationStreamArn: pulumi.Input<string>;
    roleArn: pulumi.Input<string>;
}
