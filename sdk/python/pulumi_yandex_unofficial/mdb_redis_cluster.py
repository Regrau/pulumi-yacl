# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MdbRedisClusterArgs', 'MdbRedisCluster']

@pulumi.input_type
class MdbRedisClusterArgs:
    def __init__(__self__, *,
                 config: pulumi.Input['MdbRedisClusterConfigArgs'],
                 environment: pulumi.Input[str],
                 hosts: pulumi.Input[Sequence[pulumi.Input['MdbRedisClusterHostArgs']]],
                 network_id: pulumi.Input[str],
                 resources: pulumi.Input['MdbRedisClusterResourcesArgs'],
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 maintenance_window: Optional[pulumi.Input['MdbRedisClusterMaintenanceWindowArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistence_mode: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sharded: Optional[pulumi.Input[bool]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a MdbRedisCluster resource.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "hosts", hosts)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "resources", resources)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistence_mode is not None:
            pulumi.set(__self__, "persistence_mode", persistence_mode)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if sharded is not None:
            pulumi.set(__self__, "sharded", sharded)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input['MdbRedisClusterConfigArgs']:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input['MdbRedisClusterConfigArgs']):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Input[Sequence[pulumi.Input['MdbRedisClusterHostArgs']]]:
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: pulumi.Input[Sequence[pulumi.Input['MdbRedisClusterHostArgs']]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Input['MdbRedisClusterResourcesArgs']:
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: pulumi.Input['MdbRedisClusterResourcesArgs']):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['MdbRedisClusterMaintenanceWindowArgs']]:
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['MdbRedisClusterMaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="persistenceMode")
    def persistence_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "persistence_mode")

    @persistence_mode.setter
    def persistence_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "persistence_mode", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter
    def sharded(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sharded")

    @sharded.setter
    def sharded(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sharded", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)


@pulumi.input_type
class _MdbRedisClusterState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input['MdbRedisClusterConfigArgs']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 health: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input['MdbRedisClusterHostArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 maintenance_window: Optional[pulumi.Input['MdbRedisClusterMaintenanceWindowArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 persistence_mode: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input['MdbRedisClusterResourcesArgs']] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sharded: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering MdbRedisCluster resources.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if health is not None:
            pulumi.set(__self__, "health", health)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if persistence_mode is not None:
            pulumi.set(__self__, "persistence_mode", persistence_mode)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if sharded is not None:
            pulumi.set(__self__, "sharded", sharded)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['MdbRedisClusterConfigArgs']]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['MdbRedisClusterConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def health(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health")

    @health.setter
    def health(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health", value)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MdbRedisClusterHostArgs']]]]:
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MdbRedisClusterHostArgs']]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['MdbRedisClusterMaintenanceWindowArgs']]:
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['MdbRedisClusterMaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="persistenceMode")
    def persistence_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "persistence_mode")

    @persistence_mode.setter
    def persistence_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "persistence_mode", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input['MdbRedisClusterResourcesArgs']]:
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input['MdbRedisClusterResourcesArgs']]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter
    def sharded(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sharded")

    @sharded.setter
    def sharded(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sharded", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)


class MdbRedisCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterConfigArgs']]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MdbRedisClusterHostArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterMaintenanceWindowArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 persistence_mode: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterResourcesArgs']]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sharded: Optional[pulumi.Input[bool]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a MdbRedisCluster resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MdbRedisClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a MdbRedisCluster resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MdbRedisClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MdbRedisClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterConfigArgs']]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MdbRedisClusterHostArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterMaintenanceWindowArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 persistence_mode: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterResourcesArgs']]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sharded: Optional[pulumi.Input[bool]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MdbRedisClusterArgs.__new__(MdbRedisClusterArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            __props__.__dict__["deletion_protection"] = deletion_protection
            __props__.__dict__["description"] = description
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            __props__.__dict__["folder_id"] = folder_id
            if hosts is None and not opts.urn:
                raise TypeError("Missing required property 'hosts'")
            __props__.__dict__["hosts"] = hosts
            __props__.__dict__["labels"] = labels
            __props__.__dict__["maintenance_window"] = maintenance_window
            __props__.__dict__["name"] = name
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["persistence_mode"] = persistence_mode
            if resources is None and not opts.urn:
                raise TypeError("Missing required property 'resources'")
            __props__.__dict__["resources"] = resources
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["sharded"] = sharded
            __props__.__dict__["tls_enabled"] = tls_enabled
            __props__.__dict__["created_at"] = None
            __props__.__dict__["health"] = None
            __props__.__dict__["status"] = None
        super(MdbRedisCluster, __self__).__init__(
            'yandex:index/mdbRedisCluster:MdbRedisCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterConfigArgs']]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            health: Optional[pulumi.Input[str]] = None,
            hosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MdbRedisClusterHostArgs']]]]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            maintenance_window: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterMaintenanceWindowArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            persistence_mode: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[pulumi.InputType['MdbRedisClusterResourcesArgs']]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            sharded: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tls_enabled: Optional[pulumi.Input[bool]] = None) -> 'MdbRedisCluster':
        """
        Get an existing MdbRedisCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MdbRedisClusterState.__new__(_MdbRedisClusterState)

        __props__.__dict__["config"] = config
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["deletion_protection"] = deletion_protection
        __props__.__dict__["description"] = description
        __props__.__dict__["environment"] = environment
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["health"] = health
        __props__.__dict__["hosts"] = hosts
        __props__.__dict__["labels"] = labels
        __props__.__dict__["maintenance_window"] = maintenance_window
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["persistence_mode"] = persistence_mode
        __props__.__dict__["resources"] = resources
        __props__.__dict__["security_group_ids"] = security_group_ids
        __props__.__dict__["sharded"] = sharded
        __props__.__dict__["status"] = status
        __props__.__dict__["tls_enabled"] = tls_enabled
        return MdbRedisCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.MdbRedisClusterConfig']:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def health(self) -> pulumi.Output[str]:
        return pulumi.get(self, "health")

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Output[Sequence['outputs.MdbRedisClusterHost']]:
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> pulumi.Output['outputs.MdbRedisClusterMaintenanceWindow']:
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="persistenceMode")
    def persistence_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "persistence_mode")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output['outputs.MdbRedisClusterResources']:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def sharded(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "sharded")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "tls_enabled")
