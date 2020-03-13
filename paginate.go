package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

func (p Pager) Paginate() PaginateResult {
	var res PaginateResult
	cur, err := collection.Find(context.TODO(), p.Filter, options.Find().SetLimit(p.Limit))
	if err != nil {
		log.Println("Error Occurred: ", err)
	}
	for cur.Next(context.TODO()) {
		var device Device
		err := cur.Decode(&device)
		if err != nil {
			log.Println("Error Occurred: ", err)
		}
		res.Data = append(res.Data, device)
	}
	limit := strconv.FormatInt(p.Limit, 10)
	if !p.FirstPage {
		res.PrevPage = p.URL + "?limit=" + limit + "&prev=" + res.Data[0].ID.Hex()
	}
	res.NextPage = p.URL + "?limit=" + limit + "&next=" + res.Data[len(res.Data)-1].ID.Hex()
	return res
}
