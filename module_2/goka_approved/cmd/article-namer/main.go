package main

import (
	"flag"
	"log"

	"github.com/lovoo/goka"

	"github.com/svirskey/kafka-practicum/module_2/goka_approved/article-namer"
	articlenamer "github.com/svirskey/kafka-practicum/module_2/goka_approved/article-namer"
)


var (
   articleId = flag.Int("articleId", -1, "article id")
   articleName  = flag.String("articleName", "", "article name")
   broker  = flag.String("broker", "localhost:29092", "boostrap Kafka broker")
   stream  = flag.String("stream", "", "stream name")
)


func main() {
   flag.Parse()
   if *articleId == -1 || *articleName == "" {
      log.Fatal("Неверные входные данные")
   }
   emitter, err := goka.NewEmitter([]string{*broker}, goka.Stream(*stream), new(articlenamer.ArticleNameEventCodec))
   if err != nil {
      log.Fatal(err)
   }
   defer emitter.Finish()


   err = emitter.EmitSync(*articleId, &articlenamer.ArticleNameEvent{Name: *articleName})
   if err != nil {
      log.Fatal(err)
   }
}