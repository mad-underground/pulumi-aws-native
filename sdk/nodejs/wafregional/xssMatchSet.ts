// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::XssMatchSet
 *
 * @deprecated XssMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class XssMatchSet extends pulumi.CustomResource {
    /**
     * Get an existing XssMatchSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): XssMatchSet {
        pulumi.log.warn("XssMatchSet is deprecated: XssMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new XssMatchSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafregional:XssMatchSet';

    /**
     * Returns true if the given object is an instance of XssMatchSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is XssMatchSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === XssMatchSet.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly xssMatchTuples!: pulumi.Output<outputs.wafregional.XssMatchSetXssMatchTuple[] | undefined>;

    /**
     * Create a XssMatchSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated XssMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: XssMatchSetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("XssMatchSet is deprecated: XssMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["xssMatchTuples"] = args ? args.xssMatchTuples : undefined;
        } else {
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["xssMatchTuples"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(XssMatchSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a XssMatchSet resource.
 */
export interface XssMatchSetArgs {
    name?: pulumi.Input<string>;
    xssMatchTuples?: pulumi.Input<pulumi.Input<inputs.wafregional.XssMatchSetXssMatchTupleArgs>[]>;
}
