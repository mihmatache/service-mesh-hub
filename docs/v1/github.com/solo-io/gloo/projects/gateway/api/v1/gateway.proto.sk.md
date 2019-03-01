
---
title: "gateway.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `gateway.solo.io` 
##### Types:


- [Gateway](#Gateway) **Top-Level Resource**
  



##### Source File: [github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto](https://github.com/solo-io/gloo/blob/master/projects/gateway/api/v1/gateway.proto)





---
### <a name="Gateway">Gateway</a>

 

A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.

```yaml
"ssl": bool
"virtual_services": []core.solo.io.ResourceRef
"bind_address": string
"bind_port": int
"plugins": .gloo.solo.io.ListenerPlugins
"status": .core.solo.io.Status
"metadata": .core.solo.io.Metadata

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `ssl` | `bool` | if set to false, only use virtual services with no ssl configured. if set to true, only use virtual services with ssl configured. |  |
| `virtual_services` | [[]core.solo.io.ResourceRef](../../../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) | names of the the virtual services, which contain the actual routes for the gateway if the list is empty, all virtual services will apply to this gateway (with accordance to tls flag above). |  |
| `bind_address` | `string` | the bind address the gateway should serve traffic on |  |
| `bind_port` | `int` | bind ports must not conflict across gateways in a namespace |  |
| `plugins` | [.gloo.solo.io.ListenerPlugins](../../../../gloo/api/v1/plugins.proto.sk#ListenerPlugins) | top level plugin configuration for all routes on the gateway |  |
| `status` | [.core.solo.io.Status](../../../../../../solo-kit/api/v1/status.proto.sk#Status) | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation |  |
| `metadata` | [.core.solo.io.Metadata](../../../../../../solo-kit/api/v1/metadata.proto.sk#Metadata) | Metadata contains the object metadata for this resource |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->