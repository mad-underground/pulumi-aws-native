// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html
 */
export class PrefixList extends pulumi.CustomResource {
    /**
     * Get an existing PrefixList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrefixList {
        return new PrefixList(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:EC2:PrefixList';

    /**
     * Returns true if the given object is an instance of PrefixList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrefixList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrefixList.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-addressfamily
     */
    public readonly addressFamily!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-entries
     */
    public readonly entries!: pulumi.Output<outputs.EC2.PrefixListEntry[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-maxentries
     */
    public readonly maxEntries!: pulumi.Output<number>;
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    public /*out*/ readonly prefixListId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-prefixlistname
     */
    public readonly prefixListName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a PrefixList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrefixListArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.addressFamily === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.maxEntries === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'maxEntries'");
            }
            if ((!args || args.prefixListName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'prefixListName'");
            }
            inputs["addressFamily"] = args ? args.addressFamily : undefined;
            inputs["entries"] = args ? args.entries : undefined;
            inputs["maxEntries"] = args ? args.maxEntries : undefined;
            inputs["prefixListName"] = args ? args.prefixListName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["prefixListId"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        } else {
            inputs["addressFamily"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["entries"] = undefined /*out*/;
            inputs["maxEntries"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["prefixListId"] = undefined /*out*/;
            inputs["prefixListName"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PrefixList.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrefixList resource.
 */
export interface PrefixListArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-addressfamily
     */
    readonly addressFamily: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-entries
     */
    readonly entries?: pulumi.Input<pulumi.Input<inputs.EC2.PrefixListEntry>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-maxentries
     */
    readonly maxEntries: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-prefixlistname
     */
    readonly prefixListName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.Tag>[]>;
}
