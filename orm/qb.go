// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import "errors"

type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	From(tables ...string) QueryBuilder
	Where(cond string) QueryBuilder
	LimitOffset(limit int, offset int) QueryBuilder
	InnerJoin(table string) QueryBuilder
	LeftJoin(table string) QueryBuilder
	On(cond string) QueryBuilder
	And(cond string) QueryBuilder
	Or(cond string) QueryBuilder
	In(vals ...string) QueryBuilder
	Subquery(query string, rename string) string
	String() string
}

func NewQueryBuilder(driver string) (qb QueryBuilder, err error) {
	if driver == "mysql" {
		qb = new(MySQLQueryBuilder)
	} else if driver == "postgres" {
		err = errors.New("postgres query builder is not supported yet!")
	} else if driver == "sqlite" {
		err = errors.New("sqlite query builder is not supported yet!")
	} else {
		err = errors.New("unknown driver for query builder!")
	}
	return
}