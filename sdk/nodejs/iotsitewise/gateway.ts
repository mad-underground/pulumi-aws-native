// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:IoTSiteWise:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewaycapabilitysummaries
     */
    public readonly gatewayCapabilitySummaries!: pulumi.Output<outputs.IoTSiteWise.GatewayGatewayCapabilitySummary[] | undefined>;
    public /*out*/ readonly gatewayId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayname
     */
    public readonly gatewayName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayplatform
     */
    public readonly gatewayPlatform!: pulumi.Output<outputs.IoTSiteWise.GatewayGatewayPlatform>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.gatewayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayName'");
            }
            if ((!args || args.gatewayPlatform === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayPlatform'");
            }
            inputs["gatewayCapabilitySummaries"] = args ? args.gatewayCapabilitySummaries : undefined;
            inputs["gatewayName"] = args ? args.gatewayName : undefined;
            inputs["gatewayPlatform"] = args ? args.gatewayPlatform : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["gatewayId"] = undefined /*out*/;
        } else {
            inputs["gatewayCapabilitySummaries"] = undefined /*out*/;
            inputs["gatewayId"] = undefined /*out*/;
            inputs["gatewayName"] = undefined /*out*/;
            inputs["gatewayPlatform"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Gateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewaycapabilitysummaries
     */
    gatewayCapabilitySummaries?: pulumi.Input<pulumi.Input<inputs.IoTSiteWise.GatewayGatewayCapabilitySummaryArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayname
     */
    gatewayName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayplatform
     */
    gatewayPlatform: pulumi.Input<inputs.IoTSiteWise.GatewayGatewayPlatformArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
