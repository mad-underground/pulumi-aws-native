// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html
 */
export class CloudFrontOriginAccessIdentity extends pulumi.CustomResource {
    /**
     * Get an existing CloudFrontOriginAccessIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CloudFrontOriginAccessIdentity {
        return new CloudFrontOriginAccessIdentity(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:CloudFront:CloudFrontOriginAccessIdentity';

    /**
     * Returns true if the given object is an instance of CloudFrontOriginAccessIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudFrontOriginAccessIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudFrontOriginAccessIdentity.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
     */
    public readonly cloudFrontOriginAccessIdentityConfig!: pulumi.Output<outputs.CloudFront.CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig>;
    public /*out*/ readonly s3CanonicalUserId!: pulumi.Output<string>;

    /**
     * Create a CloudFrontOriginAccessIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudFrontOriginAccessIdentityArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.cloudFrontOriginAccessIdentityConfig === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'cloudFrontOriginAccessIdentityConfig'");
            }
            inputs["cloudFrontOriginAccessIdentityConfig"] = args ? args.cloudFrontOriginAccessIdentityConfig : undefined;
            inputs["s3CanonicalUserId"] = undefined /*out*/;
        } else {
            inputs["cloudFrontOriginAccessIdentityConfig"] = undefined /*out*/;
            inputs["s3CanonicalUserId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CloudFrontOriginAccessIdentity.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloudFrontOriginAccessIdentity resource.
 */
export interface CloudFrontOriginAccessIdentityArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
     */
    readonly cloudFrontOriginAccessIdentityConfig: pulumi.Input<inputs.CloudFront.CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig>;
}
