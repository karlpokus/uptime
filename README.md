# uptime
Availability tool for distributed services. A toy project.

# usage
Path to config file required

```bash
$ uptime path/to/config
```

Example of a valid `config.yaml`
```yaml
services:
  name: the service name
  type: http|mongodb
  method: http method other than GET
  url: fully qualified url
  expect: expected successful server status if type http
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
