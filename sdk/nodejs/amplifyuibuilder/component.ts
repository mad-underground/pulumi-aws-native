// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::AmplifyUIBuilder::Component Resource Type
 */
export class Component extends pulumi.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Component {
        return new Component(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:amplifyuibuilder:Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    public readonly appId!: pulumi.Output<string | undefined>;
    public readonly bindingProperties!: pulumi.Output<{[key: string]: outputs.amplifyuibuilder.ComponentBindingPropertiesValue} | undefined>;
    public readonly children!: pulumi.Output<outputs.amplifyuibuilder.ComponentChild[] | undefined>;
    public readonly collectionProperties!: pulumi.Output<{[key: string]: outputs.amplifyuibuilder.ComponentDataConfiguration} | undefined>;
    public readonly componentType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly environmentName!: pulumi.Output<string | undefined>;
    public readonly events!: pulumi.Output<{[key: string]: outputs.amplifyuibuilder.ComponentEvent} | undefined>;
    public /*out*/ readonly modifiedAt!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly overrides!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly properties!: pulumi.Output<{[key: string]: outputs.amplifyuibuilder.ComponentProperty} | undefined>;
    public readonly schemaVersion!: pulumi.Output<string | undefined>;
    public readonly sourceId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly variants!: pulumi.Output<outputs.amplifyuibuilder.ComponentVariant[] | undefined>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComponentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["bindingProperties"] = args ? args.bindingProperties : undefined;
            resourceInputs["children"] = args ? args.children : undefined;
            resourceInputs["collectionProperties"] = args ? args.collectionProperties : undefined;
            resourceInputs["componentType"] = args ? args.componentType : undefined;
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["events"] = args ? args.events : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrides"] = args ? args.overrides : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["schemaVersion"] = args ? args.schemaVersion : undefined;
            resourceInputs["sourceId"] = args ? args.sourceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["variants"] = args ? args.variants : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["bindingProperties"] = undefined /*out*/;
            resourceInputs["children"] = undefined /*out*/;
            resourceInputs["collectionProperties"] = undefined /*out*/;
            resourceInputs["componentType"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["environmentName"] = undefined /*out*/;
            resourceInputs["events"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["overrides"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["schemaVersion"] = undefined /*out*/;
            resourceInputs["sourceId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["variants"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["appId", "environmentName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Component.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    appId?: pulumi.Input<string>;
    bindingProperties?: pulumi.Input<{[key: string]: pulumi.Input<inputs.amplifyuibuilder.ComponentBindingPropertiesValueArgs>}>;
    children?: pulumi.Input<pulumi.Input<inputs.amplifyuibuilder.ComponentChildArgs>[]>;
    collectionProperties?: pulumi.Input<{[key: string]: pulumi.Input<inputs.amplifyuibuilder.ComponentDataConfigurationArgs>}>;
    componentType?: pulumi.Input<string>;
    environmentName?: pulumi.Input<string>;
    events?: pulumi.Input<{[key: string]: pulumi.Input<inputs.amplifyuibuilder.ComponentEventArgs>}>;
    name?: pulumi.Input<string>;
    overrides?: pulumi.Input<{[key: string]: any}>;
    properties?: pulumi.Input<{[key: string]: pulumi.Input<inputs.amplifyuibuilder.ComponentPropertyArgs>}>;
    schemaVersion?: pulumi.Input<string>;
    sourceId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    variants?: pulumi.Input<pulumi.Input<inputs.amplifyuibuilder.ComponentVariantArgs>[]>;
}
