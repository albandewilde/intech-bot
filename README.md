# intech-bot

Discord bot to help managing the school guild

## Start the project

To start the project you must use the `Makefile`.

Run:
- `make test` to run unit tests.
- `make build` to build the project. It will produce the `intech-bot` binary.
- `make run` to run the project.
- `make ctn-build` to build an image for the project named `intech-bot`.
- `make ctn-run` to run the project in a container.

## Envoronment variable

- `TKN` → The discord bot token
- `HOST` → Host the metric server listen (default is `0.0.0.0`)
- `PORT` → Port the metric server listen (defautl is `5419`)

## Deployment

The deployment is made with kubernetes.

All files are un the `./k8s/` directory.

The deployment depends on a secret named `intech-bot-discord` that have the `tkn` key.
