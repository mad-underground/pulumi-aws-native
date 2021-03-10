// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html
 */
export class ComponentVersion extends pulumi.CustomResource {
    /**
     * Get an existing ComponentVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ComponentVersion {
        return new ComponentVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:GreengrassV2:ComponentVersion';

    /**
     * Returns true if the given object is an instance of ComponentVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComponentVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComponentVersion.__pulumiType;
    }

    public /*out*/ readonly Arn!: pulumi.Output<string>;
    public /*out*/ readonly ComponentName!: pulumi.Output<string>;
    public /*out*/ readonly ComponentVersion!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-inlinerecipe
     */
    public readonly InlineRecipe!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-lambdafunction
     */
    public readonly LambdaFunction!: pulumi.Output<outputs.GreengrassV2.ComponentVersionLambdaFunctionRecipeSource | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-tags
     */
    public readonly Tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ComponentVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComponentVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            inputs["InlineRecipe"] = args ? args.InlineRecipe : undefined;
            inputs["LambdaFunction"] = args ? args.LambdaFunction : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["Arn"] = undefined /*out*/;
            inputs["ComponentName"] = undefined /*out*/;
            inputs["ComponentVersion"] = undefined /*out*/;
        } else {
            inputs["Arn"] = undefined /*out*/;
            inputs["ComponentName"] = undefined /*out*/;
            inputs["ComponentVersion"] = undefined /*out*/;
            inputs["InlineRecipe"] = undefined /*out*/;
            inputs["LambdaFunction"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ComponentVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ComponentVersion resource.
 */
export interface ComponentVersionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-inlinerecipe
     */
    readonly InlineRecipe?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-lambdafunction
     */
    readonly LambdaFunction?: pulumi.Input<inputs.GreengrassV2.ComponentVersionLambdaFunctionRecipeSource>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-tags
     */
    readonly Tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
