// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.
 */
export class SlackChannelConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing SlackChannelConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SlackChannelConfiguration {
        return new SlackChannelConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:supportapp:SlackChannelConfiguration';

    /**
     * Returns true if the given object is an instance of SlackChannelConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SlackChannelConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SlackChannelConfiguration.__pulumiType;
    }

    /**
     * The channel ID in Slack, which identifies a channel within a workspace.
     */
    public readonly channelId!: pulumi.Output<string>;
    /**
     * The channel name in Slack.
     */
    public readonly channelName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
     */
    public readonly channelRoleArn!: pulumi.Output<string>;
    /**
     * Whether to notify when a correspondence is added to a case.
     */
    public readonly notifyOnAddCorrespondenceToCase!: pulumi.Output<boolean | undefined>;
    /**
     * The severity level of a support case that a customer wants to get notified for.
     */
    public readonly notifyOnCaseSeverity!: pulumi.Output<enums.supportapp.SlackChannelConfigurationNotifyOnCaseSeverity>;
    /**
     * Whether to notify when a case is created or reopened.
     */
    public readonly notifyOnCreateOrReopenCase!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to notify when a case is resolved.
     */
    public readonly notifyOnResolveCase!: pulumi.Output<boolean | undefined>;
    /**
     * The team ID in Slack, which uniquely identifies a workspace.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a SlackChannelConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SlackChannelConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.channelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channelId'");
            }
            if ((!args || args.channelRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channelRoleArn'");
            }
            if ((!args || args.notifyOnCaseSeverity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notifyOnCaseSeverity'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["channelId"] = args ? args.channelId : undefined;
            resourceInputs["channelName"] = args ? args.channelName : undefined;
            resourceInputs["channelRoleArn"] = args ? args.channelRoleArn : undefined;
            resourceInputs["notifyOnAddCorrespondenceToCase"] = args ? args.notifyOnAddCorrespondenceToCase : undefined;
            resourceInputs["notifyOnCaseSeverity"] = args ? args.notifyOnCaseSeverity : undefined;
            resourceInputs["notifyOnCreateOrReopenCase"] = args ? args.notifyOnCreateOrReopenCase : undefined;
            resourceInputs["notifyOnResolveCase"] = args ? args.notifyOnResolveCase : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        } else {
            resourceInputs["channelId"] = undefined /*out*/;
            resourceInputs["channelName"] = undefined /*out*/;
            resourceInputs["channelRoleArn"] = undefined /*out*/;
            resourceInputs["notifyOnAddCorrespondenceToCase"] = undefined /*out*/;
            resourceInputs["notifyOnCaseSeverity"] = undefined /*out*/;
            resourceInputs["notifyOnCreateOrReopenCase"] = undefined /*out*/;
            resourceInputs["notifyOnResolveCase"] = undefined /*out*/;
            resourceInputs["teamId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SlackChannelConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SlackChannelConfiguration resource.
 */
export interface SlackChannelConfigurationArgs {
    /**
     * The channel ID in Slack, which identifies a channel within a workspace.
     */
    channelId: pulumi.Input<string>;
    /**
     * The channel name in Slack.
     */
    channelName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
     */
    channelRoleArn: pulumi.Input<string>;
    /**
     * Whether to notify when a correspondence is added to a case.
     */
    notifyOnAddCorrespondenceToCase?: pulumi.Input<boolean>;
    /**
     * The severity level of a support case that a customer wants to get notified for.
     */
    notifyOnCaseSeverity: pulumi.Input<enums.supportapp.SlackChannelConfigurationNotifyOnCaseSeverity>;
    /**
     * Whether to notify when a case is created or reopened.
     */
    notifyOnCreateOrReopenCase?: pulumi.Input<boolean>;
    /**
     * Whether to notify when a case is resolved.
     */
    notifyOnResolveCase?: pulumi.Input<boolean>;
    /**
     * The team ID in Slack, which uniquely identifies a workspace.
     */
    teamId: pulumi.Input<string>;
}