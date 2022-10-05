package eth

import (
	"fmt"
	"math/big"
	"strings"

	apimodel "github.com/authgear/authgear-nft-indexer/pkg/api/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bunbig"

	"github.com/authgear/authgear-nft-indexer/pkg/model"
)

type NFTCollectionType string

const (
	NFTCollectionTypeERC721  NFTCollectionType = "erc721"
	NFTCollectionTypeERC1155 NFTCollectionType = "erc1155"
)

func ParseNFTCollectionType(t string) (NFTCollectionType, error) {
	tokenType := strings.ToLower(t)
	switch tokenType {
	case "erc721":
		return NFTCollectionTypeERC721, nil
	case "erc1155":
		return NFTCollectionTypeERC1155, nil
	default:
		return "", fmt.Errorf("unknown nft collection type: %+v", tokenType)
	}
}

type NFTCollection struct {
	bun.BaseModel `bun:"table:eth_nft_collection"`
	model.BaseWithID

	Blockchain      string            `bun:"blockchain,notnull"`
	Network         string            `bun:"network,notnull"`
	ContractAddress string            `bun:"contract_address,notnull"`
	Name            string            `bun:"name,notnull"`
	FromBlockHeight *bunbig.Int       `bun:"from_block_height,notnull"`
	TotalSupply     *bunbig.Int       `bun:"total_supply"`
	Type            NFTCollectionType `bun:"type,notnull"`
}

func (c NFTCollection) ToAPIModel() apimodel.NFTCollection {
	var totalSupply *big.Int
	if c.TotalSupply != nil {
		totalSupply = c.TotalSupply.ToMathBig()
	}

	return apimodel.NFTCollection{
		ID:              c.ID,
		Blockchain:      c.Blockchain,
		Network:         c.Network,
		Name:            c.Name,
		ContractAddress: c.ContractAddress,
		CreatedAt:       c.CreatedAt,
		BlockHeight:     *c.FromBlockHeight.ToMathBig(),
		TotalSupply:     totalSupply,
		Type:            string(c.Type),
	}
}
