// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a launch profile which delegates access to a collection of studio components to studio users
 */
export function getLaunchProfile(args: GetLaunchProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchProfileResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:nimblestudio:getLaunchProfile", {
        "launchProfileId": args.launchProfileId,
        "studioId": args.studioId,
    }, opts);
}

export interface GetLaunchProfileArgs {
    launchProfileId: string;
    /**
     * <p>The studio ID. </p>
     */
    studioId: string;
}

export interface GetLaunchProfileResult {
    /**
     * <p>The description.</p>
     */
    readonly description?: string;
    readonly launchProfileId?: string;
    /**
     * <p>The version number of the protocol that is used by the launch profile. The only valid
     *             version is "2021-03-31".</p>
     */
    readonly launchProfileProtocolVersions?: string[];
    /**
     * <p>The name for the launch profile.</p>
     */
    readonly name?: string;
    readonly streamConfiguration?: outputs.nimblestudio.LaunchProfileStreamConfiguration;
    /**
     * <p>Unique identifiers for a collection of studio components that can be used with this
     *             launch profile.</p>
     */
    readonly studioComponentIds?: string[];
}

export function getLaunchProfileOutput(args: GetLaunchProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLaunchProfileResult> {
    return pulumi.output(args).apply(a => getLaunchProfile(a, opts))
}

export interface GetLaunchProfileOutputArgs {
    launchProfileId: pulumi.Input<string>;
    /**
     * <p>The studio ID. </p>
     */
    studioId: pulumi.Input<string>;
}
