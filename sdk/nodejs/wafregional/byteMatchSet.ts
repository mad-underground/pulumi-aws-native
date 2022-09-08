// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::ByteMatchSet
 *
 * @deprecated ByteMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ByteMatchSet extends pulumi.CustomResource {
    /**
     * Get an existing ByteMatchSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ByteMatchSet {
        pulumi.log.warn("ByteMatchSet is deprecated: ByteMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ByteMatchSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafregional:ByteMatchSet';

    /**
     * Returns true if the given object is an instance of ByteMatchSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ByteMatchSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ByteMatchSet.__pulumiType;
    }

    public readonly byteMatchTuples!: pulumi.Output<outputs.wafregional.ByteMatchSetByteMatchTuple[] | undefined>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ByteMatchSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ByteMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: ByteMatchSetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ByteMatchSet is deprecated: ByteMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["byteMatchTuples"] = args ? args.byteMatchTuples : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        } else {
            resourceInputs["byteMatchTuples"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ByteMatchSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ByteMatchSet resource.
 */
export interface ByteMatchSetArgs {
    byteMatchTuples?: pulumi.Input<pulumi.Input<inputs.wafregional.ByteMatchSetByteMatchTupleArgs>[]>;
    name?: pulumi.Input<string>;
}
