// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package animation

import (
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tengge1/shadoweditor/helper"
	"github.com/tengge1/shadoweditor/server"
)

func init() {
	server.Mux.UsingContext().Handle(http.MethodPost, "/api/Animation/Edit", Edit)
}

// Edit change the name, category and thumbnail of an animation.
func Edit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := primitive.ObjectIDFromHex(strings.TrimSpace(r.FormValue("ID")))
	if err != nil {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  "ID is not allowed.",
		})
	}

	name := strings.TrimSpace(r.FormValue("Name"))
	if name == "" {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  "Name is not allowed to be empty.",
		})
		return
	}

	image := strings.TrimSpace(r.FormValue("Image"))
	category := strings.TrimSpace(r.FormValue("Category"))

	// update mongo
	db, err := server.Mongo()
	if err != nil {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  err.Error(),
		})
		return
	}

	pinyin := helper.ConvertToPinYin(name)

	filter := bson.M{
		"_id": id,
	}
	set := bson.M{
		"Name":        name,
		"TotalPinYin": pinyin.TotalPinYin,
		"FirstPinYin": pinyin.FirstPinYin,
		"Thumbnail":   image,
	}
	update := bson.M{
		"$set": set,
	}
	if category == "" {
		update["$unset"] = bson.M{
			"Category": 1,
		}
	} else {
		set["Category"] = category
	}

	db.UpdateOne(server.AnimationCollectionName, filter, update)

	helper.WriteJSON(w, server.Result{
		Code: 200,
		Msg:  "Saved successfully!",
	})
}
