// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Route53Resolver::ResolverRule
 */
export function getResolverRule(args: GetResolverRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetResolverRuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:route53resolver:getResolverRule", {
        "resolverRuleId": args.resolverRuleId,
    }, opts);
}

export interface GetResolverRuleArgs {
    /**
     * The ID of the endpoint that the rule is associated with.
     */
    resolverRuleId: string;
}

export interface GetResolverRuleResult {
    /**
     * The Amazon Resource Name (ARN) of the resolver rule.
     */
    readonly arn?: string;
    /**
     * DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
     */
    readonly domainName?: string;
    /**
     * The name for the Resolver rule
     */
    readonly name?: string;
    /**
     * The ID of the endpoint that the rule is associated with.
     */
    readonly resolverEndpointId?: string;
    /**
     * The ID of the endpoint that the rule is associated with.
     */
    readonly resolverRuleId?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.route53resolver.ResolverRuleTag[];
    /**
     * An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
     */
    readonly targetIps?: outputs.route53resolver.ResolverRuleTargetAddress[];
}
/**
 * Resource Type definition for AWS::Route53Resolver::ResolverRule
 */
export function getResolverRuleOutput(args: GetResolverRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResolverRuleResult> {
    return pulumi.output(args).apply((a: any) => getResolverRule(a, opts))
}

export interface GetResolverRuleOutputArgs {
    /**
     * The ID of the endpoint that the rule is associated with.
     */
    resolverRuleId: pulumi.Input<string>;
}
