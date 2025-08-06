package main

import (
	"flag"
	"log"

	"github.com/lovoo/goka"

	"github.com/svirskey/kafka-practicum/module_2/goka_approved/article-namer"
)


var (
   articleId = flag.String("articleId", "", "article id")
   articleName  = flag.String("articleName", "", "article name")
   broker  = flag.String("broker", "localhost:29092", "boostrap Kafka broker")
   stream  = flag.String("stream", "", "stream name")
)


func main() {
   flag.Parse()
   if *articleId == "" || *articleName == "" {
      log.Fatal("Неверные входные данные")
   }
   emitter, err := goka.NewEmitter([]string{*broker}, goka.Stream(*stream), new(articlenamer.ArticleNameCodec))
   if err != nil {
      log.Fatal(err)
   }
   defer emitter.Finish()

   log.Printf("[proc] article with id: %s named: %s \n", *articleId, *articleName)

   err = emitter.EmitSync(*articleId, &articlenamer.ArticleName{Name: *articleName})
   if err != nil {
      log.Fatal(err)
   }
}