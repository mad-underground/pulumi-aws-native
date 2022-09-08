// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lightsail::LoadBalancer
 */
export function getLoadBalancer(args: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lightsail:getLoadBalancer", {
        "loadBalancerName": args.loadBalancerName,
    }, opts);
}

export interface GetLoadBalancerArgs {
    /**
     * The name of your load balancer.
     */
    loadBalancerName: string;
}

export interface GetLoadBalancerResult {
    /**
     * The names of the instances attached to the load balancer.
     */
    readonly attachedInstances?: string[];
    /**
     * The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
     */
    readonly healthCheckPath?: string;
    readonly loadBalancerArn?: string;
    /**
     * Configuration option to enable session stickiness.
     */
    readonly sessionStickinessEnabled?: boolean;
    /**
     * Configuration option to adjust session stickiness cookie duration parameter.
     */
    readonly sessionStickinessLBCookieDurationSeconds?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.lightsail.LoadBalancerTag[];
    /**
     * The name of the TLS policy to apply to the load balancer.
     */
    readonly tlsPolicyName?: string;
}

export function getLoadBalancerOutput(args: GetLoadBalancerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadBalancerResult> {
    return pulumi.output(args).apply(a => getLoadBalancer(a, opts))
}

export interface GetLoadBalancerOutputArgs {
    /**
     * The name of your load balancer.
     */
    loadBalancerName: pulumi.Input<string>;
}
