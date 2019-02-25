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
    method: http method (defaults to GET, only applies to http)
    url: fully qualified url
    expect: server response status considered successful (defaults to 200, only applies to http)
    auth: basic auth credentials in format user:pwd (only applies to http)
    pwd: password (only applies to redis)
```

# todos
- [x] http method should default to GET
- [x] status
- [x] concurrent service calls
- [x] http connection timeout
- [x] http request timeout
- [ ] cron + storage
- [x] conf defaults

# license
MIT
