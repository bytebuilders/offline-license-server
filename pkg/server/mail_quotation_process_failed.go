/*
Copyright AppsCode Inc.

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

package server

import (
	"fmt"

	"github.com/avct/uasurfer"
	"gomodules.xyz/mailer"
	"sigs.k8s.io/yaml"
)

func NewQuotationProcessFailedMailer(gen *QuotationGenerator, err error) mailer.Mailer {
	var src string

	info := struct {
		Contact  ProductQuotation    `json:"contact"`
		UA       *uasurfer.UserAgent `json:"ua"`
		Location GeoLocation         `json:"location"`
		Err      string              `json:"error"`
	}{
		Contact:  gen.Contact,
		UA:       gen.UA,
		Location: gen.Location,
		Err:      err.Error(),
	}
	data, err := yaml.Marshal(info)
	if err != nil {
		src = fmt.Sprintf("%+v", info)
	} else {
		src = string(data)
	}

	return mailer.Mailer{
		Sender:          MailSales,
		BCC:             "",
		ReplyTo:         MailSales,
		Subject:         "[URGENT] Quotation request processing failed",
		Body:            src,
		Params:          nil,
		AttachmentBytes: nil,
	}
}
