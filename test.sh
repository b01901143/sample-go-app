cd config
go build
cd ..
go install b01901143.git/sample-go-app
sample-go-app --config-path containers.yaml 
