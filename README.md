# Go Project Template

## How to build this project

1. How to build this project

```bash
# build
make build
```

2. How to run this project

```bash
# run the bin
./dist/go-template-project

# run the version
./dist/go-template-project --version
```

## How to install this project

```bash
# install bin
cp ./dist/go-template-project /usr/local/bin/go-template

# install config file
cp ./examples/config.json /etc/go-template.config


systemctl daemon-reload
systemctl enable go-template-project
systemctl start go-template-project
systemctl status go-template-project

# uninstall
systemctl stop go-template-project
systemctl disable go-template-project
systemctl daemon-reload
```
