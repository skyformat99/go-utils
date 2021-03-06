/* Copyright 2016 sky<skygangsta@hotmail.com>. All rights reserved.
 *
 * Licensed under the Apache License, version 2.0 (the "License").
 * You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package dbx

// 如果编译时间过长，不使用时可注释
import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Host         string
	Port         int
	DbName       string
	Dba          string
	Pwd          string
	MaxIdleConns int
	MaxOpenConns int
	Timeout      int
	Conn         *xorm.Engine
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (this *Postgres) Init() error {
	var (
		err error
	)

	if this.Timeout == 0 {
		this.Timeout = 10
	}

	this.Conn, err = xorm.NewEngine("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&connect_timeout=%d",
			this.Dba,
			this.Pwd,
			this.Host,
			this.Port,
			this.DbName,
			this.Timeout))
	if err == nil {
		this.Conn.SetMaxIdleConns(this.MaxIdleConns)
		this.Conn.SetMaxOpenConns(this.MaxOpenConns)
	}

	return err
}
