// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html
 */
export class Asset extends pulumi.CustomResource {
    /**
     * Get an existing Asset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Asset {
        return new Asset(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:MediaPackage:Asset';

    /**
     * Returns true if the given object is an instance of Asset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Asset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Asset.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-egressendpoints
     */
    public readonly egressEndpoints!: pulumi.Output<outputs.MediaPackage.AssetEgressEndpoint[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-id
     */
    public readonly id!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-packaginggroupid
     */
    public readonly packagingGroupId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-resourceid
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-sourcearn
     */
    public readonly sourceArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-sourcerolearn
     */
    public readonly sourceRoleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Asset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.id === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'id'");
            }
            if ((!args || args.packagingGroupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'packagingGroupId'");
            }
            if ((!args || args.sourceArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sourceArn'");
            }
            if ((!args || args.sourceRoleArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sourceRoleArn'");
            }
            inputs["egressEndpoints"] = args ? args.egressEndpoints : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["packagingGroupId"] = args ? args.packagingGroupId : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["sourceArn"] = args ? args.sourceArn : undefined;
            inputs["sourceRoleArn"] = args ? args.sourceRoleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["egressEndpoints"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["packagingGroupId"] = undefined /*out*/;
            inputs["resourceId"] = undefined /*out*/;
            inputs["sourceArn"] = undefined /*out*/;
            inputs["sourceRoleArn"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Asset.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Asset resource.
 */
export interface AssetArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-egressendpoints
     */
    readonly egressEndpoints?: pulumi.Input<pulumi.Input<inputs.MediaPackage.AssetEgressEndpoint>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-id
     */
    readonly id: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-packaginggroupid
     */
    readonly packagingGroupId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-resourceid
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-sourcearn
     */
    readonly sourceArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-sourcerolearn
     */
    readonly sourceRoleArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.Tag>[]>;
}
