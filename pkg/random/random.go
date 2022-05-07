/**
 * Copyright (C) 2022 Yeno
 *
 * This file is part of Emotebox.
 *
 * Emotebox is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Emotebox is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Emotebox.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package random provides functions focused on randomness
package random

import (
	"crypto/rand"
)

// RandBytes generates random bytes
func RandBytes(len int) ([]byte, error) {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// RandString generates a random string
func RandString(len int) (string, error) {
	b, err := RandBytes(len)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}
