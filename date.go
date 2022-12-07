// Copyright 2022 RetailNext, Inc.
//
// Licensed under the BSD 3-Clause License (the "License");
// you may not use this file except in compliance with the License.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package easypost

import "time"

type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	if err := d.Time.UnmarshalJSON(data); err == nil {
		return nil
	}

	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
