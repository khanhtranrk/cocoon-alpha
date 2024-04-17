package cocoon

import (
	"log"
)

type Office struct {
  borker *Broker

  IncomingLetter chan *Letter
  OutgoingLetter chan *Letter
  AcceptedLetter chan *Letter
  RejectedLetter chan *Letter
  PendingLetter chan *Letter
  UnsentLetter chan *Letter
}

func (o *Office) HandleInComingLetter() {
  msgs, err := o.borker.channel.Consume("mavis", "rabit", true, false, false, false, nil)
  if err != nil {
      log.Fatalf(err.Error())
  }

  for d := range msgs {
    b.hub.broadcast <-d.Body
  }
}
