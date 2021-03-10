// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html
 */
export class LoggingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing LoggingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoggingConfiguration {
        return new LoggingConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:NetworkFirewall:LoggingConfiguration';

    /**
     * Returns true if the given object is an instance of LoggingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoggingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoggingConfiguration.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallarn
     */
    public readonly FirewallArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallname
     */
    public readonly FirewallName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration
     */
    public readonly LoggingConfiguration!: pulumi.Output<outputs.NetworkFirewall.LoggingConfigurationLoggingConfiguration>;

    /**
     * Create a LoggingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoggingConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.FirewallArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'FirewallArn'");
            }
            if ((!args || args.LoggingConfiguration === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'LoggingConfiguration'");
            }
            inputs["FirewallArn"] = args ? args.FirewallArn : undefined;
            inputs["FirewallName"] = args ? args.FirewallName : undefined;
            inputs["LoggingConfiguration"] = args ? args.LoggingConfiguration : undefined;
        } else {
            inputs["FirewallArn"] = undefined /*out*/;
            inputs["FirewallName"] = undefined /*out*/;
            inputs["LoggingConfiguration"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LoggingConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoggingConfiguration resource.
 */
export interface LoggingConfigurationArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallarn
     */
    readonly FirewallArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallname
     */
    readonly FirewallName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration
     */
    readonly LoggingConfiguration: pulumi.Input<inputs.NetworkFirewall.LoggingConfigurationLoggingConfiguration>;
}
