package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/andrewtanjaya21/test_go/graph/generated"
	"github.com/andrewtanjaya21/test_go/graph/model"
)

func (r *mutationResolver) CreateFood(ctx context.Context, input model.NewFood) (*model.Food, error) {
	food := model.Food{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
	_,err := r.DB.Model(&food).Insert()

	if(err != nil){
		return nil, errors.New("Insert Food Failed")
	}

	return &food, nil
}

func (r *mutationResolver) UpdateFood(ctx context.Context, id string, input model.NewFood) (*model.Food, error) {
	var food model.Food
	err := r.DB.Model(&food).Where("id = ?", id).First()
	if err!= nil{
		return nil, errors.New("Update Food Failed")
	}

	food.Price = input.Price
	food.Description = input.Description
	food.Name = input.Name

	_,error := r.DB.Model(&food).Where("id = ?", id).Update()

	if error != nil{
		return nil, errors.New("Update Food Failed")
	}
	return &food, nil
}

func (r *mutationResolver) DeleteFood(ctx context.Context, id string) (bool, error) {
	var food model.Food
	err := r.DB.Model(&food).Where("id = ?", id).First()
	if err!= nil{
		return false, errors.New("Delete Food Failed")
	}
	_,error := r.DB.Model(&food).Where("id = ?", id).Delete()

	if error != nil{
		return false, errors.New("Delete Food Failed")
	}
	return true, nil
}

func (r *queryResolver) Foods(ctx context.Context) ([]*model.Food, error) {
	var foods []*model.Food
	err := r.DB.Model(&foods).Select()
	if err != nil{
		return nil, errors.New("Failed query foods")
	}
	return foods, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
