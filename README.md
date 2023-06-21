# instance-api

** WARNING: This repo is a work in progress in it's infant stages **

Instance API is a service responsible for managing instances in the infratographer eco-system.

## To Do

- Authentication
- Authorization
- mutation APIs
- Instance Provider Specs Validation
- Emit events on queue

## Creating Instances

Instances are created with a provider configuration. The provider configuration must match the spec that is provided
when the provider is created. Instance API itself is **not** responsible for handling the provisioning of the instance
beyond emitting the event on the NATs queue that the event was created (*not yet implimented*). Providers should listen
to the NATS queue for events and then create and delete instances as needed. This allows you to create and manage
instances though the Instance API service no matter the backend where they are running.

## Future Things

In the future infratographer will provide an instance provider that can manage instances that are provided by server api.
This will make provisioning bare-metal servers possible. Additional things that an instance provider could be built for
are self hosted providers like  Proxmox and VMware vSphere. As well as dedicated server providers and cloud providers.
