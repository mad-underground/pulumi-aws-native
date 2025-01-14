// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Theme Resource Type.
 */
export class Theme extends pulumi.CustomResource {
    /**
     * Get an existing Theme resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Theme {
        return new Theme(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Theme';

    /**
     * Returns true if the given object is an instance of Theme.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Theme {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Theme.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly awsAccountId!: pulumi.Output<string>;
    public readonly baseThemeId!: pulumi.Output<string>;
    public readonly configuration!: pulumi.Output<outputs.quicksight.ThemeConfiguration>;
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly permissions!: pulumi.Output<outputs.quicksight.ThemeResourcePermission[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly themeId!: pulumi.Output<string>;
    public /*out*/ readonly type!: pulumi.Output<enums.quicksight.ThemeType>;
    public /*out*/ readonly version!: pulumi.Output<outputs.quicksight.ThemeVersion>;
    public readonly versionDescription!: pulumi.Output<string | undefined>;

    /**
     * Create a Theme resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ThemeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.awsAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountId'");
            }
            if ((!args || args.baseThemeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseThemeId'");
            }
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.themeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'themeId'");
            }
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["baseThemeId"] = args ? args.baseThemeId : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["themeId"] = args ? args.themeId : undefined;
            resourceInputs["versionDescription"] = args ? args.versionDescription : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["baseThemeId"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["permissions"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["themeId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
            resourceInputs["versionDescription"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["awsAccountId", "themeId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Theme.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Theme resource.
 */
export interface ThemeArgs {
    awsAccountId: pulumi.Input<string>;
    baseThemeId: pulumi.Input<string>;
    configuration: pulumi.Input<inputs.quicksight.ThemeConfigurationArgs>;
    name?: pulumi.Input<string>;
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.ThemeResourcePermissionArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    themeId: pulumi.Input<string>;
    versionDescription?: pulumi.Input<string>;
}
