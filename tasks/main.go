/*
 * The golabs/tasks application.
 *
 * Copyright (C) 2024 Alexandre Mulatinho <alex@mulatinho.net>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either  version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Package tasks provides a collection of functions and types
// for managing and executing various tasks.
//
// The package includes functionality for task creation, execution,
// and management, and is designed to be flexible and extensible.
package main

import (
	"github.com/mulatinho/golabs/tasks/cmd"
)

func main() {
	cmd.Start()
}
