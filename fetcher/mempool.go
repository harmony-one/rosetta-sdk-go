// Copyright 2020 Coinbase, Inc.
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

package fetcher

import (
	"context"
	"errors"

	"github.com/coinbase/rosetta-sdk-go/asserter"

	"github.com/coinbase/rosetta-sdk-go/models"
)

// UnsafeMempool returns the unvalidated response
// from the Mempool method.
func (f *Fetcher) UnsafeMempool(
	ctx context.Context,
	network *models.NetworkIdentifier,
) ([]*models.TransactionIdentifier, error) {
	mempool, _, err := f.rosettaClient.MempoolAPI.Mempool(ctx, models.MempoolRequest{
		NetworkIdentifier: network,
	})
	if err != nil {
		return nil, err
	}

	return mempool.TransactionIdentifiers, nil
}

// Mempool returns the validated response
// from the Mempool method.
func (f *Fetcher) Mempool(
	ctx context.Context,
	network *models.NetworkIdentifier,
) ([]*models.TransactionIdentifier, error) {
	mempool, err := f.UnsafeMempool(ctx, network)
	if err != nil {
		return nil, err
	}

	if err := asserter.MempoolTransactions(mempool); err != nil {
		return nil, err
	}

	return mempool, nil
}

// UnsafeMempoolTransaction returns the unvalidated response
// from the MempoolTransaction method.
func (f *Fetcher) UnsafeMempoolTransaction(
	ctx context.Context,
	network *models.NetworkIdentifier,
	transaction *models.TransactionIdentifier,
) (*models.Transaction, *map[string]interface{}, error) {
	mempoolTransaction, _, err := f.rosettaClient.MempoolAPI.MempoolTransaction(ctx,
		models.MempoolTransactionRequest{
			NetworkIdentifier:     network,
			TransactionIdentifier: transaction,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return mempoolTransaction.Transaction, mempoolTransaction.Metadata, nil
}

// MempoolTransaction returns the validated response
// from the MempoolTransaction method.
func (f *Fetcher) MempoolTransaction(
	ctx context.Context,
	network *models.NetworkIdentifier,
	transaction *models.TransactionIdentifier,
) (*models.Transaction, *map[string]interface{}, error) {
	if f.Asserter == nil {
		return nil, nil, errors.New("asserter not initialized")
	}

	mempoolTransaction, metadata, err := f.UnsafeMempoolTransaction(ctx, network, transaction)
	if err != nil {
		return nil, nil, err
	}

	if err := f.Asserter.Transaction(mempoolTransaction); err != nil {
		return nil, nil, err
	}

	return mempoolTransaction, metadata, nil
}
