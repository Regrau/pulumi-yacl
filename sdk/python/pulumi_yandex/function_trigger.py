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

__all__ = ['FunctionTriggerArgs', 'FunctionTrigger']

@pulumi.input_type
class FunctionTriggerArgs:
    def __init__(__self__, *,
                 function: pulumi.Input['FunctionTriggerFunctionArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 dlq: Optional[pulumi.Input['FunctionTriggerDlqArgs']] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 iot: Optional[pulumi.Input['FunctionTriggerIotArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 log_group: Optional[pulumi.Input['FunctionTriggerLogGroupArgs']] = None,
                 logging: Optional[pulumi.Input['FunctionTriggerLoggingArgs']] = None,
                 message_queue: Optional[pulumi.Input['FunctionTriggerMessageQueueArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_storage: Optional[pulumi.Input['FunctionTriggerObjectStorageArgs']] = None,
                 timer: Optional[pulumi.Input['FunctionTriggerTimerArgs']] = None):
        """
        The set of arguments for constructing a FunctionTrigger resource.
        """
        pulumi.set(__self__, "function", function)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dlq is not None:
            pulumi.set(__self__, "dlq", dlq)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if iot is not None:
            pulumi.set(__self__, "iot", iot)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if log_group is not None:
            pulumi.set(__self__, "log_group", log_group)
        if logging is not None:
            pulumi.set(__self__, "logging", logging)
        if message_queue is not None:
            pulumi.set(__self__, "message_queue", message_queue)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_storage is not None:
            pulumi.set(__self__, "object_storage", object_storage)
        if timer is not None:
            pulumi.set(__self__, "timer", timer)

    @property
    @pulumi.getter
    def function(self) -> pulumi.Input['FunctionTriggerFunctionArgs']:
        return pulumi.get(self, "function")

    @function.setter
    def function(self, value: pulumi.Input['FunctionTriggerFunctionArgs']):
        pulumi.set(self, "function", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def dlq(self) -> Optional[pulumi.Input['FunctionTriggerDlqArgs']]:
        return pulumi.get(self, "dlq")

    @dlq.setter
    def dlq(self, value: Optional[pulumi.Input['FunctionTriggerDlqArgs']]):
        pulumi.set(self, "dlq", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def iot(self) -> Optional[pulumi.Input['FunctionTriggerIotArgs']]:
        return pulumi.get(self, "iot")

    @iot.setter
    def iot(self, value: Optional[pulumi.Input['FunctionTriggerIotArgs']]):
        pulumi.set(self, "iot", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[pulumi.Input['FunctionTriggerLogGroupArgs']]:
        return pulumi.get(self, "log_group")

    @log_group.setter
    def log_group(self, value: Optional[pulumi.Input['FunctionTriggerLogGroupArgs']]):
        pulumi.set(self, "log_group", value)

    @property
    @pulumi.getter
    def logging(self) -> Optional[pulumi.Input['FunctionTriggerLoggingArgs']]:
        return pulumi.get(self, "logging")

    @logging.setter
    def logging(self, value: Optional[pulumi.Input['FunctionTriggerLoggingArgs']]):
        pulumi.set(self, "logging", value)

    @property
    @pulumi.getter(name="messageQueue")
    def message_queue(self) -> Optional[pulumi.Input['FunctionTriggerMessageQueueArgs']]:
        return pulumi.get(self, "message_queue")

    @message_queue.setter
    def message_queue(self, value: Optional[pulumi.Input['FunctionTriggerMessageQueueArgs']]):
        pulumi.set(self, "message_queue", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectStorage")
    def object_storage(self) -> Optional[pulumi.Input['FunctionTriggerObjectStorageArgs']]:
        return pulumi.get(self, "object_storage")

    @object_storage.setter
    def object_storage(self, value: Optional[pulumi.Input['FunctionTriggerObjectStorageArgs']]):
        pulumi.set(self, "object_storage", value)

    @property
    @pulumi.getter
    def timer(self) -> Optional[pulumi.Input['FunctionTriggerTimerArgs']]:
        return pulumi.get(self, "timer")

    @timer.setter
    def timer(self, value: Optional[pulumi.Input['FunctionTriggerTimerArgs']]):
        pulumi.set(self, "timer", value)


@pulumi.input_type
class _FunctionTriggerState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dlq: Optional[pulumi.Input['FunctionTriggerDlqArgs']] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 function: Optional[pulumi.Input['FunctionTriggerFunctionArgs']] = None,
                 iot: Optional[pulumi.Input['FunctionTriggerIotArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 log_group: Optional[pulumi.Input['FunctionTriggerLogGroupArgs']] = None,
                 logging: Optional[pulumi.Input['FunctionTriggerLoggingArgs']] = None,
                 message_queue: Optional[pulumi.Input['FunctionTriggerMessageQueueArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_storage: Optional[pulumi.Input['FunctionTriggerObjectStorageArgs']] = None,
                 timer: Optional[pulumi.Input['FunctionTriggerTimerArgs']] = None):
        """
        Input properties used for looking up and filtering FunctionTrigger resources.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dlq is not None:
            pulumi.set(__self__, "dlq", dlq)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if function is not None:
            pulumi.set(__self__, "function", function)
        if iot is not None:
            pulumi.set(__self__, "iot", iot)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if log_group is not None:
            pulumi.set(__self__, "log_group", log_group)
        if logging is not None:
            pulumi.set(__self__, "logging", logging)
        if message_queue is not None:
            pulumi.set(__self__, "message_queue", message_queue)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_storage is not None:
            pulumi.set(__self__, "object_storage", object_storage)
        if timer is not None:
            pulumi.set(__self__, "timer", timer)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def dlq(self) -> Optional[pulumi.Input['FunctionTriggerDlqArgs']]:
        return pulumi.get(self, "dlq")

    @dlq.setter
    def dlq(self, value: Optional[pulumi.Input['FunctionTriggerDlqArgs']]):
        pulumi.set(self, "dlq", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def function(self) -> Optional[pulumi.Input['FunctionTriggerFunctionArgs']]:
        return pulumi.get(self, "function")

    @function.setter
    def function(self, value: Optional[pulumi.Input['FunctionTriggerFunctionArgs']]):
        pulumi.set(self, "function", value)

    @property
    @pulumi.getter
    def iot(self) -> Optional[pulumi.Input['FunctionTriggerIotArgs']]:
        return pulumi.get(self, "iot")

    @iot.setter
    def iot(self, value: Optional[pulumi.Input['FunctionTriggerIotArgs']]):
        pulumi.set(self, "iot", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[pulumi.Input['FunctionTriggerLogGroupArgs']]:
        return pulumi.get(self, "log_group")

    @log_group.setter
    def log_group(self, value: Optional[pulumi.Input['FunctionTriggerLogGroupArgs']]):
        pulumi.set(self, "log_group", value)

    @property
    @pulumi.getter
    def logging(self) -> Optional[pulumi.Input['FunctionTriggerLoggingArgs']]:
        return pulumi.get(self, "logging")

    @logging.setter
    def logging(self, value: Optional[pulumi.Input['FunctionTriggerLoggingArgs']]):
        pulumi.set(self, "logging", value)

    @property
    @pulumi.getter(name="messageQueue")
    def message_queue(self) -> Optional[pulumi.Input['FunctionTriggerMessageQueueArgs']]:
        return pulumi.get(self, "message_queue")

    @message_queue.setter
    def message_queue(self, value: Optional[pulumi.Input['FunctionTriggerMessageQueueArgs']]):
        pulumi.set(self, "message_queue", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectStorage")
    def object_storage(self) -> Optional[pulumi.Input['FunctionTriggerObjectStorageArgs']]:
        return pulumi.get(self, "object_storage")

    @object_storage.setter
    def object_storage(self, value: Optional[pulumi.Input['FunctionTriggerObjectStorageArgs']]):
        pulumi.set(self, "object_storage", value)

    @property
    @pulumi.getter
    def timer(self) -> Optional[pulumi.Input['FunctionTriggerTimerArgs']]:
        return pulumi.get(self, "timer")

    @timer.setter
    def timer(self, value: Optional[pulumi.Input['FunctionTriggerTimerArgs']]):
        pulumi.set(self, "timer", value)


class FunctionTrigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dlq: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerDlqArgs']]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 function: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerFunctionArgs']]] = None,
                 iot: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerIotArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 log_group: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerLogGroupArgs']]] = None,
                 logging: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerLoggingArgs']]] = None,
                 message_queue: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerMessageQueueArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_storage: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerObjectStorageArgs']]] = None,
                 timer: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerTimerArgs']]] = None,
                 __props__=None):
        """
        Create a FunctionTrigger resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionTriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FunctionTrigger resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FunctionTriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionTriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dlq: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerDlqArgs']]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 function: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerFunctionArgs']]] = None,
                 iot: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerIotArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 log_group: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerLogGroupArgs']]] = None,
                 logging: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerLoggingArgs']]] = None,
                 message_queue: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerMessageQueueArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_storage: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerObjectStorageArgs']]] = None,
                 timer: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerTimerArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionTriggerArgs.__new__(FunctionTriggerArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dlq"] = dlq
            __props__.__dict__["folder_id"] = folder_id
            if function is None and not opts.urn:
                raise TypeError("Missing required property 'function'")
            __props__.__dict__["function"] = function
            __props__.__dict__["iot"] = iot
            __props__.__dict__["labels"] = labels
            __props__.__dict__["log_group"] = log_group
            __props__.__dict__["logging"] = logging
            __props__.__dict__["message_queue"] = message_queue
            __props__.__dict__["name"] = name
            __props__.__dict__["object_storage"] = object_storage
            __props__.__dict__["timer"] = timer
            __props__.__dict__["created_at"] = None
        super(FunctionTrigger, __self__).__init__(
            'yandex:index/functionTrigger:FunctionTrigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dlq: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerDlqArgs']]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            function: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerFunctionArgs']]] = None,
            iot: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerIotArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            log_group: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerLogGroupArgs']]] = None,
            logging: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerLoggingArgs']]] = None,
            message_queue: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerMessageQueueArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object_storage: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerObjectStorageArgs']]] = None,
            timer: Optional[pulumi.Input[pulumi.InputType['FunctionTriggerTimerArgs']]] = None) -> 'FunctionTrigger':
        """
        Get an existing FunctionTrigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionTriggerState.__new__(_FunctionTriggerState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["dlq"] = dlq
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["function"] = function
        __props__.__dict__["iot"] = iot
        __props__.__dict__["labels"] = labels
        __props__.__dict__["log_group"] = log_group
        __props__.__dict__["logging"] = logging
        __props__.__dict__["message_queue"] = message_queue
        __props__.__dict__["name"] = name
        __props__.__dict__["object_storage"] = object_storage
        __props__.__dict__["timer"] = timer
        return FunctionTrigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dlq(self) -> pulumi.Output[Optional['outputs.FunctionTriggerDlq']]:
        return pulumi.get(self, "dlq")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def function(self) -> pulumi.Output['outputs.FunctionTriggerFunction']:
        return pulumi.get(self, "function")

    @property
    @pulumi.getter
    def iot(self) -> pulumi.Output[Optional['outputs.FunctionTriggerIot']]:
        return pulumi.get(self, "iot")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> pulumi.Output[Optional['outputs.FunctionTriggerLogGroup']]:
        return pulumi.get(self, "log_group")

    @property
    @pulumi.getter
    def logging(self) -> pulumi.Output[Optional['outputs.FunctionTriggerLogging']]:
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter(name="messageQueue")
    def message_queue(self) -> pulumi.Output[Optional['outputs.FunctionTriggerMessageQueue']]:
        return pulumi.get(self, "message_queue")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectStorage")
    def object_storage(self) -> pulumi.Output[Optional['outputs.FunctionTriggerObjectStorage']]:
        return pulumi.get(self, "object_storage")

    @property
    @pulumi.getter
    def timer(self) -> pulumi.Output[Optional['outputs.FunctionTriggerTimer']]:
        return pulumi.get(self, "timer")
