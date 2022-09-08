// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.
 */
export class KeySigningKey extends pulumi.CustomResource {
    /**
     * Get an existing KeySigningKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KeySigningKey {
        return new KeySigningKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53:KeySigningKey';

    /**
     * Returns true if the given object is an instance of KeySigningKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeySigningKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeySigningKey.__pulumiType;
    }

    /**
     * The unique string (ID) used to identify a hosted zone.
     */
    public readonly hostedZoneId!: pulumi.Output<string>;
    /**
     * The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
     */
    public readonly keyManagementServiceArn!: pulumi.Output<string>;
    /**
     * An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
     */
    public readonly status!: pulumi.Output<enums.route53.KeySigningKeyStatus>;

    /**
     * Create a KeySigningKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeySigningKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.hostedZoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostedZoneId'");
            }
            if ((!args || args.keyManagementServiceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyManagementServiceArn'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["hostedZoneId"] = args ? args.hostedZoneId : undefined;
            resourceInputs["keyManagementServiceArn"] = args ? args.keyManagementServiceArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        } else {
            resourceInputs["hostedZoneId"] = undefined /*out*/;
            resourceInputs["keyManagementServiceArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeySigningKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KeySigningKey resource.
 */
export interface KeySigningKeyArgs {
    /**
     * The unique string (ID) used to identify a hosted zone.
     */
    hostedZoneId: pulumi.Input<string>;
    /**
     * The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
     */
    keyManagementServiceArn: pulumi.Input<string>;
    /**
     * An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
     */
    name?: pulumi.Input<string>;
    /**
     * A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
     */
    status: pulumi.Input<enums.route53.KeySigningKeyStatus>;
}
