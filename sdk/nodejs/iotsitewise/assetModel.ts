// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html
 */
export class AssetModel extends pulumi.CustomResource {
    /**
     * Get an existing AssetModel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AssetModel {
        return new AssetModel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:IoTSiteWise:AssetModel';

    /**
     * Returns true if the given object is an instance of AssetModel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AssetModel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AssetModel.__pulumiType;
    }

    public /*out*/ readonly assetModelArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
     */
    public readonly assetModelDescription!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
     */
    public readonly assetModelHierarchies!: pulumi.Output<outputs.IoTSiteWise.AssetModelAssetModelHierarchy[] | undefined>;
    public /*out*/ readonly assetModelId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
     */
    public readonly assetModelName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
     */
    public readonly assetModelProperties!: pulumi.Output<outputs.IoTSiteWise.AssetModelAssetModelProperty[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a AssetModel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssetModelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.assetModelName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'assetModelName'");
            }
            inputs["assetModelDescription"] = args ? args.assetModelDescription : undefined;
            inputs["assetModelHierarchies"] = args ? args.assetModelHierarchies : undefined;
            inputs["assetModelName"] = args ? args.assetModelName : undefined;
            inputs["assetModelProperties"] = args ? args.assetModelProperties : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["assetModelArn"] = undefined /*out*/;
            inputs["assetModelId"] = undefined /*out*/;
        } else {
            inputs["assetModelArn"] = undefined /*out*/;
            inputs["assetModelDescription"] = undefined /*out*/;
            inputs["assetModelHierarchies"] = undefined /*out*/;
            inputs["assetModelId"] = undefined /*out*/;
            inputs["assetModelName"] = undefined /*out*/;
            inputs["assetModelProperties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AssetModel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AssetModel resource.
 */
export interface AssetModelArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
     */
    readonly assetModelDescription?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
     */
    readonly assetModelHierarchies?: pulumi.Input<pulumi.Input<inputs.IoTSiteWise.AssetModelAssetModelHierarchy>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
     */
    readonly assetModelName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
     */
    readonly assetModelProperties?: pulumi.Input<pulumi.Input<inputs.IoTSiteWise.AssetModelAssetModelProperty>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.Tag>[]>;
}
