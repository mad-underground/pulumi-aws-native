// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lightsail:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    public /*out*/ readonly certificateArn!: pulumi.Output<string>;
    /**
     * The name for the certificate.
     */
    public readonly certificateName!: pulumi.Output<string>;
    /**
     * The domain name (e.g., example.com ) for the certificate.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The validation status of the certificate.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
     */
    public readonly subjectAlternativeNames!: pulumi.Output<string[] | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.lightsail.CertificateTag[] | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["subjectAlternativeNames"] = args ? args.subjectAlternativeNames : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["certificateArn"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["certificateArn"] = undefined /*out*/;
            inputs["certificateName"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["subjectAlternativeNames"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The name for the certificate.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * The domain name (e.g., example.com ) for the certificate.
     */
    domainName: pulumi.Input<string>;
    /**
     * An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
     */
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.lightsail.CertificateTagArgs>[]>;
}
