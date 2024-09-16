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

package core

import (
	"log"
	"net/http/httptest"
	"testing"
)

func RestCallSimpleTest(t *testing.T) {
	SetupRouter()
	TestRecord := httptest.NewRecorder()
	restCall := RestCall("GET", URL_API_TASKS, nil)
	taskApp.gin.ServeHTTP(TestRecord, restCall.RestRequest)

	if restCall.RestError != nil {
		log.Println(restCall.RestError)
	}

	equals(t, 200, restCall.RestStatus)
}
