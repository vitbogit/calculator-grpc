package converter

import (
	"math/big"
	"sum/internal/model"
	desc "sum/pkg/sum_v1"
)

// func ToUserFromService(user *model.User) *desc.User {
// 	var updatedAt *timestamppb.Timestamp
// 	if user.UpdatedAt != nil {
// 		updatedAt = timestamppb.New(*user.UpdatedAt)
// 	}

// 	return &desc.User{
// 		Uuid:      user.UUID,
// 		Info:      ToUserInfoFromService(user.Info),
// 		CreatedAt: timestamppb.New(user.CreatedAt),
// 		UpdatedAt: updatedAt,
// 	}
// }

// func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
// 	return &desc.UserInfo{
// 		FirstName: info.FirstName,
// 		LastName:  info.LastName,
// 		Age:       info.Age,
// 	}
// }

// func ToIntegersFromDesc(info *desc.CalculateRequest) *model.Integers {
// 	return &model.Integers{
// 		A: info.GetA(),
// 		B: info.GetB(),
// 	}
// }

func ToFloatsFromDesc(info *desc.CalculateRequest) *model.Floats {
	a := new(big.Float)
	a.SetString(info.GetA())

	b := new(big.Float)
	b.SetString(info.GetB())

	return &model.Floats{
		A: a,
		B: b,
	}
}

func ToFractionalsFromDesc(info *desc.CalculateFractionalRequest) *model.Fractionals {
	a1 := new(big.Float)
	a1.SetString(info.GetA1())

	a2 := new(big.Float)
	a2.SetString(info.GetA2())

	b1 := new(big.Float)
	b1.SetString(info.GetB1())

	b2 := new(big.Float)
	b2.SetString(info.GetB2())

	return &model.Fractionals{
		A1: a1,
		A2: a2,
		B1: b1,
		B2: b2,
	}
}
