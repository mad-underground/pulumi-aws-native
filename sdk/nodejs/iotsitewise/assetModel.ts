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

    public /*out*/ readonly AssetModelArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
     */
    public readonly AssetModelDescription!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
     */
    public readonly AssetModelHierarchies!: pulumi.Output<outputs.IoTSiteWise.AssetModelAssetModelHierarchy[] | undefined>;
    public /*out*/ readonly AssetModelId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
     */
    public readonly AssetModelName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
     */
    public readonly AssetModelProperties!: pulumi.Output<outputs.IoTSiteWise.AssetModelAssetModelProperty[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
     */
    public readonly Tags!: pulumi.Output<outputs.Tag[] | undefined>;

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
            if ((!args || args.AssetModelName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'AssetModelName'");
            }
            inputs["AssetModelDescription"] = args ? args.AssetModelDescription : undefined;
            inputs["AssetModelHierarchies"] = args ? args.AssetModelHierarchies : undefined;
            inputs["AssetModelName"] = args ? args.AssetModelName : undefined;
            inputs["AssetModelProperties"] = args ? args.AssetModelProperties : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["AssetModelArn"] = undefined /*out*/;
            inputs["AssetModelId"] = undefined /*out*/;
        } else {
            inputs["AssetModelArn"] = undefined /*out*/;
            inputs["AssetModelDescription"] = undefined /*out*/;
            inputs["AssetModelHierarchies"] = undefined /*out*/;
            inputs["AssetModelId"] = undefined /*out*/;
            inputs["AssetModelName"] = undefined /*out*/;
            inputs["AssetModelProperties"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
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
    readonly AssetModelDescription?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
     */
    readonly AssetModelHierarchies?: pulumi.Input<pulumi.Input<inputs.IoTSiteWise.AssetModelAssetModelHierarchy>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
     */
    readonly AssetModelName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
     */
    readonly AssetModelProperties?: pulumi.Input<pulumi.Input<inputs.IoTSiteWise.AssetModelAssetModelProperty>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
     */
    readonly Tags?: pulumi.Input<pulumi.Input<inputs.Tag>[]>;
}
