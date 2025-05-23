// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package code

import (
	"context"

	"github.com/teamgram/teamgram-server/pkg/code/conf"
	"github.com/teamgram/teamgram-server/pkg/code/none"
)

type VerifyCodeInterface interface {
	SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string) (string, error)
	VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error
	IsTestEnvironment(ctx context.Context) bool
}

func NewVerifyCode(c *conf.SmsVerifyCodeConfig) VerifyCodeInterface {
	if c == nil {
		c = new(conf.SmsVerifyCodeConfig)
	}

	switch c.Name {
	// case "predefined":
	// 	return predefined.New(c)
	case "none":
		return none.New(c)
	}
	return none.New(c)
}
