// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html
 */
export class CertificateAuthorityActivation extends pulumi.CustomResource {
    /**
     * Get an existing CertificateAuthorityActivation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateAuthorityActivation {
        return new CertificateAuthorityActivation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ACMPCA:CertificateAuthorityActivation';

    /**
     * Returns true if the given object is an instance of CertificateAuthorityActivation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateAuthorityActivation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateAuthorityActivation.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificate
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificateauthorityarn
     */
    public readonly certificateAuthorityArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificatechain
     */
    public readonly certificateChain!: pulumi.Output<string | undefined>;
    public /*out*/ readonly completeCertificateChain!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-status
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateAuthorityActivation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateAuthorityActivationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            if ((!args || args.certificateAuthorityArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAuthorityArn'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["certificateAuthorityArn"] = args ? args.certificateAuthorityArn : undefined;
            inputs["certificateChain"] = args ? args.certificateChain : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["completeCertificateChain"] = undefined /*out*/;
        } else {
            inputs["certificate"] = undefined /*out*/;
            inputs["certificateAuthorityArn"] = undefined /*out*/;
            inputs["certificateChain"] = undefined /*out*/;
            inputs["completeCertificateChain"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CertificateAuthorityActivation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateAuthorityActivation resource.
 */
export interface CertificateAuthorityActivationArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificate
     */
    certificate: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificateauthorityarn
     */
    certificateAuthorityArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificatechain
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-status
     */
    status?: pulumi.Input<string>;
}
