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
 * @file app.go
 * @package engine
 * @author Dr.NP <conan.np@gmail.com>
 * @since 06/23/2022
 */

package engine

import (
	"context"

	"github.com/herewetech/go-engine/interfaces/null"
)

// App : application definition
type App struct {
	option      *Option
	mainContext context.Context
	mainCancle  context.CancelFunc
}

/* {{{ [NewApp] - Creates engine application */

// NewApp creates engine application
func NewApp(o *Option) (*App, error) {
	// logger
	initLogger()

	// Load config
	initConfig(o.Group, o.Name)

	if o.Interface == nil {
		o.Interface = null.NewInterface()
	}

	ctx, cancle := context.WithCancel(context.Background())

	app := new(App)
	app.option = o
	app.mainContext = ctx
	app.mainCancle = cancle

	return app, nil
}

/* }}} */

/* {{{ [App::*] */

// Start application
func (app *App) Start() error {
	// Start interface
	if app.option.Interface != nil {
		return app.option.Interface.Start(app.mainContext)
	}

	return nil
}

// Stop application
func (app *App) Stop() error {
	app.mainCancle()

	// Stop interface
	if app.option.Interface != nil {
		return app.option.Interface.Stop()
	}

	return nil
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
