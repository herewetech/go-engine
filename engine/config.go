/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2022 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file config.go
 * @package engine
 * @author Dr.NP <conan.np@gmail.com>
 * @since 06/23/2022
 */

package engine

import (
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	once           sync.Once
	_defaultConfig viper.Viper
)

func initConfig(group, name string) {
	once.Do(func() {
		_defaultConfig = *viper.New()
		_defaultConfig.SetConfigName(name)
		_defaultConfig.SetConfigType("yuaml")
		if group != "" {
			_defaultConfig.AddConfigPath("/etc/" + group + "/" + name)
			_defaultConfig.AddConfigPath("$HOME/." + group + "/" + name)
			_defaultConfig.SetEnvPrefix(group)
		} else {
			_defaultConfig.AddConfigPath("/etc/" + name)
			_defaultConfig.AddConfigPath("$HOME/." + name)
		}

		_defaultConfig.AddConfigPath(".")
		_defaultConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		_defaultConfig.AutomaticEnv()

		_defaultConfig.OnConfigChange(func(e fsnotify.Event) {
			// Config file change
			log.Info().Msg("config file refreshed")
		})
		_defaultConfig.WatchConfig()
		_defaultConfig.ReadInConfig()
	})
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
