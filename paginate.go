package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (p Pager) PageCheck(c echo.Context) error {
	var (
		err    error
		fp     bool
		filter bson.D
		key    = p.QueryKey
		val    = p.QueryVal
	)
	if lim := c.QueryParam("limit"); lim == "" {
		p.Limit = 20
		if n := c.QueryParam("next"); n != "" {
			objID, err := primitive.ObjectIDFromHex(n)
			if err != nil {
				return err
			}
			filter = bson.D{
				{key, val},
				{"_id", bson.D{{"$gt", objID}}},
			}
			fp = false
		} else if p := c.QueryParam("previous"); p != "" {
			objID, err := primitive.ObjectIDFromHex(p)
			if err != nil {
				return err
			}
			filter = bson.D{
				{key, val},
				{"_id", bson.D{{"$lt", objID}}},
			}
			fp = false
		} else {
			filter = bson.D{
				{key, val},
			}
			fp = true
		}
	} else {
		p.Limit, err = strconv.ParseInt(lim, 32, 64)
		if err != nil {
			log.Println("Error Occurred: ", err)
			return err
		}
		if n := c.QueryParam("next"); n != "" {
			objID, err := primitive.ObjectIDFromHex(c.QueryParam("next"))
			if err != nil {
				return err
			}
			filter = bson.D{
				{key, val},
				{"_id", bson.D{{"$gt", objID}}},
			}
			fp = false
		} else if p := c.QueryParam("previous"); p != "" {
			objID, err := primitive.ObjectIDFromHex(p)
			if err != nil {
				return err
			}
			filter = bson.D{
				{key, val},
				{"_id", bson.D{{"$lt", objID}}},
			}
			fp = false
		} else {
			filter = bson.D{
				{key, val},
			}
			fp = true
		}
	}
	p.Filter = filter
	p.FirstPage = fp
	return nil
}
