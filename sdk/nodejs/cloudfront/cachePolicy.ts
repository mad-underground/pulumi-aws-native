// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFront::CachePolicy
 */
export class CachePolicy extends pulumi.CustomResource {
    /**
     * Get an existing CachePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CachePolicy {
        return new CachePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudfront:CachePolicy';

    /**
     * Returns true if the given object is an instance of CachePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CachePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CachePolicy.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly cachePolicyConfig!: pulumi.Output<outputs.cloudfront.CachePolicyConfig>;
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;

    /**
     * Create a CachePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CachePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cachePolicyConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cachePolicyConfig'");
            }
            resourceInputs["cachePolicyConfig"] = args ? args.cachePolicyConfig : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["cachePolicyConfig"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CachePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CachePolicy resource.
 */
export interface CachePolicyArgs {
    cachePolicyConfig: pulumi.Input<inputs.cloudfront.CachePolicyConfigArgs>;
}
