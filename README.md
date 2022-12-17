# uniplay

## Local development
1. Run service deps
```sh
$ make compose-up
```
2. Run the application
```sh
$ make run-migrate
```

## Error codes

App error codes unique for any specific error and can be only >= 600.

Codes per domain model:
`Match` - >= 600
`Metric` - >= 700
`Player` - >= 800
`Team` - >= 900
`WeaponStats` - >= 1000