# Custom Go handlers for AWS IoT Jobs
[AWS IoT Jobs](https://docs.aws.amazon.com/iot/latest/developerguide/jobs-what-is.html) allows 
us to define a set of remote operations that we can send to one or more devices. This is critical 
not just for admin/monitoring tasks but also for us to safely deploy new versions of the software
we run on remote devices.

The [AWS IoT Device Client](https://github.com/awslabs/aws-iot-device-client) is a device agent that
implements most of the features available in AWS IoT. It provides a [schema](https://github.com/awslabs/aws-iot-device-client/blob/main/source/jobs/README.md)
for AWS IoT Jobs and reference implementations for [AWS managed templates](https://docs.aws.amazon.com/iot/latest/developerguide/job-templates-managed.html).

New custom handlers can be implemented as a simple `bash` script. This is fine for simple operations
but for complex handlers it can get out of hand pretty quickly. For these cases, writing that logic in
Go can result in handlers that are more robust, easier to maintain and test.

This repository contains a collection of custom handlers to carry out common tasks. Each handler is
implemented as a [sub-command](https://pkg.go.dev/github.com/google/subcommands) and self-contained in
its own package. New tasks can be added by:

- creating a new package under `internal` with the command name
- creating a type that implements `subcommands.Command`
- registering that type in `main.go`