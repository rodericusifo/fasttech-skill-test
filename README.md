# Fasttech Skill Test

## Guidance

### Folder: Logic Test

- Step 1: Move to logic-test directory
- Step 2: Run

```bash
go mod tidy
```

### Folder: Questions

- Step 1: Just open the file, i answered the questions on those files

### Folder: Mini App

- Step 1: Check Go Version (Must go1.19+)

```bash
go version
# go version go1.19
```

- Step 2: Install Wire by running

```bash
go install github.com/google/wire/cmd/wire@latest
```

- Step 3: Install Depedencies by Running

```bash
go mod tidy
```

- Step 4: Copy `<dev|stag|prod|test>`.application.yaml.example to `<dev|stag|prod|test>`.application.yaml. `NOTE`: choose type dev

- Step 5: Running the grpc server

```bash
make start-dev
```

### Documentation: Mini App

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-e617f456-d76a-4635-9e44-f44247265b08?action=collection%2Ffork&collection-url=entityId%3D10344918-e617f456-d76a-4635-9e44-f44247265b08%26entityType%3Dcollection%26workspaceId%3D43df7931-feec-460c-8889-25210781dc3f)
