// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Omics::ReferenceStore Resource Type
 */
export class ReferenceStore extends pulumi.CustomResource {
    /**
     * Get an existing ReferenceStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReferenceStore {
        return new ReferenceStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:omics:ReferenceStore';

    /**
     * Returns true if the given object is an instance of ReferenceStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReferenceStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReferenceStore.__pulumiType;
    }

    /**
     * The store's ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * When the store was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * A description for the store.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A name for the store.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly referenceStoreId!: pulumi.Output<string>;
    public readonly sseConfig!: pulumi.Output<outputs.omics.ReferenceStoreSseConfig | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ReferenceStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ReferenceStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sseConfig"] = args ? args.sseConfig : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["referenceStoreId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["referenceStoreId"] = undefined /*out*/;
            resourceInputs["sseConfig"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "name", "sseConfig", "tags.*"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ReferenceStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReferenceStore resource.
 */
export interface ReferenceStoreArgs {
    /**
     * A description for the store.
     */
    description?: pulumi.Input<string>;
    /**
     * A name for the store.
     */
    name?: pulumi.Input<string>;
    sseConfig?: pulumi.Input<inputs.omics.ReferenceStoreSseConfigArgs>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
