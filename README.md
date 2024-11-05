# Serverless Inferencing Host

This will allow incoming inferencing jobs in the OpenAI format and push / 
expose them to other services (kserve / Acaisia)

## Generating new OpenAI API spec files

run `go generate ./...`


# Running the application

`go run .`

## Log settings

Example debug log level for serve system only:
```shell
GOLOG_LOG_LEVEL="info,serve=debug" go run .
```

Some other common env vars:
- `GOLOG_FILE` will write to file instead of stderr
- `GOLOG_OUTPUT="stderr+file"` will write to file and stderr
- `GOLOG_LOG_FMT="json"` will set output to JSON (other options are `color` or `nocolor`)
- `GOLOG_LOG_LABELS="app=example_app,dc=sjc-1"` will add `{"app": "example_app", "dc": "sjc-1"}` to every log entry.
