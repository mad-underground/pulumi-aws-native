// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFront::KeyValueStore
 */
export class KeyValueStore extends pulumi.CustomResource {
    /**
     * Get an existing KeyValueStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KeyValueStore {
        return new KeyValueStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudfront:KeyValueStore';

    /**
     * Returns true if the given object is an instance of KeyValueStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyValueStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyValueStore.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly importSource!: pulumi.Output<outputs.cloudfront.KeyValueStoreImportSource | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a KeyValueStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyValueStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["importSource"] = args ? args.importSource : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["comment"] = undefined /*out*/;
            resourceInputs["importSource"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(KeyValueStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KeyValueStore resource.
 */
export interface KeyValueStoreArgs {
    comment?: pulumi.Input<string>;
    importSource?: pulumi.Input<inputs.cloudfront.KeyValueStoreImportSourceArgs>;
    name?: pulumi.Input<string>;
}
