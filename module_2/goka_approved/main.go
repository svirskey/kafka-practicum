package main


import (
   "github.com/lovoo/goka"

   "github.com/svirskey/kafka-practicum/module_2/goka_approved/blocker"
   "github.com/svirskey/kafka-practicum/module_2/goka_approved/emitter"
   "github.com/svirskey/kafka-practicum/module_2/goka_approved/filter"
   "github.com/svirskey/kafka-practicum/module_2/goka_approved/user"
)


var brokers = []string{"127.0.0.1:29092"}


var (
   Emitter2FilterTopic       goka.Stream = "emitter2filter-stream"
   Filter2UserProcessorTopic goka.Stream = "filter2userprocessor-stream"
   BlockerTopic              goka.Stream = "blocker-stream"
)


func main() {
   go blocker.RunBlocker(brokers, BlockerTopic)
   go emitter.RunEmitter(brokers, Emitter2FilterTopic)
   go filter.RunFilter(brokers, Emitter2FilterTopic, Filter2UserProcessorTopic)
   go user.RunUserProcessor(brokers, Filter2UserProcessorTopic)
   user.RunUserView(brokers)
}