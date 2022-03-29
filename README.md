# Nyckel

Nyckel is the swedish word for key. And now it's a neat utility to work with
Kubernets opaque secret files. It let's you add, edit, and delete data in
opaque secret files.

Kubectl already offers a useful interface for injecting secrets directly in
your cluster. However, that's not always what you want. So instead of doing
that, nyckel only helps you create your opaque secrets on disk so you can later
deploy them.

## Getting started

I couldn't be bothered with a Makefile just yet. So you'll need to have Go
installed on your machine to make this utility work.

Installing the tool is easy enough:

```shell
go install github.com/wmeints/nyckel/cmd/nyckel@latest
```

After installation it will be available if you have `$GOBIN` in your `$PATH`
variable.

## Supported commands

### Creating a new opaque secret file

You can create a completely new opaque secret file using the following command:

```shell
nyckel create <arguments>
```

| Argument       | Description                                                                       |
| -------------- | --------------------------------------------------------------------------------- |
| `--path`       | The path to the opaque secret file                                                |
| `--key`        | The key for the data to add                                                       |
| `--data`       | The data to be added, can be any string, optionally escaped using double quotes   |
| `--input-file` | The input file to load the secret data from. You can use this instead of `--data` |

### Adding data to an existing opaque secret file

You can add data to an existing opaque secret file using the command:

```shell
nyckel add <arguments>
```

| Argument       | Description                                                                       |
| -------------- | --------------------------------------------------------------------------------- |
| `--path`       | The path to the opaque secret file                                                |
| `--key`        | The key for the data to add                                                       |
| `--data`       | The data to be added, can be any string, optionally escaped using double quotes   |
| `--input-file` | The input file to load the secret data from. You can use this instead of `--data` |

### Editing data in an existing opaque secret file

You can edit data in an existing opaque secret file using the command:

```shell
nyckel update <arguments>
```

| Argument       | Description                                                                       |
| -------------- | --------------------------------------------------------------------------------- |
| `--path`       | The path to the opaque secret file                                                |
| `--key`        | The key for the data to add                                                       |
| `--data`       | The data to be added, can be any string, optionally escaped using double quotes   |
| `--input-file` | The input file to load the secret data from. You can use this instead of `--data` |

### Removing data from an existing opaque secret file

You can add data to an existing opaque secret file using the command:

```shell
nyckel remove <arguments>
```

| Argument | Description                        |
| -------- | ---------------------------------- |
| `--path` | The path to the opaque secret file |
| `--key`  | The key for the data to add        |
