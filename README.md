## Description
This module provides reading config for you app. The module was created in the 8th homework assignment at the Tinkoff Backend Academy.

## Last Version
Last version is `1.0.0`

## How to install
```
go get github.com/JustNik8/cfgreader
```

## Usage
Declare struct, that describe your app configuration. You should use tag `yaml` to declare key in yaml file. For example:
``` go
type Config struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
```
This structure describes the following yaml file:
``` yaml
host: localhost
port: 4000
timeout: 10s
```

Then declare variable config and call `cfgreader.ReadCfg(filename, cfg)`. 
``` go
var cfg = &Config{}
err := cfgreader.ReadCfg("./cfg.yaml", cfg)
```

1 param: filename string - path of yaml file

2 param: cfg any - pass pointer of config struct
