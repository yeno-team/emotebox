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

package example

import "fmt"

type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println(message)
}

func (l *Logger) Logf(format string, values ...interface{}) {
	fmt.Printf(format, values...)
}

func (l *Logger) Warn(message string) {
	l.Log(message)
}

func (l *Logger) Warnf(format string, values ...interface{}) {
	l.Logf(format, values...)
}

func (l *Logger) Fatal(message string) {
	l.Log(message)
}

func (l *Logger) Fatalf(format string, values ...interface{}) {
	l.Logf(format, values...)
}
