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

__all__ = [
    'GetFunctionResult',
    'AwaitableGetFunctionResult',
    'get_function',
    'get_function_output',
]

@pulumi.output_type
class GetFunctionResult:
    """
    A collection of values returned by getFunction.
    """
    def __init__(__self__, connectivity=None, created_at=None, description=None, entrypoint=None, environment=None, execution_timeout=None, folder_id=None, function_id=None, id=None, image_size=None, labels=None, loggroup_id=None, memory=None, name=None, runtime=None, secrets=None, service_account_id=None, tags=None, version=None):
        if connectivity and not isinstance(connectivity, dict):
            raise TypeError("Expected argument 'connectivity' to be a dict")
        pulumi.set(__self__, "connectivity", connectivity)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if entrypoint and not isinstance(entrypoint, str):
            raise TypeError("Expected argument 'entrypoint' to be a str")
        pulumi.set(__self__, "entrypoint", entrypoint)
        if environment and not isinstance(environment, dict):
            raise TypeError("Expected argument 'environment' to be a dict")
        pulumi.set(__self__, "environment", environment)
        if execution_timeout and not isinstance(execution_timeout, str):
            raise TypeError("Expected argument 'execution_timeout' to be a str")
        pulumi.set(__self__, "execution_timeout", execution_timeout)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if function_id and not isinstance(function_id, str):
            raise TypeError("Expected argument 'function_id' to be a str")
        pulumi.set(__self__, "function_id", function_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_size and not isinstance(image_size, int):
            raise TypeError("Expected argument 'image_size' to be a int")
        pulumi.set(__self__, "image_size", image_size)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if loggroup_id and not isinstance(loggroup_id, str):
            raise TypeError("Expected argument 'loggroup_id' to be a str")
        pulumi.set(__self__, "loggroup_id", loggroup_id)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if runtime and not isinstance(runtime, str):
            raise TypeError("Expected argument 'runtime' to be a str")
        pulumi.set(__self__, "runtime", runtime)
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        pulumi.set(__self__, "secrets", secrets)
        if service_account_id and not isinstance(service_account_id, str):
            raise TypeError("Expected argument 'service_account_id' to be a str")
        pulumi.set(__self__, "service_account_id", service_account_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def connectivity(self) -> Optional['outputs.GetFunctionConnectivityResult']:
        """
        Function version connectivity. If specified the version will be attached to specified network.
        * `connectivity.0.network_id` - Network the version will have access to. It's essential to specify network with subnets in all availability zones.
        """
        return pulumi.get(self, "connectivity")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Creation timestamp of the Yandex Cloud Function
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the Yandex Cloud Function
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def entrypoint(self) -> str:
        """
        Entrypoint for Yandex Cloud Function
        """
        return pulumi.get(self, "entrypoint")

    @property
    @pulumi.getter
    def environment(self) -> Mapping[str, str]:
        """
        A set of key/value environment variables for Yandex Cloud Function
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="executionTimeout")
    def execution_timeout(self) -> str:
        """
        Execution timeout in seconds for Yandex Cloud Function
        """
        return pulumi.get(self, "execution_timeout")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[str]:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[str]:
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageSize")
    def image_size(self) -> int:
        """
        Image size for Yandex Cloud Function.
        """
        return pulumi.get(self, "image_size")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Function
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="loggroupId")
    def loggroup_id(self) -> str:
        """
        Log group ID size for Yandex Cloud Function.
        """
        return pulumi.get(self, "loggroup_id")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def runtime(self) -> str:
        """
        Runtime for Yandex Cloud Function
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter
    def secrets(self) -> Sequence['outputs.GetFunctionSecretResult']:
        """
        Secrets for Yandex Cloud Function.
        """
        return pulumi.get(self, "secrets")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> str:
        """
        Service account ID for Yandex Cloud Function
        """
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags for Yandex Cloud Function. Tag "$latest" isn't returned.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version for Yandex Cloud Function.
        """
        return pulumi.get(self, "version")


class AwaitableGetFunctionResult(GetFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionResult(
            connectivity=self.connectivity,
            created_at=self.created_at,
            description=self.description,
            entrypoint=self.entrypoint,
            environment=self.environment,
            execution_timeout=self.execution_timeout,
            folder_id=self.folder_id,
            function_id=self.function_id,
            id=self.id,
            image_size=self.image_size,
            labels=self.labels,
            loggroup_id=self.loggroup_id,
            memory=self.memory,
            name=self.name,
            runtime=self.runtime,
            secrets=self.secrets,
            service_account_id=self.service_account_id,
            tags=self.tags,
            version=self.version)


def get_function(connectivity: Optional[pulumi.InputType['GetFunctionConnectivityArgs']] = None,
                 folder_id: Optional[str] = None,
                 function_id: Optional[str] = None,
                 name: Optional[str] = None,
                 secrets: Optional[Sequence[pulumi.InputType['GetFunctionSecretArgs']]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionResult:
    """
    Get information about a Yandex Cloud Function. For more information about Yandex Cloud Functions, see
    [Yandex Cloud Functions](https://cloud.yandex.com/docs/functions/).

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_function = yandex.get_function(function_id="are1samplefunction11")
    ```

    This data source is used to define [Yandex Cloud Function](https://cloud.yandex.com/docs/functions/concepts/function) that can be used by other resources.


    :param pulumi.InputType['GetFunctionConnectivityArgs'] connectivity: Function version connectivity. If specified the version will be attached to specified network.
           * `connectivity.0.network_id` - Network the version will have access to. It's essential to specify network with subnets in all availability zones.
    :param str folder_id: Folder ID for the Yandex Cloud Function
    :param str function_id: Yandex Cloud Function id used to define function
    :param str name: Yandex Cloud Function name used to define function
    :param Sequence[pulumi.InputType['GetFunctionSecretArgs']] secrets: Secrets for Yandex Cloud Function.
    """
    __args__ = dict()
    __args__['connectivity'] = connectivity
    __args__['folderId'] = folder_id
    __args__['functionId'] = function_id
    __args__['name'] = name
    __args__['secrets'] = secrets
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getFunction:getFunction', __args__, opts=opts, typ=GetFunctionResult).value

    return AwaitableGetFunctionResult(
        connectivity=__ret__.connectivity,
        created_at=__ret__.created_at,
        description=__ret__.description,
        entrypoint=__ret__.entrypoint,
        environment=__ret__.environment,
        execution_timeout=__ret__.execution_timeout,
        folder_id=__ret__.folder_id,
        function_id=__ret__.function_id,
        id=__ret__.id,
        image_size=__ret__.image_size,
        labels=__ret__.labels,
        loggroup_id=__ret__.loggroup_id,
        memory=__ret__.memory,
        name=__ret__.name,
        runtime=__ret__.runtime,
        secrets=__ret__.secrets,
        service_account_id=__ret__.service_account_id,
        tags=__ret__.tags,
        version=__ret__.version)


@_utilities.lift_output_func(get_function)
def get_function_output(connectivity: Optional[pulumi.Input[Optional[pulumi.InputType['GetFunctionConnectivityArgs']]]] = None,
                        folder_id: Optional[pulumi.Input[Optional[str]]] = None,
                        function_id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        secrets: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetFunctionSecretArgs']]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionResult]:
    """
    Get information about a Yandex Cloud Function. For more information about Yandex Cloud Functions, see
    [Yandex Cloud Functions](https://cloud.yandex.com/docs/functions/).

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_function = yandex.get_function(function_id="are1samplefunction11")
    ```

    This data source is used to define [Yandex Cloud Function](https://cloud.yandex.com/docs/functions/concepts/function) that can be used by other resources.


    :param pulumi.InputType['GetFunctionConnectivityArgs'] connectivity: Function version connectivity. If specified the version will be attached to specified network.
           * `connectivity.0.network_id` - Network the version will have access to. It's essential to specify network with subnets in all availability zones.
    :param str folder_id: Folder ID for the Yandex Cloud Function
    :param str function_id: Yandex Cloud Function id used to define function
    :param str name: Yandex Cloud Function name used to define function
    :param Sequence[pulumi.InputType['GetFunctionSecretArgs']] secrets: Secrets for Yandex Cloud Function.
    """
    ...
