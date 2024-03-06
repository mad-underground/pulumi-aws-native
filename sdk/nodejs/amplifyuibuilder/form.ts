// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::AmplifyUIBuilder::Form Resource Type
 */
export class Form extends pulumi.CustomResource {
    /**
     * Get an existing Form resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Form {
        return new Form(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:amplifyuibuilder:Form';

    /**
     * Returns true if the given object is an instance of Form.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Form {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Form.__pulumiType;
    }

    public readonly appId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly cta!: pulumi.Output<outputs.amplifyuibuilder.FormCta | undefined>;
    public readonly dataType!: pulumi.Output<outputs.amplifyuibuilder.FormDataTypeConfig | undefined>;
    public readonly environmentName!: pulumi.Output<string | undefined>;
    public readonly fields!: pulumi.Output<{[key: string]: outputs.amplifyuibuilder.FormFieldConfig} | undefined>;
    public readonly formActionType!: pulumi.Output<enums.amplifyuibuilder.FormActionType | undefined>;
    public readonly labelDecorator!: pulumi.Output<enums.amplifyuibuilder.FormLabelDecorator | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly schemaVersion!: pulumi.Output<string | undefined>;
    public readonly sectionalElements!: pulumi.Output<{[key: string]: outputs.amplifyuibuilder.FormSectionalElement} | undefined>;
    public readonly style!: pulumi.Output<outputs.amplifyuibuilder.FormStyle | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Form resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FormArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["cta"] = args ? args.cta : undefined;
            resourceInputs["dataType"] = args ? args.dataType : undefined;
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["formActionType"] = args ? args.formActionType : undefined;
            resourceInputs["labelDecorator"] = args ? args.labelDecorator : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["schemaVersion"] = args ? args.schemaVersion : undefined;
            resourceInputs["sectionalElements"] = args ? args.sectionalElements : undefined;
            resourceInputs["style"] = args ? args.style : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["cta"] = undefined /*out*/;
            resourceInputs["dataType"] = undefined /*out*/;
            resourceInputs["environmentName"] = undefined /*out*/;
            resourceInputs["fields"] = undefined /*out*/;
            resourceInputs["formActionType"] = undefined /*out*/;
            resourceInputs["labelDecorator"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["schemaVersion"] = undefined /*out*/;
            resourceInputs["sectionalElements"] = undefined /*out*/;
            resourceInputs["style"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["appId", "environmentName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Form.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Form resource.
 */
export interface FormArgs {
    appId?: pulumi.Input<string>;
    cta?: pulumi.Input<inputs.amplifyuibuilder.FormCtaArgs>;
    dataType?: pulumi.Input<inputs.amplifyuibuilder.FormDataTypeConfigArgs>;
    environmentName?: pulumi.Input<string>;
    fields?: pulumi.Input<{[key: string]: pulumi.Input<inputs.amplifyuibuilder.FormFieldConfigArgs>}>;
    formActionType?: pulumi.Input<enums.amplifyuibuilder.FormActionType>;
    labelDecorator?: pulumi.Input<enums.amplifyuibuilder.FormLabelDecorator>;
    name?: pulumi.Input<string>;
    schemaVersion?: pulumi.Input<string>;
    sectionalElements?: pulumi.Input<{[key: string]: pulumi.Input<inputs.amplifyuibuilder.FormSectionalElementArgs>}>;
    style?: pulumi.Input<inputs.amplifyuibuilder.FormStyleArgs>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
