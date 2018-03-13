/*
   Restfool-go

   Copyright (C) 2018 Carsten Seeger

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.

   @author Carsten Seeger
   @copyright Copyright (C) 2018 Carsten Seeger
   @license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
   @link https://github.com/cseeger-epages/rest-api-go-skeleton
*/

package restfool

import (
	"fmt"
	"net/http"
	"strings"
)

func (a RestAPI) addDefaultHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if a.Conf.Cors.AllowCrossOrigin {
			w.Header().Set("Access-Control-Allow-Origin", a.Conf.Cors.AllowFrom)
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(a.Conf.Cors.CorsMethods, ","))
			/* tbd
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			*/
		}

		if a.Conf.TLS.Hsts {
			hsts := fmt.Sprintf("max-age=%d; includeSubDomains", a.Conf.TLS.HstsMaxAge)
			w.Header().Add("Strict-Transport-Security", hsts)
		}
		h.ServeHTTP(w, r)
	})
}
