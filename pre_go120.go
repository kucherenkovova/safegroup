// Copyright 2023 The Go Authors. All rights reserved.
// Copyright 2024 Volodymyr Kucherenko <kucherenkovova@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.20

package safegroup

import "context"

func withCancelCause(parent context.Context) (context.Context, func(error)) {
	ctx, cancel := context.WithCancel(parent)
	return ctx, func(error) { cancel() }
}
