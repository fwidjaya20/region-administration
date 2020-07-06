# Regional Administration
Regional Administration of Indonesia. 
Data based on [Permendagri No. 37 Tahun 2017](https://www.kemendagri.go.id/page/read/40/permendagri-no137-tahun-2017).

#### Prerequisites
- Go

#### Domain
- Province
- Regency
- District
- Village

## Getting Started
#### How To Serve
```shell script
go run ./cmd/app
```
or
```shell script
make run_dev
```

## Tool
#### Make
1. Migration
```shell script
make migration version=YOUR_VERSION_NUMBER filename=YOUR_FILE_NAME
```
2. Domain
```shell script
make domain name=YOUR_DOMAIN_NAME
```