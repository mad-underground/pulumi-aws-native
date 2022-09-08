// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Registers a CA Certificate in IoT.
 */
export class CACertificate extends pulumi.CustomResource {
    /**
     * Get an existing CACertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CACertificate {
        return new CACertificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iot:CACertificate';

    /**
     * Returns true if the given object is an instance of CACertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CACertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CACertificate.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly autoRegistrationStatus!: pulumi.Output<enums.iot.CACertificateAutoRegistrationStatus | undefined>;
    public readonly cACertificatePem!: pulumi.Output<string>;
    public readonly certificateMode!: pulumi.Output<enums.iot.CACertificateCertificateMode | undefined>;
    public readonly registrationConfig!: pulumi.Output<outputs.iot.CACertificateRegistrationConfig | undefined>;
    public readonly status!: pulumi.Output<enums.iot.CACertificateStatus>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.iot.CACertificateTag[] | undefined>;
    /**
     * The private key verification certificate.
     */
    public readonly verificationCertificatePem!: pulumi.Output<string | undefined>;

    /**
     * Create a CACertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CACertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cACertificatePem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cACertificatePem'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["autoRegistrationStatus"] = args ? args.autoRegistrationStatus : undefined;
            resourceInputs["cACertificatePem"] = args ? args.cACertificatePem : undefined;
            resourceInputs["certificateMode"] = args ? args.certificateMode : undefined;
            resourceInputs["registrationConfig"] = args ? args.registrationConfig : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["verificationCertificatePem"] = args ? args.verificationCertificatePem : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["autoRegistrationStatus"] = undefined /*out*/;
            resourceInputs["cACertificatePem"] = undefined /*out*/;
            resourceInputs["certificateMode"] = undefined /*out*/;
            resourceInputs["registrationConfig"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["verificationCertificatePem"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CACertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CACertificate resource.
 */
export interface CACertificateArgs {
    autoRegistrationStatus?: pulumi.Input<enums.iot.CACertificateAutoRegistrationStatus>;
    cACertificatePem: pulumi.Input<string>;
    certificateMode?: pulumi.Input<enums.iot.CACertificateCertificateMode>;
    registrationConfig?: pulumi.Input<inputs.iot.CACertificateRegistrationConfigArgs>;
    status: pulumi.Input<enums.iot.CACertificateStatus>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iot.CACertificateTagArgs>[]>;
    /**
     * The private key verification certificate.
     */
    verificationCertificatePem?: pulumi.Input<string>;
}
