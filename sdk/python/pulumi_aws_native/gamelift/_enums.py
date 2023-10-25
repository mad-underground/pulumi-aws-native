# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AliasRoutingStrategyType',
    'BuildOperatingSystem',
    'FleetCertificateConfigurationCertificateType',
    'FleetComputeType',
    'FleetInstanceRoleCredentialsProvider',
    'FleetIpPermissionProtocol',
    'FleetNewGameSessionProtectionPolicy',
    'FleetType',
    'GameServerGroupBalancingStrategy',
    'GameServerGroupDeleteOption',
    'GameServerGroupGameServerProtectionPolicy',
]


class AliasRoutingStrategyType(str, Enum):
    """
    Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.
    """
    SIMPLE = "SIMPLE"
    TERMINAL = "TERMINAL"


class BuildOperatingSystem(str, Enum):
    """
    The operating system that the game server binaries are built to run on. This value determines the type of fleet resources that you can use for this build. If your game build contains multiple executables, they all must run on the same operating system. If an operating system is not specified when creating a build, Amazon GameLift uses the default value (WINDOWS_2012). This value cannot be changed later.
    """
    AMAZON_LINUX = "AMAZON_LINUX"
    AMAZON_LINUX2 = "AMAZON_LINUX_2"
    AMAZON_LINUX2023 = "AMAZON_LINUX_2023"
    WINDOWS2012 = "WINDOWS_2012"
    WINDOWS2016 = "WINDOWS_2016"


class FleetCertificateConfigurationCertificateType(str, Enum):
    DISABLED = "DISABLED"
    GENERATED = "GENERATED"


class FleetComputeType(str, Enum):
    """
    ComputeType to differentiate EC2 hardware managed by GameLift and Anywhere hardware managed by the customer.
    """
    EC2 = "EC2"
    ANYWHERE = "ANYWHERE"


class FleetInstanceRoleCredentialsProvider(str, Enum):
    """
    Credentials provider implementation that loads credentials from the Amazon EC2 Instance Metadata Service.
    """
    SHARED_CREDENTIAL_FILE = "SHARED_CREDENTIAL_FILE"


class FleetIpPermissionProtocol(str, Enum):
    """
    The network communication protocol used by the fleet.
    """
    TCP = "TCP"
    UDP = "UDP"


class FleetNewGameSessionProtectionPolicy(str, Enum):
    """
    A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.
    """
    FULL_PROTECTION = "FullProtection"
    NO_PROTECTION = "NoProtection"


class FleetType(str, Enum):
    """
    Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.
    """
    ON_DEMAND = "ON_DEMAND"
    SPOT = "SPOT"


class GameServerGroupBalancingStrategy(str, Enum):
    """
    The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.
    """
    SPOT_ONLY = "SPOT_ONLY"
    SPOT_PREFERRED = "SPOT_PREFERRED"
    ON_DEMAND_ONLY = "ON_DEMAND_ONLY"


class GameServerGroupDeleteOption(str, Enum):
    """
    The type of delete to perform.
    """
    SAFE_DELETE = "SAFE_DELETE"
    FORCE_DELETE = "FORCE_DELETE"
    RETAIN = "RETAIN"


class GameServerGroupGameServerProtectionPolicy(str, Enum):
    """
    A flag that indicates whether instances in the game server group are protected from early termination.
    """
    NO_PROTECTION = "NO_PROTECTION"
    FULL_PROTECTION = "FULL_PROTECTION"
