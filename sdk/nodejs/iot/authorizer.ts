// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html
 */
export class Authorizer extends pulumi.CustomResource {
    /**
     * Get an existing Authorizer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Authorizer {
        return new Authorizer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:IoT:Authorizer';

    /**
     * Returns true if the given object is an instance of Authorizer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Authorizer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Authorizer.__pulumiType;
    }

    public /*out*/ readonly Arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-authorizerfunctionarn
     */
    public readonly AuthorizerFunctionArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-authorizername
     */
    public readonly AuthorizerName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-signingdisabled
     */
    public readonly SigningDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-status
     */
    public readonly Status!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tags
     */
    public readonly Tags!: pulumi.Output<outputs.IoT.AuthorizerTags | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tokenkeyname
     */
    public readonly TokenKeyName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tokensigningpublickeys
     */
    public readonly TokenSigningPublicKeys!: pulumi.Output<outputs.IoT.AuthorizerTokenSigningPublicKeys | undefined>;

    /**
     * Create a Authorizer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.AuthorizerFunctionArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'AuthorizerFunctionArn'");
            }
            inputs["AuthorizerFunctionArn"] = args ? args.AuthorizerFunctionArn : undefined;
            inputs["AuthorizerName"] = args ? args.AuthorizerName : undefined;
            inputs["SigningDisabled"] = args ? args.SigningDisabled : undefined;
            inputs["Status"] = args ? args.Status : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["TokenKeyName"] = args ? args.TokenKeyName : undefined;
            inputs["TokenSigningPublicKeys"] = args ? args.TokenSigningPublicKeys : undefined;
            inputs["Arn"] = undefined /*out*/;
        } else {
            inputs["Arn"] = undefined /*out*/;
            inputs["AuthorizerFunctionArn"] = undefined /*out*/;
            inputs["AuthorizerName"] = undefined /*out*/;
            inputs["SigningDisabled"] = undefined /*out*/;
            inputs["Status"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
            inputs["TokenKeyName"] = undefined /*out*/;
            inputs["TokenSigningPublicKeys"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Authorizer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Authorizer resource.
 */
export interface AuthorizerArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-authorizerfunctionarn
     */
    readonly AuthorizerFunctionArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-authorizername
     */
    readonly AuthorizerName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-signingdisabled
     */
    readonly SigningDisabled?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-status
     */
    readonly Status?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tags
     */
    readonly Tags?: pulumi.Input<inputs.IoT.AuthorizerTags>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tokenkeyname
     */
    readonly TokenKeyName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tokensigningpublickeys
     */
    readonly TokenSigningPublicKeys?: pulumi.Input<inputs.IoT.AuthorizerTokenSigningPublicKeys>;
}
