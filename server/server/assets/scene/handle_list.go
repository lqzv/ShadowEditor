// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package scene

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tengge1/shadoweditor/helper"
	"github.com/tengge1/shadoweditor/server"
	"github.com/tengge1/shadoweditor/server/category"
)

func init() {
	server.Mux.UsingContext().Handle(http.MethodGet, "/api/Scene/List", List)
}

// List returns scene list.
func List(w http.ResponseWriter, r *http.Request) {
	db, err := server.Mongo()
	if err != nil {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  err.Error(),
		})
		return
	}

	// get all categories
	categories := []category.Model{}
	db.FindAll(server.CategoryCollectionName, &categories)

	docs := bson.A{}
	opts := options.FindOptions{
		Sort: bson.M{
			"UpdateTime": -1,
		},
	}

	if server.Config.Authority.Enabled {
		user, _ := server.GetCurrentUser(r)

		if user != nil {
			filter1 := bson.M{
				"$or": bson.A{
					bson.M{
						"UserID": user.ID,
					},
					bson.M{
						"IsPublic": true,
					},
				},
			}

			if user.Name == "Administrator" {
				filter1 = bson.M{
					"$or": bson.A{
						filter1,
						bson.M{
							"UserID": bson.M{
								"$exists": 0,
							},
						},
					},
				}
			}
			db.FindMany(server.SceneCollectionName, filter1, &docs, &opts)
		} else { // no login user can see public scenes
			filter1 := bson.M{
				"IsPublic": true,
			}
			db.FindMany(server.SceneCollectionName, filter1, &docs, &opts)
		}
	} else {
		db.FindAll(server.SceneCollectionName, &docs, &opts)
	}

	list := []Model{}

	for _, i := range docs {
		doc := i.(primitive.D).Map()

		categoryID := ""
		categoryName := ""

		if doc["Category"] != nil {
			for _, category := range categories {
				if category.ID == doc["Category"].(string) {
					categoryID = category.ID
					categoryName = category.Name
					break
				}
			}
		}

		thumbnail := ""
		if doc["Thumbnail"] != nil {
			thumbnail = doc["Thumbnail"].(string)
		}

		isPublic := false
		if doc["IsPublic"] != nil {
			isPublic = doc["IsPublic"].(bool)
		}

		info := Model{
			ID:             doc["ID"].(primitive.ObjectID).Hex(),
			Name:           doc["Name"].(string),
			CategoryID:     categoryID,
			CategoryName:   categoryName,
			TotalPinYin:    helper.PinYinToString(doc["TotalPinYin"]),
			FirstPinYin:    helper.PinYinToString(doc["FirstPinYin"]),
			CollectionName: doc["CollectionName"].(string),
			Version:        int(doc["Version"].(int32)),
			CreateTime:     doc["CreateTime"].(primitive.DateTime).Time(),
			UpdateTime:     doc["UpdateTime"].(primitive.DateTime).Time(),
			Thumbnail:      thumbnail,
			IsPublic:       isPublic,
		}
		list = append(list, info)
	}

	helper.WriteJSON(w, server.Result{
		Code: 200,
		Msg:  "Get Successfully!",
		Data: list,
	})
}
