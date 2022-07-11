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
 * @file null.go
 * @package null
 * @author Dr.NP <conan.np@gmail.com>
 * @since 07/06/2022
 */

package null

import (
	"context"

	"github.com/herewetech/go-engine/interfaces"
	"github.com/rs/zerolog/log"
)

type NullInterface struct{}

func NewInterface() interfaces.Interface {
	i := new(NullInterface)

	return i
}

func (i *NullInterface) Start(ctx context.Context) error {
	log.Info().Msg("Null interface started, switch to correct interface in NewApp()")

	return nil
}

func (i *NullInterface) Stop() error {
	log.Info().Msg("See you, bye!")

	return nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
