# uptime
Availability tool for distributed services. A toy project. Includes support for mongodb and any http end point (which works for rabbitmq, elasticsearch)

# usage
Path to config file required

```bash
$ uptime path/to/config
```

Example of a valid `config.yaml`
```yaml
services: # a list of services
  - name: service name
    type: http|mongodb|redis
    method: http method (defaults to GET)
    url: fully qualified url
    expect: server response status considered successful (only applies to type http)
    auth: basic auth credentials in format user:pwd (only applies to type http)
    pwd: password (only applies to type redis)
```

# todos
- [x] http method should default to GET
- [x] status
- [x] concurrent service calls
- [x] http connection timeout
- [x] http request timeout
- [ ] cron + storage

# license
MIT
