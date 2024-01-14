package node

import (
	"context"

	"log"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func (n *OracleNode) Enroll(_ context.Context, request *SendEnrollRequest) (*SendEnrollResponse, error) {
// 	//	此时接收到报名请求
// 	success := n.aggregator.Enroll(request.Enroll)
// 	return &SendEnrollResponse{EnrollSuccess: success}, nil
// }

// func (n *OracleNode) GetEnrollNodes(_ context.Context, request *SendGetEnrollNodesRequest) (*SendEnrollNodesResponse, error) {
// 	enrollNodes, success := n.aggregator.getEnrollNodes(request.GetNodes)
// 	if success {
// 		enrollNodesBytes, err := json.Marshal(enrollNodes)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &SendEnrollNodesResponse{EnrollNodes: enrollNodesBytes, EnrollSuccess: success}, nil
// 	}
// 	return &SendEnrollNodesResponse{EnrollNodes: nil, EnrollSuccess: success}, nil
// }

// 这个函数的功能是验证器来验证的过程，以及构造出应答
func (n *OracleNode) Validate(ctx context.Context, request *ValidateRequest) (*ValidateResponse, error) {

	var result *ValidateResult
	var err error

	// switch request.Type {
	// case ValidateRequest_block:
	// 	result, err = n.validator.ValidateBlock(
	// 		ctx,
	// 		common.BytesToHash(request.Hash),
	// 	)
	// case ValidateRequest_transaction:
	// 	result, err = n.validator.ValidateTransaction(
	// 		ctx,
	// 		common.BytesToHash(request.Hash),
	// 		request.Size,
	// 		request.MinRank,
	// 	)
	// }
	result, err = n.validator.ValidateTransaction(
		ctx,
		common.BytesToHash(request.Hash),
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "validate %s: %v", request.Hash, err)
	}

	resultStr := "valid"
	if !result.valid {
		resultStr = "invalid"
	}
	log.Println("Validated hash:", common.BytesToHash(request.Hash), "with result: ", resultStr)

	return ValidateResultToResponse(result), nil
}

func ValidateResultToResponse(result *ValidateResult) *ValidateResponse {
	resp := &ValidateResponse{
		Hash:      result.hash[:],
		Valid:     result.valid,
		Signature: result.signature,
		R:         result.R,
		Message:   result.message,
	}

	if result.blockNumber != nil {
		resp.BlockNumber = result.blockNumber.Int64()
	}

	return resp
}
