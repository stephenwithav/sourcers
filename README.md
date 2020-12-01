# sourcers

Adds simple, additional `Sourcers` to [ardanlabs/conf](https://github.com/ardanlabs/conf/).


## Rationale

You like using `conf` and don't want to switch to [spf13/viper](https://github.com/spf13/viper), but your command line or list of environment variables is getting long.  You wish you could simply drop a `config.json`, `config.yaml`, or `config.toml` in your container to simplify your life.

Or, you want to write your own [Sourcer](https://github.com/ardanlabs/conf/blob/master/conf.go#L28-L33) and aren't quite sure how.

This package is your answer.
