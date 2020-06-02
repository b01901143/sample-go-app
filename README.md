# sample-go-app
Sample go app demonstrating use of go modules.

## Flags
```
  -config-path string
    	Path to config.yaml.
  -dynamic
    	Load config.yaml to dynamic structure.
```
## Usage
Build go app
```
# build
go build

# or install
go install b01901143.git/sample-go-app
```

Load configuration yaml file into a dynamic structure.

```
sample-go-app --config-path config.yaml --dynamic 
```

Or load configuration yaml file into a static & well-defined structure.

```
sample-go-app --config-path config.yaml
```