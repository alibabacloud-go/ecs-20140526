// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseElasticityAssuranceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PurchaseElasticityAssuranceResponseBody
	GetRequestId() *string
}

type PurchaseElasticityAssuranceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurchaseElasticityAssuranceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseElasticityAssuranceResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseElasticityAssuranceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseElasticityAssuranceResponseBody) SetRequestId(v string) *PurchaseElasticityAssuranceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseElasticityAssuranceResponseBody) Validate() error {
	return dara.Validate(s)
}
