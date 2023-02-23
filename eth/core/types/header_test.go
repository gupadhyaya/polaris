// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
)

var _ = Describe("Header", func() {
	var h *types.Header
	var ha common.Hash
	var sh *types.StargazerHeader

	BeforeEach(func() {
		ha = common.BytesToHash([]byte{1})
		h = &types.Header{
			Coinbase: common.BytesToAddress([]byte{2}),
			Bloom:    types.BytesToBloom([]byte{3}),
		}
		sh = types.NewStargazerHeader(h, ha)
	})

	It("should return the correct values", func() {
		Expect(sh.Author()).To(Equal(common.BytesToAddress([]byte{2})))
		// Expect(sh.Hash()).To(Equal(common.BytesToHash([]byte{1})))

		// sh.SetHash(common.Hash{})
		// Expect(sh.Hash()).To(Equal(h.Hash()))

		data, err := sh.MarshalBinary()
		Expect(err).To(BeNil())
		sh2 := types.NewEmptyStargazerHeader()
		err = sh2.UnmarshalBinary(data)
		Expect(err).To(BeNil())

		Expect(sh2.Author()).To(Equal(sh.Author()))
		Expect(sh2.Hash()).To(Equal(sh.Hash()))
		Expect(sh2.Header.Bloom).To(Equal(sh.Header.Bloom))
		Expect(sh2.Header.ReceiptHash).To(Equal(sh.Header.ReceiptHash))
	})
})