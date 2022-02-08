// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::NetworkAclEntry
 */
export function getNetworkAclEntry(args: GetNetworkAclEntryArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkAclEntryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getNetworkAclEntry", {
        "id": args.id,
    }, opts);
}

export interface GetNetworkAclEntryArgs {
    id: string;
}

export interface GetNetworkAclEntryResult {
    /**
     * The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). Requirement is conditional: You must specify the CidrBlock or Ipv6CidrBlock property
     */
    readonly cidrBlock?: string;
    /**
     * The Internet Control Message Protocol (ICMP) code and type. Requirement is conditional: Required if specifying 1 (ICMP) for the protocol parameter
     */
    readonly icmp?: outputs.ec2.NetworkAclEntryIcmp;
    readonly id?: string;
    /**
     * The IPv6 network range to allow or deny, in CIDR notation (for example 2001:db8:1234:1a00::/64)
     */
    readonly ipv6CidrBlock?: string;
    /**
     * The IPv4 network range to allow or deny, in CIDR notation (for example 172.16.0.0/24). We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18
     */
    readonly portRange?: outputs.ec2.NetworkAclEntryPortRange;
    /**
     * The protocol number. A value of "-1" means all protocols. If you specify "-1" or a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on all ports is allowed, regardless of any ports or ICMP types or codes that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR block, traffic for all ICMP types and codes allowed, regardless of any that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR block, you must specify an ICMP type and code
     */
    readonly protocol?: number;
    /**
     * Indicates whether to allow or deny the traffic that matches the rule
     */
    readonly ruleAction?: string;
}

export function getNetworkAclEntryOutput(args: GetNetworkAclEntryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkAclEntryResult> {
    return pulumi.output(args).apply(a => getNetworkAclEntry(a, opts))
}

export interface GetNetworkAclEntryOutputArgs {
    id: pulumi.Input<string>;
}
