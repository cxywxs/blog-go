package dao

import (
	"blog-go/entity"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

func GetAllArticle() []entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	var result []entity.Article
	iter := collection.Find(nil).Iter()
	_ = iter.All(&result)
	return result

}

func GetArticle(condition entity.SelectCondition) entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	article := entity.Article{}
	_ = collection.Find(bson.M{condition.Field: bson.M{condition.Key: condition.Value}}).One(&article)
	return article
}

func GetArticles(condition entity.SelectCondition) []entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	var articles []entity.Article
	_ = collection.Find(bson.M{condition.Field: bson.M{condition.Key: condition.Value}}).All(&articles)
	return articles
}

func GetArticleSort(field string) []entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	var articles []entity.Article
	_ = collection.Find(nil).Sort("-" + field).All(&articles)
	return articles
}

func InsertArticle(article entity.Article) {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	err = collection.Insert(article)
	if err != nil {
		log.Fatalln(err)
	}

}
