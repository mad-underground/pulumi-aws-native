// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::ResourceExplorer2::View Resource Type
 */
export class View extends pulumi.CustomResource {
    /**
     * Get an existing View resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): View {
        return new View(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:resourceexplorer2:View';

    /**
     * Returns true if the given object is an instance of View.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is View {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === View.__pulumiType;
    }

    public readonly filters!: pulumi.Output<outputs.resourceexplorer2.ViewSearchFilter | undefined>;
    public readonly includedProperties!: pulumi.Output<outputs.resourceexplorer2.ViewIncludedProperty[] | undefined>;
    public readonly scope!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly viewArn!: pulumi.Output<string>;
    public readonly viewName!: pulumi.Output<string>;

    /**
     * Create a View resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ViewArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["includedProperties"] = args ? args.includedProperties : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["viewName"] = args ? args.viewName : undefined;
            resourceInputs["viewArn"] = undefined /*out*/;
        } else {
            resourceInputs["filters"] = undefined /*out*/;
            resourceInputs["includedProperties"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["viewArn"] = undefined /*out*/;
            resourceInputs["viewName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["scope", "viewName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(View.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a View resource.
 */
export interface ViewArgs {
    filters?: pulumi.Input<inputs.resourceexplorer2.ViewSearchFilterArgs>;
    includedProperties?: pulumi.Input<pulumi.Input<inputs.resourceexplorer2.ViewIncludedPropertyArgs>[]>;
    scope?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    viewName?: pulumi.Input<string>;
}
