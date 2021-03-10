// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html
 */
export class RealtimeLogConfig extends pulumi.CustomResource {
    /**
     * Get an existing RealtimeLogConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RealtimeLogConfig {
        return new RealtimeLogConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:CloudFront:RealtimeLogConfig';

    /**
     * Returns true if the given object is an instance of RealtimeLogConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealtimeLogConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealtimeLogConfig.__pulumiType;
    }

    public /*out*/ readonly Arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-endpoints
     */
    public readonly EndPoints!: pulumi.Output<outputs.CloudFront.RealtimeLogConfigEndPoint[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-fields
     */
    public readonly Fields!: pulumi.Output<string[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-name
     */
    public readonly Name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-samplingrate
     */
    public readonly SamplingRate!: pulumi.Output<number>;

    /**
     * Create a RealtimeLogConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealtimeLogConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.EndPoints === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'EndPoints'");
            }
            if ((!args || args.Fields === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'Fields'");
            }
            if ((!args || args.Name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'Name'");
            }
            if ((!args || args.SamplingRate === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'SamplingRate'");
            }
            inputs["EndPoints"] = args ? args.EndPoints : undefined;
            inputs["Fields"] = args ? args.Fields : undefined;
            inputs["Name"] = args ? args.Name : undefined;
            inputs["SamplingRate"] = args ? args.SamplingRate : undefined;
            inputs["Arn"] = undefined /*out*/;
        } else {
            inputs["Arn"] = undefined /*out*/;
            inputs["EndPoints"] = undefined /*out*/;
            inputs["Fields"] = undefined /*out*/;
            inputs["Name"] = undefined /*out*/;
            inputs["SamplingRate"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RealtimeLogConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RealtimeLogConfig resource.
 */
export interface RealtimeLogConfigArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-endpoints
     */
    readonly EndPoints: pulumi.Input<pulumi.Input<inputs.CloudFront.RealtimeLogConfigEndPoint>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-fields
     */
    readonly Fields: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-name
     */
    readonly Name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-samplingrate
     */
    readonly SamplingRate: pulumi.Input<number>;
}
