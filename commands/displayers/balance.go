/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package displayers

import (
	"io"

	"github.com/digitalocean/doctl/do"
)

type Balance struct {
	*do.Balance
}

var _ Displayable = &Balance{}

func (a *Balance) JSON(out io.Writer) error {
	return writeJSON(a.Balance, out)
}

func (a *Balance) Cols() []string {
	return []string{
		"MonthToDateBalance", "AccountBalance", "MonthToDateUsage", "GeneratedAt",
	}
}

func (a *Balance) ColMap() map[string]string {
	return map[string]string{
		"MonthToDateBalance": "Month-to-date balance",
		"AccountBalance": "Account Balance",
		"MonthToDateUsage": "Month-to-date Usage",
		"GeneratedAt": "Generated At",
	}
}

func (a *Balance) KV() []map[string]interface{} {
	out := []map[string]interface{}{}
	x := map[string]interface{}{
		"MonthToDateBalance": a.MonthToDateBalance,
		"AccountBalance": a.AccountBalance,
		"MonthToDateBalance": a.MonthToDateBalance,
		"GeneratedAt": a.GeneratedAt,
	}
	out = append(out, x)

	return out
}
