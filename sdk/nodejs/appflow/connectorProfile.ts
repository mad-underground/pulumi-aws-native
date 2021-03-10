// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html
 */
export class ConnectorProfile extends pulumi.CustomResource {
    /**
     * Get an existing ConnectorProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectorProfile {
        return new ConnectorProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:AppFlow:ConnectorProfile';

    /**
     * Returns true if the given object is an instance of ConnectorProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectorProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectorProfile.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectionmode
     */
    public readonly ConnectionMode!: pulumi.Output<string>;
    public /*out*/ readonly ConnectorProfileArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorprofileconfig
     */
    public readonly ConnectorProfileConfig!: pulumi.Output<outputs.AppFlow.ConnectorProfileConnectorProfileConfig | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorprofilename
     */
    public readonly ConnectorProfileName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectortype
     */
    public readonly ConnectorType!: pulumi.Output<string>;
    public /*out*/ readonly CredentialsArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-kmsarn
     */
    public readonly KMSArn!: pulumi.Output<string | undefined>;

    /**
     * Create a ConnectorProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.ConnectionMode === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'ConnectionMode'");
            }
            if ((!args || args.ConnectorProfileName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'ConnectorProfileName'");
            }
            if ((!args || args.ConnectorType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'ConnectorType'");
            }
            inputs["ConnectionMode"] = args ? args.ConnectionMode : undefined;
            inputs["ConnectorProfileConfig"] = args ? args.ConnectorProfileConfig : undefined;
            inputs["ConnectorProfileName"] = args ? args.ConnectorProfileName : undefined;
            inputs["ConnectorType"] = args ? args.ConnectorType : undefined;
            inputs["KMSArn"] = args ? args.KMSArn : undefined;
            inputs["ConnectorProfileArn"] = undefined /*out*/;
            inputs["CredentialsArn"] = undefined /*out*/;
        } else {
            inputs["ConnectionMode"] = undefined /*out*/;
            inputs["ConnectorProfileArn"] = undefined /*out*/;
            inputs["ConnectorProfileConfig"] = undefined /*out*/;
            inputs["ConnectorProfileName"] = undefined /*out*/;
            inputs["ConnectorType"] = undefined /*out*/;
            inputs["CredentialsArn"] = undefined /*out*/;
            inputs["KMSArn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConnectorProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectorProfile resource.
 */
export interface ConnectorProfileArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectionmode
     */
    readonly ConnectionMode: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorprofileconfig
     */
    readonly ConnectorProfileConfig?: pulumi.Input<inputs.AppFlow.ConnectorProfileConnectorProfileConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectorprofilename
     */
    readonly ConnectorProfileName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-connectortype
     */
    readonly ConnectorType: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connectorprofile.html#cfn-appflow-connectorprofile-kmsarn
     */
    readonly KMSArn?: pulumi.Input<string>;
}
